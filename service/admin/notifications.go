package admin

import (
	"strconv"

	"github.com/bots-house/birzzha/core"
)

type DeclineNotification string

func (d DeclineNotification) Build() string {
	return `
        🙅 <b>Лот <a href="{{ .Lot.Link }}">{{ .Lot.Name }}</a> (#{{ .Lot.ID }})</b> не прошел модерацию и был отклонен модератором.

Причина: {{ .Lot.DeclineReason.String }}
    `
}

type CreatePostNotification struct {
	postTgID core.PostID
}

func (c CreatePostNotification) Build() string {
	return `
        👌 <b>Лот <a href="{{ .Lot.Link }}">{{ .Lot.Name }}</a> (#{{ .Lot.ID }})</b> опубликован в канале

https://t.me/birzzha/` + strconv.Itoa(int(c.postTgID))
}
