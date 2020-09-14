package notifications

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"
	"sync"
	"text/template"
	"time"

	"github.com/bots-house/birzzha/core"
	tgbotapi "github.com/bots-house/telegram-bot-api"
	"github.com/goodsign/monday"
	"github.com/lithammer/dedent"
)

// Log it's it's Telegram channel with notifications.
type Notifications struct {
	client      *tgbotapi.BotAPI
	adminChatID int64
	pending     chan Notification
	settings    core.SettingsStore
	user        core.UserStore
	wg          sync.WaitGroup
	funcs       template.FuncMap
}

type Paths struct {
	Site string
	Lots string
}

func New(client *tgbotapi.BotAPI, user core.UserStore, settings core.SettingsStore, adminChatID int64, paths Paths) *Notifications {
	ns := &Notifications{
		client: client,

		adminChatID: adminChatID,
		pending:     make(chan Notification, 10),
		settings:    settings,
		user:        user,
		funcs: template.FuncMap{
			"lotSiteURL": func(id core.LotID) string {
				return paths.Site + paths.Lots + "/" + strconv.Itoa(int(id)) + "?from=bot"
			},
			"mskTime": func(v time.Time) string {
				loc, err := time.LoadLocation("Europe/Moscow")
				if err == nil {
					v = v.In(loc)
				}
				return monday.Format(v, "02 Jan 15:04 MST", monday.LocaleRuRU)
			},
			"percent": func(v float64) string {
				return fmt.Sprintf("%.2f%%", v*100)
			},
		},
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

			ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
			defer cancel()

			settings, err := ns.settings.Get(ctx)
			if err != nil {
				log.Printf("get settings failed: %v", err)
				return
			}

			t := dedent.Dedent(notify.NotificationTemplate())
			t = strings.TrimSpace(t)

			tmpl, err := template.New("notification").
				Funcs(ns.funcs).
				Parse(t)

			if err != nil {
				log.Printf("parse notification template failed: %v", err)
				return
			}

			res := &bytes.Buffer{}

			vars := struct {
				Self     Notification
				Settings *core.Settings
			}{
				Settings: settings,
			}

			if n, ok := notify.(userNotificationWrapper); ok {
				vars.Self = n.Notification
			} else {
				vars.Self = notify
			}

			if err := tmpl.Execute(res, vars); err != nil {
				log.Printf("fail to execute notification template: %v", err)
				return
			}

			var chatID int64
			if n, ok := notify.(userNotificationWrapper); ok {
				user, err := ns.user.Query().ID(n.UserID()).One(ctx)
				if err != nil {
					log.Printf("fail to resolve user chat id: %v", err)
					return
				}

				chatID = int64(user.Telegram.ID)
			} else {
				chatID = ns.adminChatID
			}

			msg := tgbotapi.NewMessage(chatID, res.String())
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

func (ns *Notifications) SendUser(id core.UserID, n Notification) {
	ns.Send(userNotificationWrapper{
		Notification: n,
		userID:       id,
	})
}

func (ns *Notifications) Close() {
	ns.wg.Wait()
	close(ns.pending)
}
