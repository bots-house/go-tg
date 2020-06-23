package catalog

import (
	"context"

	"github.com/bots-house/birzzha/core"
	"github.com/bots-house/birzzha/pkg/tg"
)

type Service struct {
	Topic    core.TopicStore
	Resolver tg.Resolver
}

func (srv *Service) GetTopics(ctx context.Context) (core.TopicSlice, error) {
	return srv.Topic.Query().All(ctx)
}
