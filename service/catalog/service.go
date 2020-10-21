package catalog

import (
	"context"

	"github.com/bots-house/birzzha/core"
	"github.com/bots-house/birzzha/pkg/stat"
	"github.com/bots-house/birzzha/pkg/storage"
	"github.com/bots-house/birzzha/pkg/tg"
	"github.com/bots-house/birzzha/store"
	"github.com/pkg/errors"
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

type Topics struct {
	Lots   int
	Topics core.TopicSlice
}

func (srv *Service) GetTopics(ctx context.Context) (*Topics, error) {
	lots, err := srv.Lot.Query().Statuses(core.LotStatusPublished).Count(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get lots count")
	}

	topics, err := srv.Topic.Query().All(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get topics")
	}

	return &Topics{
		Lots:   lots,
		Topics: topics,
	}, nil
}
