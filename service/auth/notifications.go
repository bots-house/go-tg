package auth

import (
	"github.com/bots-house/birzzha/core"
)

type NewUserNotification struct {
	User      *core.User
	channelID int64
}

func (n NewUserNotification) NotificationTemplate() string {
	return `
        👤 <b>Новый пользователь!</b>

		<a href="{{ .User.TelegramLink }}">{{ .User.Name }}</a> {{ if .User.Telegram.Username.Valid }}@{{ .User.Telegram.Username.String }}{{ end }}

		<b>Способ регистрации:</b> <i>{{ if eq .User.JoinedFrom  "bot" }}бот{{ else }}cайт{{ end }}</i>.

		#user{{ .User.ID }}
	`
}

func (n NewUserNotification) ChatID() int64 {
	return n.channelID
}
