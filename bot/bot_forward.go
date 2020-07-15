package bot

import (
	"context"
	"strings"
	"time"

	"github.com/bots-house/birzzha/service/admin"
	tgbotapi "github.com/bots-house/telegram-bot-api"
	"github.com/pkg/errors"
)

const (
	textReviewAddedSuccess   = "👍 Отзыв успешно добавлен."
	reviewTextCantBeNullable = "В отзыве должен присутствовать текст."
)

func newReviewInput(msg *tgbotapi.Message) *admin.ReviewInput {
	reviewInput := &admin.ReviewInput{}

	if msg.ForwardFrom != nil {
		reviewInput.TelegramID = msg.ForwardFrom.ID
		reviewInput.FirstName = msg.ForwardFrom.FirstName
		reviewInput.LastName = msg.ForwardFrom.LastName
		reviewInput.Username = msg.ForwardFrom.UserName
	} else {
		names := strings.Split(msg.ForwardSenderName, " ")
		if len(names) == 1 {
			reviewInput.FirstName = names[0]
		} else {
			reviewInput.FirstName = names[0]
			reviewInput.LastName = names[1]
		}
	}
	reviewInput.Text = msg.Text
	reviewInput.CreatedAt = time.Unix(int64(msg.ForwardDate), 0)
	return reviewInput
}

func (bot *Bot) onForward(ctx context.Context, msg *tgbotapi.Message) error {
	if msg.Text == "" {
		answ := bot.newAnswerMsg(msg, reviewTextCantBeNullable)
		answ.ReplyToMessageID = msg.MessageID
		return bot.send(answ)
	}

	in := newReviewInput(msg)

	err := bot.adminSrv.AddReview(ctx, getUserCtx(ctx), in)
	if err != nil {
		return errors.Wrap(err, "Add review")
	}

	answ := bot.newAnswerMsg(msg, textReviewAddedSuccess)
	answ.ReplyToMessageID = msg.MessageID
	return bot.send(answ)
}
