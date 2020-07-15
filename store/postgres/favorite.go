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

type FavoriteStore struct {
	boil.ContextExecutor
}

func (store *FavoriteStore) toRow(favorite *core.Favorite) *dal.Favorite {
	return &dal.Favorite{
		ID:        int(favorite.ID),
		LotID:     int(favorite.LotID),
		UserID:    int(favorite.UserID),
		CreatedAt: favorite.CreatedAt,
	}
}

func (store *FavoriteStore) fromRow(row *dal.Favorite) *core.Favorite {
	return &core.Favorite{
		ID:        core.FavoriteID(row.ID),
		LotID:     core.LotID(row.LotID),
		UserID:    core.UserID(row.UserID),
		CreatedAt: row.CreatedAt,
	}
}

func (store *FavoriteStore) fromRowSlice(rows dal.FavoriteSlice) core.FavoriteSlice {
	result := make(core.FavoriteSlice, len(rows))
	for i, row := range rows {
		result[i] = store.fromRow(row)
	}
	return result
}

func (store *FavoriteStore) Add(ctx context.Context, favorite *core.Favorite) error {
	row := store.toRow(favorite)
	if err := row.Insert(ctx, shared.GetExecutorOrDefault(ctx, store.ContextExecutor), boil.Infer()); err != nil {
		return errors.Wrap(err, "insert query")
	}
	*favorite = *store.fromRow(row)
	return nil
}

func (store *FavoriteStore) Query() core.FavoriteStoreQuery {
	return &FavoriteStoreQuery{store: store}
}

type FavoriteStoreQuery struct {
	mods  []qm.QueryMod
	store *FavoriteStore
}

func (fsq *FavoriteStoreQuery) ID(ids ...core.FavoriteID) core.FavoriteStoreQuery {
	idsInt := make([]int, len(ids))
	for i, id := range ids {
		idsInt[i] = int(id)
	}
	fsq.mods = append(fsq.mods, dal.FavoriteWhere.ID.IN(idsInt))
	return fsq
}

func (fsq *FavoriteStoreQuery) LotID(ids ...core.LotID) core.FavoriteStoreQuery {
	idsInt := make([]int, len(ids))
	for i, id := range ids {
		idsInt[i] = int(id)
	}

	fsq.mods = append(fsq.mods, dal.FavoriteWhere.LotID.IN(idsInt))
	return fsq
}

func (fsq *FavoriteStoreQuery) UserID(id core.UserID) core.FavoriteStoreQuery {
	fsq.mods = append(fsq.mods, dal.FavoriteWhere.UserID.EQ(int(id)))
	return fsq
}

func (fsq *FavoriteStoreQuery) One(ctx context.Context) (*core.Favorite, error) {
	executor := shared.GetExecutorOrDefault(ctx, fsq.store.ContextExecutor)

	row, err := dal.Favorites(fsq.mods...).One(ctx, executor)
	if err == sql.ErrNoRows {
		return nil, core.ErrFavoriteNotFound
	} else if err != nil {
		return nil, err
	}
	return fsq.store.fromRow(row), nil
}

func (fsq *FavoriteStoreQuery) Delete(ctx context.Context) error {
	executor := shared.GetExecutorOrDefault(ctx, fsq.store.ContextExecutor)

	deleted, err := dal.Favorites(fsq.mods...).DeleteAll(ctx, executor)
	if err != nil {
		return err
	} else if deleted == 0 {
		return core.ErrFavoriteNotFound
	}
	return nil
}

func (fsq *FavoriteStoreQuery) All(ctx context.Context) (core.FavoriteSlice, error) {
	executor := shared.GetExecutorOrDefault(ctx, fsq.store.ContextExecutor)

	rows, err := dal.Favorites(fsq.mods...).All(ctx, executor)
	if err != nil {
		return nil, err
	}
	return fsq.store.fromRowSlice(rows), nil
}
