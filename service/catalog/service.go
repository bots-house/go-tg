package catalog

import (
	"context"

	"github.com/bots-house/birzzha/core"
	"github.com/bots-house/birzzha/pkg/storage"
	"github.com/bots-house/birzzha/pkg/tg"
	"github.com/bots-house/birzzha/store"
)

type Service struct {
	Topic    core.TopicStore
	Lot      core.LotStore
	LotTopic core.LotTopicStore
	Settings core.SettingsStore
	User     core.UserStore

	Resolver tg.Resolver
	Storage  storage.Storage
	Txier    store.Txier
}

func (srv *Service) GetTopics(ctx context.Context) (core.TopicSlice, error) {
	return srv.Topic.Query().All(ctx)
}
