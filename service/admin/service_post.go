package admin

import (
	"context"
	"time"

	"github.com/bots-house/birzzha/core"
	"github.com/pkg/errors"
)

type PostInput struct {
	LotID                 core.LotID
	Text                  string
	DisableWebPagePreview bool
	ScheduledAt           time.Time
}

var (
	ErrLotIsAlreadyPublished = core.NewError("lot_is_already_published", "lot is already published")
)

func (srv *Service) CreatePost(ctx context.Context, user *core.User, in *PostInput) (*core.Post, error) {
	if err := srv.IsAdmin(user); err != nil {
		return nil, err
	}
	post := core.NewPost(in.LotID, in.Text, in.DisableWebPagePreview, in.ScheduledAt)

	if err := srv.Txier(ctx, func(ctx context.Context) error {
		lot, err := srv.Lot.Query().ID(in.LotID).One(ctx)
		if err != nil {
			return errors.Wrap(err, "get lot")
		}

		if lot.Status == core.LotStatusPublished {
			return ErrLotIsAlreadyPublished
		}

		if err := srv.Post.Add(ctx, post); err != nil {
			return errors.Wrap(err, "add post")
		}

		lot.Status = core.LotStatusScheduled
		if err := srv.Lot.Update(ctx, lot); err != nil {
			return errors.Wrap(err, "update lot")
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return post, nil

}
