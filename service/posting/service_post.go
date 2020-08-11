package posting

import (
	"context"
	"time"

	"github.com/bots-house/birzzha/core"
	tgbotapi "github.com/bots-house/telegram-bot-api"
	"github.com/pkg/errors"
	"github.com/volatiletech/null/v8"
)

func (srv *Service) GetText(ctx context.Context, id core.LotID) (string, error) {
	lot, err := srv.Lot.Query().ID(id).One(ctx)
	if err != nil {
		return "", errors.Wrap(err, "get lot")
	}

	topics, err := srv.Topic.Query().ID(lot.TopicIDs...).All(ctx)
	if err != nil {
		return "", errors.Wrap(err, "get topics")
	}

	owner, err := srv.User.Query().ID(lot.OwnerID).One(ctx)
	if err != nil {
		return "", errors.Wrap(err, "get owner")
	}

	settings, err := srv.Settings.Get(ctx)
	if err != nil {
		return "", errors.Wrap(err, "get settings")
	}

	lt := &LotPostText{
		Lot:      lot,
		Topics:   topics,
		Owner:    owner,
		Settings: settings,
	}

	return lt.Render()
}

func (srv *Service) SendPreview(ctx context.Context, user *core.User, post string) error {
	_, err := srv.TgClient.Send(tgbotapi.MessageConfig{
		BaseChat:              tgbotapi.BaseChat{ChatID: int64(user.Telegram.ID)},
		Text:                  post,
		ParseMode:             "HTML",
		DisableWebPagePreview: true,
	})
	if err != nil {
		return errors.Wrap(err, "send lot text")
	}
	return nil
}

func (srv *Service) SendPosts(ctx context.Context) error {
	return srv.Txier(ctx, func(ctx context.Context) error {
		settings, err := srv.Settings.Get(ctx)
		if err != nil {
			return errors.Wrap(err, "get settings")
		}

		posts, err := srv.Post.Pull(ctx)
		if err != nil {
			return errors.Wrap(err, "get posts")
		}

		lots, err := srv.Lot.Query().ID(posts.LotIDs()...).All(ctx)
		if err != nil {
			return errors.Wrap(err, "get lots")
		}

		for _, post := range posts {
			_, err := srv.TgClient.Send(tgbotapi.MessageConfig{
				BaseChat:              tgbotapi.BaseChat{ChatID: settings.Channel.PrivateID},
				Text:                  post.Text,
				ParseMode:             "HTML",
				DisableWebPagePreview: post.DisableWebPagePreview,
			})
			if err != nil {
				return errors.Wrap(err, "send post")
			}

			post.PublishedAt = null.TimeFrom(time.Now())
			if err := srv.Post.Update(ctx, post); err != nil {
				return errors.Wrap(err, "update post")
			}

			lot := lots.Find(post.LotID)
			lot.Status = core.LotStatusPublished
			lot.PublishedAt = post.PublishedAt
			if err := srv.Lot.Update(ctx, lot); err != nil {
				return errors.Wrap(err, "update lot")
			}
		}
		return nil
	})
}
