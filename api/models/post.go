package models

import (
	"github.com/bots-house/birzzha/api/gen/models"
	"github.com/bots-house/birzzha/core"
	"github.com/go-openapi/swag"
)

func NewPost(in *core.Post) *models.Post {
	return &models.Post{
		ID:                    swag.Int64(int64(in.ID)),
		LotID:                 swag.Int64(int64(in.LotID)),
		Text:                  swag.String(in.Text),
		DisableWebPagePreview: swag.Bool(in.DisableWebPagePreview),
		ScheduledAt:           timeToUnix(in.ScheduledAt),
		PublishedAt:           nullTimeToUnix(in.PublishedAt),
	}
}
