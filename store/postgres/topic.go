package postgres

import (
	"context"
	"database/sql"

	"github.com/bots-house/birzzha/core"
	"github.com/bots-house/birzzha/store/postgres/dal"
	"github.com/bots-house/birzzha/store/postgres/shared"
	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type TopicStore struct {
	boil.ContextExecutor
}

func (store *TopicStore) toRow(topic *core.Topic) *dal.Topic {
	return &dal.Topic{
		ID:        int(topic.ID),
		Name:      topic.Name,
		Slug:      topic.Slug,
		CreatedAt: topic.CreatedAt,
	}
}

func (store *TopicStore) fromRow(row *dal.Topic) *core.Topic {
	return &core.Topic{
		ID:        core.TopicID(row.ID),
		Name:      row.Name,
		Slug:      row.Slug,
		CreatedAt: row.CreatedAt,
	}
}

func (store *TopicStore) Add(ctx context.Context, topic *core.Topic) error {
	row := store.toRow(topic)
	if err := row.Insert(ctx, shared.GetExecutorOrDefault(ctx, store.ContextExecutor), boil.Infer()); err != nil {
		return errors.Wrap(err, "insert query")
	}
	*topic = *store.fromRow(row)
	return nil
}

func (store *TopicStore) Update(ctx context.Context, topic *core.Topic) error {
	row := store.toRow(topic)
	n, err := row.Update(ctx, shared.GetExecutorOrDefault(ctx, store.ContextExecutor), boil.Infer())
	if err != nil {
		return errors.Wrap(err, "update query")
	}
	if n == 0 {
		return core.ErrTopicNotFound
	}
	return nil
}

func (store *TopicStore) Query() core.TopicStoreQuery {
	return &TopicStoreQuery{store: store}
}

type TopicStoreQuery struct {
	mods  []qm.QueryMod
	store *TopicStore
}

func (usq *TopicStoreQuery) ID(ids ...core.TopicID) core.TopicStoreQuery {
	idsInt := make([]int, len(ids))
	for i, id := range ids {
		idsInt[i] = int(id)
	}

	usq.mods = append(usq.mods, dal.TopicWhere.ID.IN(idsInt))
	return usq
}

func (usq *TopicStoreQuery) One(ctx context.Context) (*core.Topic, error) {
	executor := shared.GetExecutorOrDefault(ctx, usq.store.ContextExecutor)

	row, err := dal.Topics(usq.mods...).One(ctx, executor)
	if err == sql.ErrNoRows {
		return nil, core.ErrTopicNotFound
	} else if err != nil {
		return nil, err
	}

	return usq.store.fromRow(row), nil
}

func (usq *TopicStoreQuery) All(ctx context.Context) (core.TopicSlice, error) {
	executor := shared.GetExecutorOrDefault(ctx, usq.store.ContextExecutor)

	rows, err := dal.Topics(usq.mods...).All(ctx, executor)
	if err != nil {
		return nil, err
	}

	result := make(core.TopicSlice, len(rows))
	for i, row := range rows {
		result[i] = usq.store.fromRow(row)
	}

	return result, nil
}
