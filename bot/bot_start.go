package bot

import (
	"context"
	"strings"

	"github.com/pkg/errors"

	tgbotapi "github.com/bots-house/telegram-bot-api"

	"github.com/bots-house/birzzha/service/auth"
)

const (
	textStart         = "Привет! Я бот канала @birzzha. Что тебя интересует?"
	textStartLogin    = "Привет! Для авторизации на сайте нажми на кнопку ниже"
	textStartNotFound = "Кажется, ссылка устарела..."
	loginPrefix       = "login_"
)

func joinSitePath(site, path string) string {
	path = strings.TrimPrefix(path, "/")
	site = strings.TrimSuffix(site, "/")

	return site + "/" + path
}

func (bot *Bot) onStart(ctx context.Context, msg *tgbotapi.Message) error {
	if strings.HasPrefix(msg.CommandArguments(), loginPrefix) {
		return bot.onStartLogin(ctx, msg)
	}

	answ := bot.newAnswerMsg(ctx, msg, textStart)
	answ.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.InlineKeyboardButton{
				Text: "💳 Купить канал",
				LoginURL: &tgbotapi.LoginURL{
					URL: joinSitePath(bot.cfg.Site, bot.cfg.PathListChannel),
				},
			},
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.InlineKeyboardButton{
				Text: "🤑 Продать канал",
				LoginURL: &tgbotapi.LoginURL{
					URL: joinSitePath(bot.cfg.Site, bot.cfg.PathSellChannel),
				},
			},
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.InlineKeyboardButton{
				Text: "🚀 О нас",
				LoginURL: &tgbotapi.LoginURL{
					URL: bot.cfg.Site,
				},
			},
		),
	)

	return bot.send(ctx, answ)
}

func (bot *Bot) onStartLogin(ctx context.Context, msg *tgbotapi.Message) error {
	id := strings.TrimPrefix(msg.CommandArguments(), loginPrefix)

	info, err := bot.authSrv.PopLoginViaBot(ctx, id)
	if err == auth.ErrBotLoginNotFound {
		return bot.send(ctx, bot.newAnswerMsg(ctx, msg, textStartNotFound))
	} else if err != nil {
		return errors.Wrap(err, "pop login via bot")
	}

	answ := bot.newAnswerMsg(ctx, msg, textStartLogin)
	answ.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.InlineKeyboardButton{
				Text: "🔓 Авторизоватся",
				LoginURL: &tgbotapi.LoginURL{
					URL: info.CallbackURL,
				},
			},
		),
	)

	return bot.send(ctx, answ)
}
