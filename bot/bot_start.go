package bot

import (
	"context"
	"net/url"
	"strings"

	"github.com/pkg/errors"

	tgbotapi "github.com/bots-house/telegram-bot-api"

	"github.com/bots-house/birzzha/pkg/log"
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

	if err := bot.send(ctx, answ); err != nil {
		if isBotDomainInvalidError(err) {
			log.Warn(ctx, "user fallback for domain invalid", "callback_url", info.CallbackURL)
			callbackURL, err := url.Parse(info.CallbackURL)
			if err != nil {
				return errors.Wrap(err, "parse callback url")
			}

			query := callbackURL.Query()
			user := getUserCtx(ctx)

			vs := bot.authSrv.GetLoginWidgetInfo(ctx, user)

			for k := range vs {
				query.Set(k, vs.Get(k))
			}

			callbackURL.RawQuery = query.Encode()

			cb := callbackURL.String()

			answ := bot.newAnswerMsg(ctx, msg, textStartLogin)
			answ.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.InlineKeyboardButton{
						Text: "🔓 Авторизоватся",
						URL:  &cb,
					},
				),
			)

			return bot.send(ctx, answ)
		}

		return err
	}

	return nil
}

func isBotDomainInvalidError(err error) bool {
	tgerr, ok := err.(*tgbotapi.Error)
	return ok && strings.Contains(tgerr.Message, "BOT_DOMAIN_INVALID")
}
