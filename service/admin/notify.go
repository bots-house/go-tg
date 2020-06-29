package admin

import (
	"bytes"
	"log"
	"strings"
	"sync"
	"text/template"

	tgbotapi "github.com/bots-house/telegram-bot-api"
	"github.com/lithammer/dedent"
)

type Notification interface {
	NotificationTemplate() string
}

// Log it's it's Telegram channel with notifications.
type Notifications struct {
	client    *tgbotapi.BotAPI
	channelID int64

	pending chan Notification
	wg      sync.WaitGroup
}

func NewNotifications(
	client *tgbotapi.BotAPI,
	channelID int64,
) *Notifications {
	ns := &Notifications{
		client:    client,
		channelID: channelID,

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

			msg := tgbotapi.NewMessage(ns.channelID, res.String())
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
