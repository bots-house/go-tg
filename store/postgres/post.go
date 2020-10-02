package postgres

import (
	"context"
	"database/sql"
	"encoding/json"
	"time"

	"github.com/bots-house/birzzha/core"
	"github.com/bots-house/birzzha/store"
	"github.com/bots-house/birzzha/store/postgres/dal"
	"github.com/bots-house/birzzha/store/postgres/shared"
	"github.com/pkg/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type PostStore struct {
	boil.ContextExecutor
}

func (store *PostStore) toRow(post *core.Post) (*dal.Post, error) {
	buttons, err := json.Marshal(post.Buttons)
	if err != nil {
		return nil, errors.Wrap(err, "marshal post buttons")
	}

	row := &dal.Post{
		ID:                    int(post.ID),
		Text:                  post.Text,
		Buttons:               null.JSONFrom(buttons),
		Title:                 post.Title,
		ScheduledAt:           post.ScheduledAt,
		PublishedAt:           post.PublishedAt,
		DisableWebPagePreview: post.DisableWebPagePreview,
		Status:                post.Status.String(),
		MessageID:             post.MessageID,
	}
	if post.LotID != 0 {
		row.LotID = null.IntFrom(int(post.LotID))
	}

	return row, nil
}

func (store *PostStore) fromRow(row *dal.Post) (*core.Post, error) {
	var buttons core.PostButtons

	if err := row.Buttons.Unmarshal(&buttons); err != nil {
		return nil, errors.Wrap(err, "unmarshal post buttons")
	}

	status, err := core.ParsePostStatus(row.Status)
	if err != nil {
		return nil, errors.Wrap(err, "parse status")
	}

	post := &core.Post{
		ID:                    core.PostID(row.ID),
		Text:                  row.Text,
		Title:                 row.Title,
		Buttons:               buttons,
		ScheduledAt:           row.ScheduledAt,
		PublishedAt:           row.PublishedAt,
		DisableWebPagePreview: row.DisableWebPagePreview,
		Status:                status,
		MessageID:             row.MessageID,
	}

	if row.LotID.Valid {
		post.LotID = core.LotID(row.LotID.Int)
	}

	return post, nil
}

func (store *PostStore) fromRowSlice(rows dal.PostSlice) (core.PostSlice, error) {
	out := make(core.PostSlice, len(rows))
	for i, row := range rows {
		var err error
		out[i], err = store.fromRow(row)
		if err != nil {
			return nil, err
		}
	}
	return out, nil
}

func (store *PostStore) Add(ctx context.Context, post *core.Post) error {
	row, err := store.toRow(post)
	if err != nil {
		return err
	}

	if err := row.Insert(ctx, shared.GetExecutorOrDefault(ctx, store.ContextExecutor), boil.Infer()); err != nil {
		return errors.Wrap(err, "insert query")
	}

	p, err := store.fromRow(row)
	if err != nil {
		return err
	}

	*post = *p
	return nil
}

func (store *PostStore) Update(ctx context.Context, post *core.Post) error {
	row, err := store.toRow(post)
	if err != nil {
		return err
	}

	n, err := row.Update(ctx, shared.GetExecutorOrDefault(ctx, store.ContextExecutor), boil.Infer())
	if err != nil {
		return errors.Wrap(err, "update query")
	}
	if n == 0 {
		return core.ErrPostNotFound
	}

	return nil
}

func (store *PostStore) Delete(ctx context.Context, id core.PostID) error {
	rows, err := (&dal.Post{ID: int(id)}).Delete(ctx, shared.GetExecutorOrDefault(ctx, store.ContextExecutor))
	if err != nil {
		return errors.Wrap(err, "delete query")
	}

	if rows == 0 {
		return core.ErrPostNotFound
	}
	return nil
}

func (store *PostStore) Pull(ctx context.Context) (core.PostSlice, error) {
	executor := shared.GetExecutorOrDefault(ctx, store.ContextExecutor)

	rows, err := dal.Posts(
		dal.PostWhere.ScheduledAt.LTE(time.Now()),
		dal.PostWhere.Status.EQ(dal.PostStatusScheduled),
		qm.For("UPDATE SKIP LOCKED"),
	).All(ctx, executor)
	if err != nil {
		return nil, errors.Wrap(err, "select query")
	}

	return store.fromRowSlice(rows)
}

func (store *PostStore) Query() core.PostStoreQuery {
	return &PostStoreQuery{store: store}
}

type PostStoreQuery struct {
	mods  []qm.QueryMod
	store *PostStore
}

func (psq *PostStoreQuery) ID(ids ...core.PostID) core.PostStoreQuery {
	idsInt := make([]int, len(ids))
	for i, id := range ids {
		idsInt[i] = int(id)
	}

	psq.mods = append(psq.mods, dal.PostWhere.ID.IN(idsInt))
	return psq
}

func (psq *PostStoreQuery) SortBy(field core.PostField, typ store.SortType) core.PostStoreQuery {
	var orderBy string

	if field == core.PostFieldScheduledAt {
		orderBy = dal.PostColumns.ScheduledAt
	}

	orderBy += store.SortTypeString(typ)

	psq.mods = append(psq.mods, qm.OrderBy(orderBy))
	return psq
}

func (psq *PostStoreQuery) Statuses(statuses ...core.PostStatus) core.PostStoreQuery {
	vs := make([]string, len(statuses))

	for i, status := range statuses {
		vs[i] = status.String()
	}

	psq.mods = append(psq.mods, dal.PostWhere.Status.IN(vs))
	return psq
}

func (psq *PostStoreQuery) Offset(v int) core.PostStoreQuery {
	psq.mods = append(psq.mods, qm.Offset(v))
	return psq
}

func (psq *PostStoreQuery) Limit(v int) core.PostStoreQuery {
	psq.mods = append(psq.mods, qm.Limit(v))
	return psq
}

func (psq *PostStoreQuery) One(ctx context.Context) (*core.Post, error) {
	executor := shared.GetExecutorOrDefault(ctx, psq.store.ContextExecutor)

	row, err := dal.Posts(psq.mods...).One(ctx, executor)
	if err == sql.ErrNoRows {
		return nil, core.ErrPostNotFound
	} else if err != nil {
		return nil, err
	}

	return psq.store.fromRow(row)
}

func (psq *PostStoreQuery) All(ctx context.Context) (core.PostSlice, error) {
	executor := shared.GetExecutorOrDefault(ctx, psq.store.ContextExecutor)
	rows, err := dal.Posts(psq.mods...).All(ctx, executor)
	if err != nil {
		return nil, err
	}

	return psq.store.fromRowSlice(rows)
}

func (psq *PostStoreQuery) Count(ctx context.Context) (int, error) {
	executor := shared.GetExecutorOrDefault(ctx, psq.store.ContextExecutor)

	count, err := dal.Posts(psq.mods...).Count(ctx, executor)

	return int(count), err
}
