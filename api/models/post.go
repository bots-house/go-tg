package models

import (
	"github.com/bots-house/birzzha/api/gen/models"
	"github.com/bots-house/birzzha/core"
	"github.com/bots-house/birzzha/pkg/storage"
	"github.com/bots-house/birzzha/service/admin"
	"github.com/go-openapi/swag"
)

func NewPost(in *core.Post) *models.Post {
	return &models.Post{
		ID:                    swag.Int64(int64(in.ID)),
		LotID:                 swag.Int64(int64(in.LotID)),
		Text:                  swag.String(in.Text),
		Title:                 nullStringToString(in.Title),
		DisableWebPagePreview: swag.Bool(in.DisableWebPagePreview),
		ScheduledAt:           timeToUnix(in.ScheduledAt),
		PublishedAt:           nullTimeToUnix(in.PublishedAt),
	}
}

func NewPostItem(s storage.Storage, in *admin.PostItem) *models.PostItem {
	item := &models.PostItem{
		ID:                    swag.Int64(int64(in.ID)),
		Text:                  swag.String(in.Text),
		Title:                 nullStringToString(in.Title),
		DisableWebPagePreview: swag.Bool(in.DisableWebPagePreview),
		ScheduledAt:           timeToUnix(in.ScheduledAt),
		PublishedAt:           nullTimeToUnix(in.PublishedAt),
	}

	if in.Lot != nil {
		item.Lot = &models.PostLot{
			ID:       swag.Int64(int64(in.LotID)),
			Name:     swag.String(in.Lot.Name),
			Username: nullStringToString(in.Lot.Username),
			JoinLink: nullStringToString(in.Lot.JoinLink),
		}

		if in.Lot.Avatar.Valid {
			item.Lot.Avatar = swag.String(s.PublicURL(in.Lot.Avatar.String))
		}
	}

	return item
}

func NewFullPost(s storage.Storage, in *admin.FullPost) *models.PostListItem {
	items := make([]*models.PostItem, in.Total)
	for i, item := range in.Items {
		items[i] = NewPostItem(s, item)
	}

	return &models.PostListItem{
		Total: swag.Int64(int64(in.Total)),
		Items: items,
	}
}
