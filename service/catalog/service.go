package catalog

import (
	"context"

	"github.com/bots-house/birzzha/core"
	"github.com/bots-house/birzzha/pkg/stat"
	"github.com/bots-house/birzzha/pkg/storage"
	"github.com/bots-house/birzzha/pkg/tg"
	"github.com/bots-house/birzzha/store"
)

type Service struct {
	Topic        core.TopicStore
	Lot          core.LotStore
	LotTopic     core.LotTopicStore
	Settings     core.SettingsStore
	User         core.UserStore
	Favorite     core.FavoriteStore
	LotFile      core.LotFileStore
	TelegramStat stat.Telegram

	Resolver tg.Resolver
	Storage  storage.Storage
	Txier    store.Txier
}

func (srv *Service) GetTopics(ctx context.Context) (core.TopicSlice, error) {
	return srv.Topic.Query().All(ctx)
}
