package models

import (
	"github.com/bots-house/birzzha/api/gen/models"
	"github.com/bots-house/birzzha/core"
	"github.com/bots-house/birzzha/service/admin"
	"github.com/bots-house/birzzha/service/catalog"
	"github.com/go-openapi/swag"
)

func NewTopicItem(topic *catalog.TopicItem) *models.TopicItem {
	return &models.TopicItem{
		ID:        swag.Int64(int64(topic.ID)),
		Name:      swag.String(topic.Name),
		Slug:      swag.String(topic.Slug),
		CreatedAt: swag.Int64(topic.CreatedAt.Unix()),
		Lots:      swag.Int64(int64(topic.Lots)),
	}
}

func NewTopicIDSlice(ids []core.TopicID) []int64 {
	result := make([]int64, len(ids))
	for i, v := range ids {
		result[i] = int64(v)
	}

	return result
}

func NewTopicItemSlice(topics []*catalog.TopicItem) []*models.TopicItem {
	result := make([]*models.TopicItem, len(topics))
	for i, topic := range topics {
		result[i] = NewTopicItem(topic)
	}

	return result
}

func NewAdminFullTopic(topic *admin.FullTopic) *models.AdminTopic {
	return &models.AdminTopic{
		ID:        swag.Int64(int64(topic.ID)),
		Name:      swag.String(topic.Name),
		Slug:      swag.String(topic.Slug),
		CreatedAt: swag.Int64(topic.CreatedAt.Unix()),
		Lots:      swag.Int64(int64(topic.Lots)),
	}
}

func NewAdminFullTopicSlice(topics []*admin.FullTopic) []*models.AdminTopic {
	result := make([]*models.AdminTopic, len(topics))
	for i, topic := range topics {
		result[i] = NewAdminFullTopic(topic)
	}
	return result
}
