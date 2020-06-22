package bot

import (
	"context"
	"net/url"
	"strings"

	"github.com/bots-house/birzzha/pkg/log"
	"github.com/bots-house/birzzha/pkg/tg"
	"github.com/bots-house/birzzha/service/auth"
	"github.com/davecgh/go-spew/spew"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

	"github.com/pkg/errors"
)

type Bot struct {
	client *tgbotapi.BotAPI

	authSrv *auth.Service
	handler tg.Handler
}

func New(client *tgbotapi.BotAPI, authSrv *auth.Service) *Bot {
	bot := &Bot{
		client:  client,
		authSrv: authSrv,
	}

	bot.initHandler()

	return bot
}

func (bot *Bot) SetWebhookIfNeed(ctx context.Context, domain string, path string) error {
	webhook, err := bot.client.GetWebhookInfo()
	if err != nil {
		return errors.Wrap(err, "get webhook info")
	}

	path = strings.TrimPrefix(path, "/")

	endpoint := strings.Join([]string{domain, path}, "/")

	if webhook.URL != endpoint {
		u, err := url.Parse(endpoint)
		if err != nil {
			return errors.Wrap(err, "invalid provided webhook url")
		}

		log.Info(ctx, "update bot webhook", "old", webhook.URL, "new", u.String())
		if _, err := bot.client.SetWebhook(tgbotapi.WebhookConfig{
			URL:            u,
			MaxConnections: 40,
		}); err != nil {
			return errors.Wrap(err, "update webhook")
		}
	}

	return nil
}

func (bot *Bot) initHandler() {
	authMiddleware := newAuthMiddleware(bot.authSrv)

	handler := authMiddleware(tg.HandlerFunc(bot.onUpdate))

	bot.handler = handler
}

func (bot *Bot) onUpdate(ctx context.Context, update *tgbotapi.Update) error {
	spew.Dump(update)

	return nil
}

func (bot *Bot) HandleUpdate(ctx context.Context, update *tgbotapi.Update) error {
	return bot.handler.HandleUpdate(ctx, update)
}
