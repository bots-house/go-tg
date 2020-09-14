package personal

import "github.com/bots-house/birzzha/core"

type adminNewLotNotification struct {
	User *core.User
	Lot  *core.Lot
}

func (n adminNewLotNotification) NotificationTemplate() string {
	return `
        🆕 <b>Новая заяка!</b>

        <a href="{{ .Self.User.TelegramLink }}">{{ .Self.User.Name }}</a> {{ if .Self.User.Telegram.Username.Valid }}@{{ .Self.User.Telegram.Username.String }}{{ end }}

        <b>№{{ .Self.Lot.ID }}</b> <a href="{{ .Self.Lot.Link }}">{{ .Self.Lot.Name }}</a>

        <b>Цена:</b> {{ .Self.Lot.Price.Current }} руб.

        #signup #user{{ .Self.User.ID }} #lot{{ .Self.Lot.ID }}
    `
}

type adminNewPaymentNotification struct {
	Lot     *core.Lot
	Payment *core.Payment
	Coupon  *core.Coupon
}

func (n adminNewPaymentNotification) NotificationTemplate() string {
	return `
       💰 <b>Зачислен платеж!</b>

        <b>№{{ .Self.Lot.ID }}</b> <a href="{{ .Self.Lot.Link }}">{{ .Self.Lot.Name }}</a>

        <b>Назначение</b>: {{ .Self.Payment.Purpose.String }}
        <b>Шлюз</b>: {{ .Self.Payment.Gateway }}
        {{ if .Self.Payment.Requested }}<b>Запрошено</b>: {{ .Self.Payment.Requested.Display }}{{end }}
        {{ if .Self.Payment.Paid }}<b>Оплачено</b>: {{ .Self.Payment.Paid.Display }}{{- end }}
        {{ if .Self.Payment.Received }}<b>Зачислено</b>: {{ .Self.Payment.Received.Display }} {{- end }}
        {{ if .Self.Coupon }}<b>Купон</b>: {{ .Self.Coupon.Code }} (-{{ .Self.Coupon.Discount | percent }}) {{end }}

        #payment #user{{ .Self.Lot.OwnerID }} #lot{{ .Self.Lot.ID }} {{ if .Self.Coupon }}#coupon{{ .Self.Coupon.ID }}{{ end }}
    `
}

type userNewPaymentNotification struct {
	Lot *core.Lot
}

func (n userNewPaymentNotification) NotificationTemplate() string {
	return `
        💸 Платеж по заявке на размещение канала <a href="{{ .Self.Lot.Link }}">{{ .Self.Lot.Name }}</a> зачислен! Ожидайте модерации.

        #лот{{ .Self.Lot.ID }}
    `
}

type adminCanceledLotNotification struct {
	Lot    *core.Lot
	Reason *core.LotCanceledReason
}

func (n adminCanceledLotNotification) NotificationTemplate() string {
	return `
        👋 <b>Лот снят с продажи!</b>

        <b>№{{ .Self.Lot.ID }}</b> <a href="{{ .Self.Lot.Link }}">{{ .Self.Lot.Name }}</a>

        <b>Причина</b>: {{ .Self.Reason.Why }}

        #cancel #user{{ .Self.Lot.OwnerID }} #lot{{ .Self.Lot.ID }}
    `
}
