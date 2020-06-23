package bot

import (
	"context"

	tgbotapi "github.com/bots-house/telegram-bot-api"

	"github.com/pkg/errors"

	"github.com/bots-house/birzzha/pkg/log"
	"github.com/bots-house/birzzha/pkg/tg"
	"github.com/bots-house/birzzha/service/auth"
)

func newAuthMiddleware(srv *auth.Service) tg.Middleware {
	return func(next tg.Handler) tg.Handler {
		return tg.HandlerFunc(func(ctx context.Context, update *tgbotapi.Update) error {
			var tgUser *tgbotapi.User

			switch {
			case update.Message != nil:
				tgUser = update.Message.From
			case update.EditedMessage != nil:
				tgUser = update.EditedMessage.From
			case update.CallbackQuery != nil:
				tgUser = update.CallbackQuery.From
			case update.InlineQuery != nil:
				tgUser = update.InlineQuery.From
			default:
				log.Warn(ctx, "unsupported update", "id", update.UpdateID)
				return nil
			}

			user, err := srv.AuthorizeInBot(ctx, &auth.TelegramUserInfo{
				ID:           tgUser.ID,
				FirstName:    tgUser.FirstName,
				LastName:     tgUser.LastName,
				Username:     tgUser.UserName,
				LanguageCode: tgUser.LanguageCode,
			})

			if err != nil {
				return errors.Wrap(err, "auth service")
			}

			ctx = withUser(ctx, user)

			return next.HandleUpdate(ctx, update)
		})
	}
}
