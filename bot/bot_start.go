package bot

import (
	"context"
	"fmt"
	"strings"

	"github.com/pkg/errors"

	"github.com/bots-house/birzzha/service/auth"
	tgbotapi "github.com/bots-house/telegram-bot-api"
)

const (
	textStart         = "Привет! Я бот @birzzha."
	textStartNotFound = "Кажется, ссылка устарела..."
	loginPrefix       = "login_"
)

func (bot *Bot) onStart(ctx context.Context, msg *tgbotapi.Message) error {
	if strings.HasPrefix(msg.CommandArguments(), loginPrefix) {
		return bot.onStartLogin(ctx, msg)
	}

	return nil
}

func (bot *Bot) onStartLogin(ctx context.Context, msg *tgbotapi.Message) error {
	id := strings.TrimPrefix(msg.CommandArguments(), loginPrefix)

	info, err := bot.authSrv.PopLoginViaBot(ctx, id)
	if err == auth.ErrBotLoginNotFound {
		return bot.send(ctx, bot.newAnswerMsg(ctx, msg, textStartNotFound))
	} else if err != nil {
		return errors.Wrap(err, "pop login via bot")
	}

	fmt.Println(info.CallbackURL)

	answ := bot.newAnswerMsg(ctx, msg, "Для авторизации нажмите на кнопку 👇")
	answ.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.InlineKeyboardButton{
				Text: "Авторизоватся",
				LoginURL: &tgbotapi.LoginURL{
					URL: info.CallbackURL,
				},
			},
		),
	)

	return bot.send(ctx, answ)
}
