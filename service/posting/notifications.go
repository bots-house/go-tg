package posting

import "github.com/bots-house/birzzha/core"

type userLotPublishedNotification struct {
	PostID int
	Lot    *core.Lot
}

func (l userLotPublishedNotification) NotificationTemplate() string {
	return `
		👀 Пост о продаже канала <a href="{{ .Self.Lot.Link }}">{{ .Self.Lot.Name }}</a> опубликован <a href="{{ .Settings.Channel.PostLink .Self.PostID }}">в канале</a>. 

		#лот{{ .Self.Lot.ID }}
	`
}
