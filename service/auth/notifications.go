package auth

import (
	"github.com/bots-house/birzzha/core"
)

type newUserNotification struct {
	User *core.User
}

func (n newUserNotification) NotificationTemplate() string {
	return `
        👤 <b>Новый пользователь!</b>

		<a href="{{ .Self.User.TelegramLink }}">{{ .Self.User.Name }}</a> {{ if .Self.User.Telegram.Username.Valid }}@{{ .Self.User.Telegram.Username.String }}{{ end }}

		<b>Способ регистрации:</b> <i>{{ if eq .Self.User.JoinedFrom  "bot" }}бот{{ else }}cайт{{ end }}</i>.

		#user{{ .Self.User.ID }}
	`
}
