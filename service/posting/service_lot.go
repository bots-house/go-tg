package posting

import (
	"context"

	"github.com/bots-house/birzzha/core"
	tgbotapi "github.com/bots-house/telegram-bot-api"
	"github.com/pkg/errors"
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
