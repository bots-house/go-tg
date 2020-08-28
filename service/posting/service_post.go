package posting

import (
	"context"
	"fmt"
	"strconv"
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

	return lt.Render()
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
							Text: "Подробней",
							LoginURL: &tgbotapi.LoginURL{
								URL: joinSitePath(srv.Config.Site, fmt.Sprintf("lots/%d?from=channel", post.LotID)),
							},
						},
					),
				)
				msg.BaseChat.ReplyMarkup = markup
			}
			_, err := srv.TgClient.Send(msg)
			if err != nil {
				return errors.Wrap(err, "send post")
			}

			post.PublishedAt = null.TimeFrom(time.Now())
			if err := srv.Post.Update(ctx, post); err != nil {
				return errors.Wrap(err, "update post")
			}

			if post.LotID != 0 {
				lot := lots.Find(post.LotID)
				lot.Status = core.LotStatusPublished
				lot.PublishedAt = post.PublishedAt
				if err := srv.Lot.Update(ctx, lot); err != nil {
					return errors.Wrap(err, "update lot")
				}

				date, err := formatDate(lot.ScheduledAt.Time)
				if err == nil {
					srv.UserNotification.Send(ctx, lot, LotPublishedNotification{PublishedAt: date})
				} else {
					log.Error(ctx, "failed to format date to send lot published notification", "error", err, "time", lot.ScheduledAt.Time)
				}
			}
		}
		return nil
	})
}

func formatDate(t time.Time) (string, error) {
	location, err := time.LoadLocation("Europe/Moscow")
	if err != nil {
		return "", errors.Wrap(err, "load moscow location")
	}
	now := time.Now().In(location)
	t = t.In(location)
	if t.Day() == now.Day() {
		return "сегодня в " + t.Format("15:04"), nil
	}

	month, ok := mskMonth[t.Month()]
	if !ok {
		return "", errors.New("failed to match defined month in MSK")
	}

	return strconv.Itoa(t.Day()) + " " + month + " в " + t.Format("15:04"), nil
}

var mskMonth = map[time.Month]string{
	time.January:   "января",
	time.February:  "ферваля",
	time.March:     "марта",
	time.April:     "апреля",
	time.May:       "мая",
	time.June:      "июня",
	time.July:      "июля",
	time.August:    "августа",
	time.September: "сентября",
	time.October:   "октябрь",
	time.November:  "ноябрь",
	time.December:  "декабрь",
}
