package auth

import (
	"github.com/bots-house/birzzha/core"
)

type NewUserNotification struct {
	User *core.User
}

func (n NewUserNotification) NotificationTemplate() string {
	return `
		🆕 *Новый пользователь!* 
		
		[{{ .User.Name }}]({{ .User.TelegramLink }}) {{ if .User.Telegram.Username.Valid }}@{{ .User.Telegram.Username.String }}{{ end }} 
		
		**Способ регистрации:** _{{ if eq .User.JoinedFrom  "bot" }}бот{{ else }}cайт{{ end }}_.

		#user{{ .User.ID }}
	`
}
