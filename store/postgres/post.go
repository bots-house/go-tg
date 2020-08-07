package postgres

import (
	"context"
	"encoding/json"
	"time"

	"github.com/bots-house/birzzha/core"
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

	return &dal.Post{
		ID:                    int(post.ID),
		LotID:                 int(post.LotID),
		Text:                  post.Text,
		Buttons:               null.JSONFrom(buttons),
		ScheduledAt:           post.ScheduledAt,
		PublishedAt:           post.PublishedAt,
		DisableWebPagePreview: post.DisableWebPagePreview,
	}, nil
}

func (store *PostStore) fromRow(row *dal.Post) (*core.Post, error) {
	var buttons core.PostButtons

	if err := row.Buttons.Unmarshal(&buttons); err != nil {
		return nil, errors.Wrap(err, "unmarshal post buttons")
	}

	return &core.Post{
		ID:                    core.PostID(row.ID),
		LotID:                 core.LotID(row.LotID),
		Text:                  row.Text,
		Buttons:               buttons,
		ScheduledAt:           row.ScheduledAt,
		PublishedAt:           row.PublishedAt,
		DisableWebPagePreview: row.DisableWebPagePreview,
	}, nil
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

func (store *PostStore) Pull(ctx context.Context) (core.PostSlice, error) {
	executor := shared.GetExecutorOrDefault(ctx, store.ContextExecutor)

	rows, err := dal.Posts(
		dal.PostWhere.ScheduledAt.LTE(time.Now()),
		dal.PostWhere.PublishedAt.IsNull(),
		qm.For("UPDATE SKIP LOCKED"),
	).All(ctx, executor)
	if err != nil {
		return nil, errors.Wrap(err, "select query")
	}

	return store.fromRowSlice(rows)
}
