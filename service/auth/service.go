package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/benbjohnson/clock"
	"github.com/bots-house/birzzha/core"
	"github.com/bots-house/birzzha/pkg/log"
	"github.com/bots-house/birzzha/pkg/storage"
	tgbotapi "github.com/bots-house/telegram-bot-api"
	"github.com/cristalhq/jwt"
	"github.com/pkg/errors"
	"github.com/rs/xid"
	"github.com/volatiletech/null/v8"
)

const (
	userAvatarDir = "user"
)

type Service struct {
	UserStore core.UserStore
	Config    Config
	Clock     clock.Clock
	Storage   storage.Storage
	Bot       *tgbotapi.BotAPI

	botLogins     map[string]*LoginViaBotInfo
	botLoginsLock sync.Mutex
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

func (srv *Service) newUserFromtelegramUserInfo(ctx context.Context, info *TelegramUserInfo) (*core.User, error) {
	user := &core.User{
		Telegram: core.UserTelegram{
			ID:       info.ID,
			Username: null.NewString(info.Username, info.Username != ""),
		},
		FirstName: info.FirstName,
		LastName:  null.NewString(info.LastName, info.LastName != ""),
		// Avatar:     null.NewString(avatar, avatar != ""),
		JoinedFrom: core.JoinedFromBot,
		JoinedAt:   time.Now(),
	}

	if err := srv.UserStore.Add(ctx, user); err != nil {
		return nil, errors.Wrap(err, "add user to store")
	}

	return user, nil
}

func (srv *Service) AuthorizeInBot(ctx context.Context, info *TelegramUserInfo) (*core.User, error) {
	user, err := srv.UserStore.Query().TelegramID(info.ID).One(ctx)
	if err == core.ErrUserNotFound {
		user, err = srv.newUserFromtelegramUserInfo(ctx, info)
		if err != nil {
			return nil, errors.Wrap(err, "new user from telegram")
		}
	} else if err != nil {
		return nil, errors.Wrap(err, "query user from store")
	}

	return user, nil
}

type LoginViaBotInfo struct {
	CallbackURL string

	createdAt time.Time
}

func (srv *Service) saveLoginViaBot(ctx context.Context, info *LoginViaBotInfo) (string, error) {
	id := xid.New().String()

	info.createdAt = time.Now()

	srv.botLoginsLock.Lock()
	if srv.botLogins == nil {
		srv.botLogins = make(map[string]*LoginViaBotInfo)
	}
	srv.botLogins[id] = info
	srv.botLoginsLock.Unlock()

	return id, nil
}

func (srv *Service) LoginViaBot(ctx context.Context, info *LoginViaBotInfo) (string, error) {
	id, err := srv.saveLoginViaBot(ctx, info)
	if err != nil {
		return "", errors.Wrap(err, "save login via bot")
	}
	url := fmt.Sprintf("https://t.me/%s?start=login_%s", srv.Bot.Self.UserName, id)

	return url, nil
}

var ErrBotLoginNotFound = errors.New("bot login not found")

func (srv *Service) PopLoginViaBot(ctx context.Context, id string) (*LoginViaBotInfo, error) {
	srv.botLoginsLock.Lock()
	defer srv.botLoginsLock.Unlock()

	v, ok := srv.botLogins[id]
	if !ok {
		return nil, ErrBotLoginNotFound
	}

	if srv.Clock.Since(v.createdAt) > time.Minute*3 {
		delete(srv.botLogins, id)
		return nil, ErrBotLoginNotFound
	}

	return v, nil
}
