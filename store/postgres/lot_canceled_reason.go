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

type LotCanceledReasonStore struct {
	boil.ContextExecutor
}

func (store *LotCanceledReasonStore) toRow(lcr *core.LotCanceledReason) *dal.LotCanceledReason {
	return &dal.LotCanceledReason{
		ID:        int(lcr.ID),
		Why:       lcr.Why,
		IsPublic:  lcr.IsPublic,
		CreatedAt: lcr.CreatedAt,
		UpdatedAt: lcr.UpdatedAt,
	}
}

func (store *LotCanceledReasonStore) fromRow(row *dal.LotCanceledReason) *core.LotCanceledReason {
	return &core.LotCanceledReason{
		ID:        core.LotCanceledReasonID(row.ID),
		Why:       row.Why,
		IsPublic:  row.IsPublic,
		CreatedAt: row.CreatedAt,
		UpdatedAt: row.UpdatedAt,
	}
}

func (store *LotCanceledReasonStore) Add(ctx context.Context, lcr *core.LotCanceledReason) error {
	row := store.toRow(lcr)
	if err := row.Insert(ctx, shared.GetExecutorOrDefault(ctx, store.ContextExecutor), boil.Infer()); err != nil {
		return errors.Wrap(err, "insert query")
	}
	*lcr = *store.fromRow(row)
	return nil
}

type LotCanceledReasonStoreQuery struct {
	mods  []qm.QueryMod
	store *LotCanceledReasonStore
}

func (store *LotCanceledReasonStore) Query() core.LotCanceledReasonStoreQuery {
	return &LotCanceledReasonStoreQuery{store: store}
}

func (query *LotCanceledReasonStoreQuery) ID(id core.LotCanceledReasonID) core.LotCanceledReasonStoreQuery {
	query.mods = append(query.mods, dal.LotCanceledReasonWhere.ID.EQ(int(id)))
	return query
}

func (query *LotCanceledReasonStoreQuery) One(ctx context.Context) (*core.LotCanceledReason, error) {
	row, err := dal.LotCanceledReasons(query.mods...).One(ctx,
		shared.GetExecutorOrDefault(ctx, query.store.ContextExecutor),
	)
	if err == sql.ErrNoRows {
		return nil, core.ErrLotCanceledReasonNotFound
	} else if err != nil {
		return nil, err
	}

	return query.store.fromRow(row), nil
}

func (query *LotCanceledReasonStoreQuery) All(ctx context.Context) ([]*core.LotCanceledReason, error) {
	rows, err := dal.LotCanceledReasons(query.mods...).All(ctx,
		shared.GetExecutorOrDefault(ctx, query.store.ContextExecutor),
	)
	if err != nil {
		return nil, err
	}

	res := make([]*core.LotCanceledReason, len(rows))

	for i, row := range rows {
		res[i] = query.store.fromRow(row)
	}

	return res, nil
}
