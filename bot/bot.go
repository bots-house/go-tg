package bot

import (
	"context"
	"net/url"
	"strings"

	"github.com/bots-house/birzzha/core"
	"github.com/bots-house/birzzha/pkg/log"
	"github.com/bots-house/birzzha/pkg/tg"
	"github.com/bots-house/birzzha/service/admin"
	"github.com/bots-house/birzzha/service/auth"
	tgbotapi "github.com/bots-house/telegram-bot-api"

	"github.com/pkg/errors"
)

type Config struct {
	Site            string
	PathSellChannel string
	PathListChannel string
}

type Bot struct {
	client *tgbotapi.BotAPI

	cfg      Config
	authSrv  *auth.Service
	adminSrv *admin.Service
	handler  tg.Handler
}

func New(cfg Config, client *tgbotapi.BotAPI, authSrv *auth.Service, adminSrv *admin.Service) *Bot {
	bot := &Bot{
		client:   client,
		authSrv:  authSrv,
		adminSrv: adminSrv,
		cfg:      cfg,
	}

	bot.initHandler()

	return bot
}

func (bot *Bot) Client() *tgbotapi.BotAPI {
	return bot.client
}

func NewLinkBuilder(username string) *core.BotLinkBuilder {
	return &core.BotLinkBuilder{
		Useraname: username,

		LoginPrefix:   startLoginPrefix,
		ContactPrefix: startContactPrefix,
	}
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
	sentryMiddleware := newSentryMiddleware()

	handler := sentryMiddleware(authMiddleware(tg.HandlerFunc(bot.onUpdate)))

	bot.handler = handler
}

func isForward(msg *tgbotapi.Message) bool {
	return msg.ForwardFrom != nil || msg.ForwardSenderName != ""
}

func (bot *Bot) onUpdate(ctx context.Context, update *tgbotapi.Update) error {

	if msg := update.Message; msg != nil {
		if msg.Command() == "start" {
			return bot.onStart(ctx, msg)
		}
		if isForward(msg) {
			return bot.onForward(ctx, msg)
		}
	}

	return nil
}

func (bot *Bot) HandleUpdate(ctx context.Context, update *tgbotapi.Update) error {
	return bot.handler.HandleUpdate(ctx, update)
}
