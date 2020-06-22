package models

import (
	"github.com/bots-house/birzzha/api/gen/models"
	"github.com/bots-house/birzzha/core"
	"github.com/go-openapi/swag"
)

func NewTopic(topic *core.Topic) *models.Topic {
	return &models.Topic{
		ID:        swag.Int64(int64(topic.ID)),
		Name:      swag.String(topic.Name),
		Slug:      swag.String(topic.Slug),
		CreatedAt: swag.Int64(topic.CreatedAt.Unix()),
	}
}

func NewTopicSlice(topics core.TopicSlice) []*models.Topic {
	result := make([]*models.Topic, len(topics))

	for i, topic := range topics {
		result[i] = NewTopic(topic)
	}

	return result
}
