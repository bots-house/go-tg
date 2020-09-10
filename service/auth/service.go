package auth

import (
	"context"
	"encoding/json"
	"net/url"
	"strconv"
	"time"

	"github.com/benbjohnson/clock"
	"github.com/bots-house/birzzha/core"
	"github.com/bots-house/birzzha/pkg/log"
	"github.com/bots-house/birzzha/pkg/storage"
	"github.com/bots-house/birzzha/service/notifications"
	"github.com/cristalhq/jwt"
	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
	"github.com/rs/xid"
	"github.com/volatiletech/null/v8"
)

const (
	userAvatarDir = "user"

	loginViaBotKey = "login:via:bot:"
)

type Service struct {
	UserStore      core.UserStore
	Config         Config
	Clock          clock.Clock
	Storage        storage.Storage
	Notifications  *notifications.Notifications
	BotLinkBuilder *core.BotLinkBuilder
	Redis          redis.UniversalClient

	AdminNotificationsChannelID int64
}

var (
	ErrInvalidTelegramWidgetInfo = core.NewError("telegram_widget_info_invalid", "telegram widget info signature is invalid")
	ErrExpiredTelegramWidgetInfo = core.NewError("telegram_widget_info_expired", "telegram widget info is expired")
)

type Token struct {
	Token string
	User  *core.User
}

func (srv *Service) newUserFromTelegramWidgetInfo(ctx context.Context, info *TelegramWidgetInfo) (*core.User, error) {
	var avatar string

	if info.PhotoURL != "" {
		var err error
		avatar, err = srv.Storage.AddByURL(ctx, userAvatarDir, info.PhotoURL)
		if err != nil {
			log.Warn(ctx, "can't download user avatar", "url", info.PhotoURL, "err", err)
		}
	}

	user := &core.User{
		Telegram: core.UserTelegram{
			ID:       info.ID,
			Username: null.NewString(info.Username, info.Username != ""),
		},
		FirstName:  info.FirstName,
		LastName:   null.NewString(info.LastName, info.LastName != ""),
		Avatar:     null.NewString(avatar, avatar != ""),
		JoinedFrom: core.JoinedFromSite,
		JoinedAt:   time.Now(),
	}

	if err := srv.UserStore.Add(ctx, user); err != nil {
		return nil, errors.Wrap(err, "add user to store")
	}

	srv.Notifications.Send(newUserNotification{
		User: user,
	})

	return user, nil
}

func (srv *Service) newToken(user *core.User) (*Token, error) {
	signer, err := jwt.NewHS256([]byte(srv.Config.TokenSecret))
	if err != nil {
		return nil, errors.Wrap(err, "create signer")
	}

	builder := jwt.NewTokenBuilder(signer)

	issuedAt := srv.Clock.Now()
	expiresAt := issuedAt.Add(srv.Config.TokenLifeTime)

	tkn, err := builder.Build(&jwt.StandardClaims{
		Subject:   strconv.Itoa(int(user.ID)),
		IssuedAt:  jwt.Timestamp(issuedAt.Unix()),
		ExpiresAt: jwt.Timestamp(expiresAt.Unix()),
	})
	if err != nil {
		return nil, errors.Wrap(err, "build token")
	}

	return &Token{
		User:  user,
		Token: string(tkn.Raw()),
	}, nil
}

func (srv *Service) GetLoginWidgetInfo(ctx context.Context, user *core.User) url.Values {
	lwi := &TelegramWidgetInfo{
		ID:        user.Telegram.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName.String,
		Username:  user.Telegram.Username.String,
		AuthDate:  time.Now().Unix(),
	}

	if user.Avatar.Valid {
		lwi.PhotoURL = srv.Storage.PublicURL(user.Avatar.String)
	}

	return lwi.Encode(srv.Config.getBotTokenHash())
}

