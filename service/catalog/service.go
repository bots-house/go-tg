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

type TopicItem struct {
	*core.Topic
	Lots int
}

func (srv *Service) GetTopics(ctx context.Context) ([]*TopicItem, error) {
	topics, err := srv.Topic.Query().All(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get topics")
	}

	result, err := srv.Lot.PublishedLotsCountByTopics(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get lots count")
	}

	out := make([]*TopicItem, len(topics))
	for i, topic := range topics {
		item := &TopicItem{
			Topic: topic,
			Lots:  0,
		}
		r := result.Find(topic.ID)
		if r != nil {
			item.Lots = r.Lots
		}
		out[i] = item
	}

	return out, nil
}
