package bot

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/pkg/errors"

	tgbotapi "github.com/bots-house/telegram-bot-api"

	"github.com/bots-house/birzzha/pkg/log"
	"github.com/bots-house/birzzha/service/auth"
)

const (
	textStart              = "Привет! Я бот канала @birzzha. Что тебя интересует?"
	textStartLogin         = "Привет! Для авторизации на сайте нажми на кнопку ниже"
	textStartLoginNotFound = "Кажется, ссылка устарела..."

	textStartNotAdmin = "😕Тебе сюда нельзя!"

	textStartContactSuccess        = "Знаю такого, вот [ссылка](tg://user?id=%d) на профиль."
	textStartContactQueryChatError = "Ошибка при получении данных пользователя: `%v`"
	startLoginPrefix               = "login_"
	startContactPrefix             = "contact_"
)

func joinSitePath(site, path string) string {
	path = strings.TrimPrefix(path, "/")
	site = strings.TrimSuffix(site, "/")

	return site + "/" + path
}

func (bot *Bot) onStart(ctx context.Context, msg *tgbotapi.Message) error {
	args := msg.CommandArguments()

	switch {
	case strings.HasPrefix(args, startLoginPrefix):
		return bot.onStartLogin(ctx, msg)
	case strings.HasPrefix(args, startContactPrefix):
		return bot.onStartContact(ctx, msg)
	default:
		return bot.onStartDefault(ctx, msg)
	}
}

func (bot *Bot) onStartDefault(_ context.Context, msg *tgbotapi.Message) error {
	answ := bot.newAnswerMsg(msg, textStart)
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

	return bot.send(answ)
}

func (bot *Bot) onStartContact(ctx context.Context, msg *tgbotapi.Message) error {
	user := getUserCtx(ctx)

	if !user.IsAdmin {
		return bot.send(bot.newReplyMsg(msg, textStartNotAdmin))
	}

	idStr := strings.TrimPrefix(msg.CommandArguments(), startContactPrefix)

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return bot.send(bot.newReplyMsg(msg, textStartLoginNotFound))
	}

	chat, err := bot.client.GetChat(tgbotapi.ChatConfig{
		ChatID: int64(id),
	})

	if err != nil {
		return bot.send(bot.newReplyMsg(msg, fmt.Sprintf(textStartContactQueryChatError, err)))
	}

	answ := bot.newReplyMsg(msg, fmt.Sprintf(textStartContactSuccess, chat.ID))

	return bot.send(answ)
}

func (bot *Bot) onStartLogin(ctx context.Context, msg *tgbotapi.Message) error {
	id := strings.TrimPrefix(msg.CommandArguments(), startLoginPrefix)

	info, err := bot.authSrv.PopLoginViaBot(ctx, id)
	if err == auth.ErrBotLoginNotFound {
		return bot.send(bot.newAnswerMsg(msg, textStartLoginNotFound))
	} else if err != nil {
		return errors.Wrap(err, "pop login via bot")
	}

	answ := bot.newAnswerMsg(msg, textStartLogin)
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

	if err := bot.send(answ); err != nil {
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

			answ := bot.newAnswerMsg(msg, textStartLogin)
			answ.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.InlineKeyboardButton{
						Text: "🔓 Авторизоватся",
						URL:  &cb,
					},
				),
			)

			return bot.send(answ)
		}

		return err
	}

	return nil
}

func isBotDomainInvalidError(err error) bool {
	tgerr, ok := err.(*tgbotapi.Error)
	return ok && strings.Contains(tgerr.Message, "BOT_DOMAIN_INVALID")
}
