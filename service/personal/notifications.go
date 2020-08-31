package personal

import "github.com/bots-house/birzzha/core"

type NewLotNotification struct {
	User      *core.User
	Lot       *core.Lot
	channelID int64
}

func (n NewLotNotification) NotificationTemplate() string {
	return `
        🆕 <b>Новая заяка!</b>

        <a href="{{ .User.TelegramLink }}">{{ .User.Name }}</a> {{ if .User.Telegram.Username.Valid }}@{{ .User.Telegram.Username.String }}{{ end }}

        <b>№{{ .Lot.ID }}</b> <a href="{{ .Lot.Link }}">{{ .Lot.Name }}</a>

        <b>Цена:</b> {{ .Lot.Price.Current }} руб.

        #signup #user{{ .User.ID }} #lot{{ .Lot.ID }}
    `
}

func (n NewLotNotification) ChatID() int64 {
	return n.channelID
}

type NewPaymentNotification struct {
	Lot       *core.Lot
	Payment   *core.Payment
	channelID int64
}

func (n NewPaymentNotification) NotificationTemplate() string {
	return `
       💰 <b>Зачислен платеж!</b>

        <b>№{{ .Lot.ID }}</b> <a href="{{ .Lot.Link }}">{{ .Lot.Name }}</a>

        <b>Назначение</b>: {{ .Payment.Purpose.String }}
        <b>Шлюз</b>: {{ .Payment.Gateway }}
        <b>Запрошено</b>: {{ .Payment.Requested.Display }}
        <b>Оплачено</b>: {{ .Payment.Paid.Display }}
        {{if .Payment.Received }}<b>Зачислено</b>: {{ .Payment.Received.Display }} {{ end }}

        #payment #user{{ .Lot.OwnerID }} #lot{{ .Lot.ID }}
    `
}

func (n NewPaymentNotification) ChatID() int64 {
	return n.channelID
}

type CanceledLotNotification struct {
	Lot       *core.Lot
	Reason    *core.LotCanceledReason
	channelID int64
}

func (n CanceledLotNotification) NotificationTemplate() string {
	return `
        👋 <b>Лот снят с продажи!</b>

        <b>№{{ .Lot.ID }}</b> <a href="{{ .Lot.Link }}">{{ .Lot.Name }}</a>

        <b>Причина</b>: {{ .Reason.Why }}

        #cancel #user{{ .Lot.OwnerID }} #lot{{ .Lot.ID }}
    `
}

func (n CanceledLotNotification) ChatID() int64 {
	return n.channelID
}
