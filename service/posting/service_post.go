package posting

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/bots-house/birzzha/core"
	"github.com/bots-house/birzzha/pkg/log"
	tgbotapi "github.com/bots-house/telegram-bot-api"
	"github.com/pkg/errors"
	"github.com/volatiletech/null/v8"
)

func joinSitePath(site, path string) string {
	path = strings.TrimPrefix(path, "/")
	site = strings.TrimSuffix(site, "/")

	return site + "/" + path
}

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
		Lot:                     lot,
		Topics:                  topics,
		Owner:                   owner,
		Settings:                settings,
		SiteWithPathListChannel: srv.SiteWithPathListChannel,
	}

	s, err := lt.Render()
	if err != nil {
		return "", err
	}

	return strings.ReplaceAll(s, "\\n", "\n"), nil
}

func (srv *Service) SendPreview(ctx context.Context, user *core.User, post string, lot *core.Lot) error {
	msg := tgbotapi.MessageConfig{
		BaseChat:              tgbotapi.BaseChat{ChatID: int64(user.Telegram.ID)},
		Text:                  post,
		ParseMode:             "HTML",
		DisableWebPagePreview: true,
	}

	if lot != nil {
		markup := tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.InlineKeyboardButton{
					Text: "Подробней",
					LoginURL: &tgbotapi.LoginURL{
						URL: joinSitePath(srv.Config.Site, fmt.Sprintf("lots/%d?from=channel", lot.ID)),
					},
				},
			),
		)
		msg.BaseChat.ReplyMarkup = markup
	}

	_, err := srv.TgClient.Send(msg)
	if err != nil {
		return errors.Wrap(err, "send lot text")
	}
	return nil
}

func (srv *Service) SendPosts(ctx context.Context) error {
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
		msg := tgbotapi.MessageConfig{
			BaseChat: tgbotapi.BaseChat{
				ChatID: settings.Channel.PrivateID,
			},
			Text:                  post.Text,
			ParseMode:             "HTML",
			DisableWebPagePreview: post.DisableWebPagePreview,
		}
		if post.LotID != 0 && post.Buttons.LotLink {
			markup := tgbotapi.NewInlineKeyboardMarkup(
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.InlineKeyboardButton{
						Text: "Подробнее",
						LoginURL: &tgbotapi.LoginURL{
							URL: joinSitePath(srv.Config.Site, fmt.Sprintf("lots/%d?from=channel", post.LotID)),
						},
					},
				),
			)
			msg.BaseChat.ReplyMarkup = markup
		}
		msg1, err := srv.TgClient.Send(msg)
		if err != nil {
			post.Status = core.PostStatusFailed
			if err := srv.Post.Update(ctx, post); err != nil {
				log.Error(ctx, "update post", "error", err)
			}
			log.Error(ctx, "send post", "error", err)
		} else {

			post.PublishedAt = null.TimeFrom(time.Now())
			post.MessageID = null.IntFrom(msg1.MessageID)
			post.Status = core.PostStatusPublished
			if err := srv.Post.Update(ctx, post); err != nil {
				log.Error(ctx, "update post", "error", err)
			}

			if post.LotID != 0 {
				lot := lots.Find(post.LotID)
				lot.Status = core.LotStatusPublished
				lot.PublishedAt = post.PublishedAt
				if err := srv.Lot.Update(ctx, lot); err != nil {
					log.Error(ctx, "update lot", "error", err)
				}

				srv.Notify.SendUser(lot.OwnerID, userLotPublishedNotification{
					PostID: post.MessageID.Int,
					Lot:    lot,
				})
			}
		}
	}
	return nil
}
