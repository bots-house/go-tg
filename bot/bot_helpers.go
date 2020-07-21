package bot

import (
	tgbotapi "github.com/bots-house/telegram-bot-api"
)

func (bot *Bot) send(s tgbotapi.Chattable) error {
	// spew.Dump(msg)
	_, err := bot.client.Send(s)
	return err
}

func (bot *Bot) newReplyMsg(msg *tgbotapi.Message, text string) *tgbotapi.MessageConfig {
	result := bot.newAnswerMsg(msg, text)

	result.ReplyToMessageID = msg.MessageID

	return result
}

func (bot *Bot) newAnswerMsg(msg *tgbotapi.Message, text string) *tgbotapi.MessageConfig {
	result := tgbotapi.NewMessage(
		int64(msg.From.ID),
		text,
	)

	result.ParseMode = tgbotapi.ModeMarkdown

	return &result
}
