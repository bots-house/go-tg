package catalog

import (
	"context"

	"github.com/bots-house/birzzha/core"
)

type Service struct {
	Topic core.TopicStore
}

func (srv *Service) GetTopics(ctx context.Context) (core.TopicSlice, error) {
	return srv.Topic.Query().All(ctx)
}
