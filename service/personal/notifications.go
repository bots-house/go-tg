package personal

import "github.com/bots-house/birzzha/core"

type NewLotNotification struct {
	User *core.User
	Lot  *core.Lot
}

func (n NewLotNotification) NotificationTemplate() string {
	return `
        🆕 <b>Новая заяка!</b>

        <a href="{{ .User.TelegramLink }}">{{ .User.Name }}</a> {{ if .User.Telegram.Username.Valid }}@{{ .User.Telegram.Username.String }}{{ end }}

        <b>№{{ .Lot.ID }}</b> <a href="{{ .Lot.Link }}">{{ .Lot.Name }}</a>

        <b>Цена:</b> {{ .Lot.Price.Current }} руб.

        #user{{ .User.ID }} #lot{{ .Lot.ID }}
    `
}

type NewPaymentNotification struct {
	Lot     *core.Lot
	Payment *core.Payment
}

func (n NewPaymentNotification) NotificationTemplate() string {
	return `
       💰 <b>Зачислен платеж!</b>

        <b>№{{ .Lot.ID }}</b> <a href="{{ .Lot.Link }}">{{ .Lot.Name }}</a>

        <b>Назначение</b>: {{ .Payment.Purpose.String }}
        <b>Шлюз</b>: {{ .Payment.Gateway }}
        <b>Запрошено</b>: {{ .Payment.Requested.Display }}
        <b>Оплачено</b>: {{ .Payment.Paid.Display }}
        <b>Зачислено</b>: {{ .Payment.Received.Display }}


        #user{{ .Lot.OwnerID }} #lot{{ .Lot.ID }}
    `
}
