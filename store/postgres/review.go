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

type ReviewStore struct {
	boil.ContextExecutor
}

func (store *ReviewStore) toRow(review *core.Review) *dal.Review {
	return &dal.Review{
		ID:         int(review.ID),
		TelegramID: review.User.TelegramID,
		FirstName:  review.User.FirstName,
		LastName:   review.User.LastName,
		Username:   review.User.Username,
		Avatar:     review.User.Avatar,
		Text:       review.Text,
		CreatedAt:  review.CreatedAt,
	}
}

func (store *ReviewStore) fromRow(row *dal.Review) *core.Review {
	return &core.Review{
		ID: core.ReviewID(row.ID),
		User: core.ReviewUser{
			TelegramID: row.TelegramID,
			FirstName:  row.FirstName,
			LastName:   row.LastName,
			Username:   row.Username,
			Avatar:     row.Avatar,
		},
		Text:      row.Text,
		CreatedAt: row.CreatedAt,
	}
}

func (store *ReviewStore) fromRowSlice(rows dal.ReviewSlice) core.ReviewSlice {
	result := make(core.ReviewSlice, len(rows))
	for i, row := range rows {
		result[i] = store.fromRow(row)
	}
	return result
}

func (store *ReviewStore) Add(ctx context.Context, review *core.Review) error {
	row := store.toRow(review)
	if err := row.Insert(ctx, shared.GetExecutorOrDefault(ctx, store.ContextExecutor), boil.Infer()); err != nil {
		return errors.Wrap(err, "insert query")
	}
	*review = *store.fromRow(row)
	return nil
}

func (store *ReviewStore) Delete(ctx context.Context, id core.ReviewID) error {
	rows, err := (&dal.Review{ID: int(id)}).Delete(ctx, shared.GetExecutorOrDefault(ctx, store.ContextExecutor))
	if err != nil {
		return errors.Wrap(err, "delete query")
	}

	if rows == 0 {
		return core.ErrReviewNotFound
	}
	return nil
}

func (store *ReviewStore) Update(ctx context.Context, review *core.Review) error {
	row := store.toRow(review)
	n, err := row.Update(ctx, shared.GetExecutorOrDefault(ctx, store.ContextExecutor), boil.Infer())
	if err != nil {
		return errors.Wrap(err, "update query")
	}

	if n == 0 {
		return core.ErrReviewNotFound
	}

	return nil
}

type ReviewStoreQuery struct {
	mods  []qm.QueryMod
	store *ReviewStore
}

func (store *ReviewStore) Query() core.ReviewStoreQuery {
	return &ReviewStoreQuery{store: store}
}

func (query *ReviewStoreQuery) Offset(offset int) core.ReviewStoreQuery {
	query.mods = append(query.mods, qm.Offset(offset))
	return query
}

func (query *ReviewStoreQuery) Limit(limit int) core.ReviewStoreQuery {
	query.mods = append(query.mods, qm.Limit(limit))
	return query
}

func (query *ReviewStoreQuery) OrderByCreatedAt() core.ReviewStoreQuery {
	query.mods = append(query.mods, qm.OrderBy("created_at DESC"))
	return query
}

func (query *ReviewStoreQuery) ID(ids ...core.ReviewID) core.ReviewStoreQuery {
	idsInt := make([]int, len(ids))
	for i, v := range ids {
		idsInt[i] = int(v)
	}

	query.mods = append(query.mods, dal.ReviewWhere.ID.IN(idsInt))
	return query
}

func (query *ReviewStoreQuery) Count(ctx context.Context) (int, error) {
	count, err := dal.
		Reviews(query.mods...).
		Count(ctx, shared.GetExecutorOrDefault(ctx, query.store.ContextExecutor))
	return int(count), err
}

func (query *ReviewStoreQuery) All(ctx context.Context) (core.ReviewSlice, error) {
	rows, err := dal.Reviews(query.mods...).All(ctx,
		shared.GetExecutorOrDefault(ctx, query.store.ContextExecutor),
	)
	if err != nil {
		return nil, err
	}

	return query.store.fromRowSlice(rows), nil
}

func (query *ReviewStoreQuery) One(ctx context.Context) (*core.Review, error) {
	row, err := dal.Reviews(query.mods...).One(ctx,
		shared.GetExecutorOrDefault(ctx, query.store.ContextExecutor),
	)
	if err == sql.ErrNoRows {
		return nil, core.ErrReviewNotFound
	} else if err != nil {
		return nil, err
	}

	return query.store.fromRow(row), nil
}
