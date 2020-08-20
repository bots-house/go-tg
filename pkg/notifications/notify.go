package notifications

import (
	"bytes"
	"context"
	"log"
	"strings"
	"sync"
	"text/template"

	"github.com/bots-house/birzzha/core"
	pLog "github.com/bots-house/birzzha/pkg/log"
	tgbotapi "github.com/bots-house/telegram-bot-api"
	"github.com/lithammer/dedent"
)

type Notification interface {
	NotificationTemplate() string
	ChatID() int64
}

// Log it's it's Telegram channel with notifications.
type Notifications struct {
	client *tgbotapi.BotAPI

	pending chan Notification
	wg      sync.WaitGroup
}

func New(client *tgbotapi.BotAPI) *Notifications {
	ns := &Notifications{
		client: client,

		pending: make(chan Notification, 1),
	}

	go ns.run()

	return ns
}

func (ns *Notifications) run() {
	for notify := range ns.pending {
		notify := notify

		ns.wg.Add(1)
		go func() {
			defer ns.wg.Done()

			t := dedent.Dedent(notify.NotificationTemplate())
			t = strings.TrimSpace(t)

			tmpl, err := template.New("notification").Parse(t)
			if err != nil {
				log.Printf("parse notification template failed: %v", err)
				return
			}

			res := &bytes.Buffer{}

			if err := tmpl.Execute(res, notify); err != nil {
				log.Printf("fail to execute notification template: %v", err)
				return
			}

			msg := tgbotapi.NewMessage(notify.ChatID(), res.String())
			msg.ParseMode = tgbotapi.ModeHTML
			msg.DisableWebPagePreview = true

			_, err = ns.client.Send(msg)
			if err != nil {
				log.Printf("fail to send message: %s", err)
			}
		}()
	}
}

func (ns *Notifications) Send(n Notification) {
	if ns == nil {
		return
	}
	ns.pending <- n
}

func (ns *Notifications) Close() {
	ns.wg.Wait()
	close(ns.pending)
}

type Template interface {
	Build() string
}

type UserNotification struct {
	UsrStore core.UserStore
	Notifier *Notifications
}

func (u UserNotification) Send(ctx context.Context, lot *core.Lot, tpl Template) {
	usr, err := u.UsrStore.Query().ID(lot.OwnerID).One(ctx)
	switch err {
	case nil:
		u.Notifier.Send(defaultNotification{
			chatID:   int64(usr.Telegram.ID),
			template: tpl.Build,
			Lot:      lot,
		})
	default:
		pLog.Error(ctx, "get user by owner id", "error", err, "user_id", lot.OwnerID)
	}

}

type defaultNotification struct {
	chatID   int64
	template func() string
	Lot      *core.Lot
}

func (d defaultNotification) NotificationTemplate() string {
	d.Lot.Link()
	return d.template()
}

func (d defaultNotification) ChatID() int64 {
	return d.chatID
}
