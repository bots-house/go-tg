package postgres

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/bots-house/birzzha/core"
	"github.com/bots-house/birzzha/store"
	"github.com/bots-house/birzzha/store/postgres/dal"
	"github.com/bots-house/birzzha/store/postgres/shared"
)

type LandingStore struct {
	db    *sql.DB
	txier store.Txier
}

func (store *LandingStore) get(ctx context.Context) (*dal.Landing, error) {
	executor := shared.GetExecutorOrDefault(ctx, store.db)

	return dal.Landings().One(ctx, executor)
}

func (store *LandingStore) fromRow(row *dal.Landing) *core.Landing {
	return &core.Landing{
		UniqueUsersPerMonthActual: row.UniqueUsersPerMonthActual,
		UniqueUsersPerMonthShift:  row.UniqueUsersPerMonthShift,

		AvgSiteReachActual: row.AvgSiteReachActual,
		AvgSiteReachShift:  row.AvgSiteReachShift,

		AvgChannelReachActual: row.AvgChannelReachActual,
		AvgChannelReachShift:  row.AvgChannelReachShift,
	}
}

func (store *LandingStore) toRow(landing *core.Landing) *dal.Landing {
	return &dal.Landing{
		ID: 1,

		UniqueUsersPerMonthActual: landing.UniqueUsersPerMonthActual,
		UniqueUsersPerMonthShift:  landing.UniqueUsersPerMonthShift,

		AvgSiteReachActual: landing.AvgSiteReachActual,
		AvgSiteReachShift:  landing.AvgSiteReachShift,

		AvgChannelReachActual: landing.AvgChannelReachActual,
		AvgChannelReachShift:  landing.AvgChannelReachShift,
	}
}

func (store *LandingStore) upsert(ctx context.Context, landing *core.Landing) error {
	executor := shared.GetExecutorOrDefault(ctx, store.db)

	if err := store.toRow(landing).Upsert(
		ctx,
		executor,
		true,
		[]string{"id"},
		boil.Infer(),
		boil.Infer(),
	); err != nil {
		return errors.Wrap(err, "upsert")
	}

	return nil
}

var defaultLanding = &core.Landing{}

func (store *LandingStore) Get(ctx context.Context) (*core.Landing, error) {
	var landing *core.Landing

	if err := store.txier(ctx, func(ctx context.Context) error {
		row, err := store.get(ctx)
		if err == sql.ErrNoRows {
			if err := store.upsert(ctx, defaultLanding); err != nil {
				return errors.Wrap(err, "insert default")
			}
			landing = defaultLanding
			return nil
		} else if err != nil {
			return errors.Wrap(err, "query row")
		}

		landing = store.fromRow(row)

		return nil
	}); err != nil {
		return nil, err
	}

	return landing, nil
}

func (store *LandingStore) Update(ctx context.Context, landing *core.Landing) error {
	if err := store.upsert(ctx, landing); err != nil {
		return errors.Wrap(err, "upsert")
	}

	return nil
}
