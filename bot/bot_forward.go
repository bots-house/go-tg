package bot

import (
	"context"

	"github.com/bots-house/birzzha/service/admin"
	tgbotapi "github.com/bots-house/telegram-bot-api"
	"github.com/pkg/errors"
)

const (
	textReviewAddedSuccess   = "👍 Отзыв успешно добавлен."
	reviewTextCantBeNullable = "В отзыве должен присутствовать текст."
)

func (bot *Bot) onForward(ctx context.Context, msg *tgbotapi.Message) error {
	if msg.Text == "" {
		answ := bot.newAnswerMsg(ctx, msg, reviewTextCantBeNullable)
		return bot.send(ctx, answ)
	}

	err := bot.adminSrv.AddReview(ctx, getUserCtx(ctx), &admin.ReviewInput{
		TelegramID: msg.ForwardFrom.ID,
		FirstName:  msg.ForwardFrom.FirstName,
		LastName:   msg.ForwardFrom.LastName,
		Username:   msg.ForwardFrom.UserName,
		Text:       msg.Text,
	})
	if err != nil {
		return errors.Wrap(err, "Add review")
	}

	answ := bot.newAnswerMsg(ctx, msg, textReviewAddedSuccess)
	return bot.send(ctx, answ)
}
