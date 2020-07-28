package postgres

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/pkg/errors"

	"github.com/bots-house/birzzha/pkg/log"
	"github.com/bots-house/birzzha/store"
	"github.com/bots-house/birzzha/store/postgres/migrations"
	"github.com/bots-house/birzzha/store/postgres/shared"
)

type Postgres struct {
	*sql.DB
	migrator *migrations.Migrator

	Lot               *LotStore
	LotCanceledReason *LotCanceledReasonStore
	User              *UserStore
	Topic             *TopicStore
	LotTopic          *LotTopicStore
	Review            *ReviewStore
	Settings          *SettingsStore
	Payment           *PaymentStore
	Favorite          *FavoriteStore
	LotFile           *LotFileStore
	Landing           *LandingStore
}

// NewPostgres create postgres based database with all stores.
func NewPostgres(db *sql.DB) *Postgres {
	pg := &Postgres{
		DB:                db,
		migrator:          migrations.New(db),
		User:              &UserStore{ContextExecutor: db},
		Topic:             &TopicStore{ContextExecutor: db},
		Review:            &ReviewStore{ContextExecutor: db},
		Payment:           &PaymentStore{ContextExecutor: db},
		Favorite:          &FavoriteStore{ContextExecutor: db},
		LotCanceledReason: &LotCanceledReasonStore{ContextExecutor: db},
		LotFile:           &LotFileStore{ContextExecutor: db},
	}

	pg.LotTopic = &LotTopicStore{db: db, txier: pg.Tx, executor: db}
	pg.Lot = &LotStore{ContextExecutor: db, txier: pg.Tx, lotTopicStore: pg.LotTopic}
	pg.Settings = &SettingsStore{db: db, txier: pg.Tx}
	pg.Landing = &LandingStore{db: db, txier: pg.Tx}

	return pg
}

func (p *Postgres) Migrator() store.Migrator {
	return p.migrator
}

// Tx run code in database transaction.
// Based on: https://stackoverflow.com/a/23502629.
func (p *Postgres) Tx(ctx context.Context, txFunc store.TxFunc) (err error) {
	tx := shared.GetTx(ctx)

	if tx != nil {
		return txFunc(ctx)
	}

	tx, err = p.BeginTx(ctx, nil)
	if err != nil {
		return errors.Wrap(err, "begin tx failed")
	}

	ctx = shared.WithTx(ctx, tx)

	//nolint:gocritic
	defer func() {
		if r := recover(); r != nil {
			if err := tx.Rollback(); err != nil {
				log.Warn(ctx, "tx rollback failed", "err", err)
			}
			panic(r)
		} else if err != nil {
			if err := tx.Rollback(); err != nil {
				log.Warn(ctx, "tx rollback failed", "err", err)
			}
		} else {
			err = tx.Commit()
		}
	}()

	err = txFunc(ctx)

	return err
}