func (srv *Service) CreateToken(ctx context.Context, info *TelegramWidgetInfo) (*Token, error) {
	// check signature
	if ok, err := info.Check(srv.Config.getBotTokenHash()); !ok || err != nil {
		return nil, ErrInvalidTelegramWidgetInfo
	}

	// check if telegram widget info is not too old
	if srv.Clock.Since(info.AuthDateTime()) > srv.Config.WidgetInfoLifeTime {
		return nil, ErrExpiredTelegramWidgetInfo
	}

	user, err := srv.UserStore.Query().TelegramID(info.ID).One(ctx)
	if err == core.ErrUserNotFound {
		user, err = srv.newUserFromTelegramWidgetInfo(ctx, info)
		if err != nil {
			return nil, errors.Wrap(err, "new user from telegram")
		}
	} else if err != nil {
		return nil, errors.Wrap(err, "get user from store")
	}

	return srv.newToken(user)
}

func (srv *Service) Authorize(ctx context.Context, token *jwt.Token) (*core.User, error) {
	// check signature
	signer, err := jwt.NewHS256([]byte(srv.Config.TokenSecret))
	if err != nil {
		return nil, errors.Wrap(err, "create signer")
	}

	if err := signer.Verify(token.Payload(), token.Signature()); err != nil {
		return nil, errors.Wrap(err, "verify")
	}

	// parse claims
	claims := &jwt.StandardClaims{}

	if err := json.Unmarshal(token.RawClaims(), claims); err != nil {
		return nil, errors.Wrap(err, "parse claims")
	}

	// parse user id
	id, err := strconv.ParseInt(claims.Subject, 10, 64)
	if err != nil {
		return nil, errors.Wrap(err, "parse subject")
	}

	// fetch user by id

	user, err := srv.UserStore.Query().ID(core.UserID(id)).One(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get user from store")
	}

	return user, nil
}

func (srv *Service) newUserFromTelegramUserInfo(ctx context.Context, info *TelegramUserInfo) (*core.User, error) {
	var avatarURL string

	if info.GetAvatar != nil {
		avatar, err := info.GetAvatar(ctx)
		if err != nil {
			log.Warn(ctx, "get avatar failed", "err", err)
		}

		avatarURL, err = srv.Storage.AddByURL(ctx, userAvatarDir, avatar)
		if err != nil {
			log.Warn(ctx, "upload avatar failed", "err", err)
		}
	}

	user := &core.User{
		Telegram: core.UserTelegram{
			ID:       info.ID,
			Username: null.NewString(info.Username, info.Username != ""),
		},
		FirstName:  info.FirstName,
		LastName:   null.NewString(info.LastName, info.LastName != ""),
		Avatar:     null.NewString(avatarURL, avatarURL != ""),
		JoinedFrom: core.JoinedFromBot,
		JoinedAt:   time.Now(),
	}

	if err := srv.UserStore.Add(ctx, user); err != nil {
		return nil, errors.Wrap(err, "add user to store")
	}

	srv.Notifications.Send(newUserNotification{
		User: user,
	})

	return user, nil
}

func (srv *Service) AuthorizeInBot(ctx context.Context, info *TelegramUserInfo) (*core.User, error) {
	user, err := srv.UserStore.Query().TelegramID(info.ID).One(ctx)
	if err == core.ErrUserNotFound {
		user, err = srv.newUserFromTelegramUserInfo(ctx, info)
		if err != nil {
			return nil, errors.Wrap(err, "new user from telegram")
		}
	} else if err != nil {
		return nil, errors.Wrap(err, "query user from store")
	}

	return user, nil
}

func (srv *Service) saveLoginViaBot(ctx context.Context, callback string) (string, error) {
	id := xid.New().String()
	_, err := srv.Redis.Set(ctx, loginViaBotKey+id, callback, time.Minute*3).Result()

	return id, errors.Wrap(err, "failed to save callback")
}

func (srv *Service) LoginViaBot(ctx context.Context, callback string) (string, error) {
	id, err := srv.saveLoginViaBot(ctx, callback)
	if err != nil {
		return "", err
	}

	return srv.BotLinkBuilder.LoginURL(id), nil
}

var ErrBotLoginNotFound = errors.New("bot login not found")

func (srv *Service) PopLoginViaBot(ctx context.Context, id string) (string, error) {
	callback, err := srv.Redis.Get(ctx, loginViaBotKey+id).Result()

	switch err {
	case nil:
		return callback, nil
	case redis.Nil:
		return "", ErrBotLoginNotFound
	default:
		return "", errors.Wrapf(err, "failed to get callback login=%s from redis", loginViaBotKey+id)
	}
}
