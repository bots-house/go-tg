package admin

import (
	"github.com/bots-house/birzzha/core"
)

type userNotifyScheduledLot struct {
	Lot      *core.Lot
	Settings *core.Settings
}

func (n userNotifyScheduledLot) NotificationTemplate() string {
	return `
		📅 Заявка на размещение канала <a href="{{ .Self.Lot.Link }}">{{ .Self.Lot.Name }}</a> прошла модерацию и уже доступна <a href="{{ lotSiteURL .Self.Lot.ID }}">на сайте</a>. 

		В <a href="{{ .Settings.Channel.Link }}">канале</a> пост будет размещен в <b>{{ mskTime .Self.Lot.ScheduledAt.Time }}</b>.

		#лот{{ .Self.Lot.ID }}
	`
}
