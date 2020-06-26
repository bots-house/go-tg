package bot

import (
	"context"
	"fmt"
	"strings"

	tgbotapi "github.com/bots-house/telegram-bot-api"
)

func embeddWebPagePreview(txt string, link string) string {
	return fmt.Sprintf("[‎](%s)%s", link, txt)
}

func escapeMarkdown(txt string) string {
	txt = strings.Replace(txt, "_", "\\_", -1)
	txt = strings.Replace(txt, "*", "\\*", -1)
	txt = strings.Replace(txt, "[", "\\[", -1)
	txt = strings.Replace(txt, "`", "\\`", -1)
	return txt
}

func (bot *Bot) send(ctx context.Context, s tgbotapi.Chattable) error {
	// spew.Dump(msg)
	_, err := bot.client.Send(s)
	return err
}

func (bot *Bot) newAnswerMsg(ctx context.Context, msg *tgbotapi.Message, text string) *tgbotapi.MessageConfig {
	result := tgbotapi.NewMessage(
		int64(msg.From.ID),
		text,
	)

	result.ParseMode = tgbotapi.ModeMarkdown

	return &result
}

func (bot *Bot) answerCallbackQuery(ctx context.Context, cbq *tgbotapi.CallbackQuery, text string) error {
	_, err := bot.client.AnswerCallbackQuery(tgbotapi.NewCallback(
		cbq.ID,
		text,
	))

	return err
}