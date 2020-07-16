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

type ReviewList struct {
	Total int
	Items core.ReviewSlice
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
			log.Warn(ctx, "resolve avatar", "err", err)
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

func (srv *Service) DeleteReview(ctx context.Context, user *core.User, id core.ReviewID) error {
	if err := srv.IsAdmin(user); err != nil {
		return err
	}

	if err := srv.Review.Delete(ctx, id); err != nil {
		return errors.Wrap(err, "delete review")
	}
	return nil
}

func (srv *Service) UpdateReview(ctx context.Context, user *core.User, id core.ReviewID, text string) (*core.Review, error) {
	if err := srv.IsAdmin(user); err != nil {
		return nil, err
	}

	review, err := srv.Review.Query().ID(id).One(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get review")
	}

	review.Text = text
	if err := srv.Review.Update(ctx, review); err != nil {
		return nil, errors.Wrap(err, "update review")
	}
	return review, nil
}

func (srv *Service) GetReviews(ctx context.Context, user *core.User, offset int, limit int) (*ReviewList, error) {
	if err := srv.IsAdmin(user); err != nil {
		return nil, err
	}

	reviews, err := srv.Review.Query().Limit(limit).Offset(offset).OrderByCreatedAt().All(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get reviews")
	}

	total, err := srv.Review.Query().Count(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get reviews count")
	}

	return &ReviewList{
		Total: total,
		Items: reviews,
	}, nil
}
