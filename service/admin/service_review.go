package admin

import (
	"context"
	"time"

	"github.com/bots-house/birzzha/core"
	"github.com/bots-house/birzzha/pkg/log"
	"github.com/bots-house/birzzha/pkg/tg"
	"github.com/pkg/errors"
	"github.com/volatiletech/null/v8"
)

type ReviewInput struct {
	TelegramID int
	FirstName  string
	LastName   string
	Username   string
	Text       string
	CreatedAt  time.Time
}

const (
	userAvatarDir = "user"
)

func (srv *Service) AddReview(ctx context.Context, user *core.User, in *ReviewInput) error {
	err := srv.IsAdmin(user)
	if err != nil {
		return err
	}

	var photoURL string
	var avatar string

	if in.Username != "" {
		photoURL, err = srv.AvatarResolver.ResolveAvatar(ctx, in.Username)
		if err != nil && err == tg.ErrCantDownloadAvatar {
			log.Warn(ctx, "err", errors.Wrap(err, "resolve avatar"))
		}
	}

	if photoURL != "" {
		avatar, err = srv.Storage.AddByURL(ctx, userAvatarDir, photoURL)
		if err != nil {
			log.Warn(ctx, "can't download user avatar", "url", photoURL, "err", errors.Wrap(err, "upload avatar to storage"))
		}
	}

	review := core.NewReview(
		core.NewReviewUser(
			null.NewInt(in.TelegramID, in.TelegramID != 0),
			in.FirstName,
			null.NewString(in.LastName, in.LastName != ""),
			null.NewString(in.Username, in.Username != ""),
			null.NewString(avatar, avatar != ""),
		),
		in.Text,
		in.CreatedAt,
	)

	if err := srv.Review.Add(ctx, review); err != nil {
		return err
	}

	return nil
}
