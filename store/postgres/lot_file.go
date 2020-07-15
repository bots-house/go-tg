package postgres

import (
	"context"

	"github.com/bots-house/birzzha/core"
	"github.com/bots-house/birzzha/store/postgres/dal"
	"github.com/bots-house/birzzha/store/postgres/shared"
	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type LotFileStore struct {
	boil.ContextExecutor
}

func (store *LotFileStore) toRow(lf *core.LotFile) *dal.LotFile {
	file := &dal.LotFile{
		Name:      lf.Name,
		Size:      lf.Size,
		MimeType:  lf.MIMEType,
		Path:      lf.Path,
		CreatedAt: lf.CreatedAt,
	}

	if lf.LotID != 0 {
		file.LotID = shared.ToNullInt(int(lf.LotID))
	}

	if lf.ID != 0 {
		file.ID = int(lf.ID)
	}

	return file
}

func (store *LotFileStore) fromRow(row *dal.LotFile) *core.LotFile {
	file := &core.LotFile{
		ID:        core.LotFileID(row.ID),
		Name:      row.Name,
		Size:      row.Size,
		MIMEType:  row.MimeType,
		Path:      row.Path,
		CreatedAt: row.CreatedAt,
	}
	if row.LotID.Valid {
		file.LotID = core.LotID(row.LotID.Int)
	}
	return file
}

func (store *LotFileStore) fromRowSlice(rows dal.LotFileSlice) core.LotFileSlice {
	result := make(core.LotFileSlice, len(rows))
	for i, row := range rows {
		result[i] = store.fromRow(row)
	}
	return result
}

func (store *LotFileStore) Add(ctx context.Context, lf *core.LotFile) error {
	row := store.toRow(lf)
	if err := row.Insert(ctx, shared.GetExecutorOrDefault(ctx, store.ContextExecutor), boil.Infer()); err != nil {
		return errors.Wrap(err, "insert query")
	}
	*lf = *store.fromRow(row)
	return nil
}

func (store *LotFileStore) Update(ctx context.Context, lf *core.LotFile) error {
	row := store.toRow(lf)
	n, err := row.Update(ctx, shared.GetExecutorOrDefault(ctx, store.ContextExecutor), boil.Infer())
	if err != nil {
		return errors.Wrap(err, "update query")
	}

	if n == 0 {
		return core.ErrLotFileNotFound
	}

	return nil
}

func (store *LotFileStore) Query() core.LotFileStoreQuery {
	return &LotFileStoreQuery{store: store}
}

type LotFileStoreQuery struct {
	mods  []qm.QueryMod
	store *LotFileStore
}

func (lfsq *LotFileStoreQuery) ID(ids ...core.LotFileID) core.LotFileStoreQuery {
	idsInt := make([]int, len(ids))
	for i, id := range ids {
		idsInt[i] = int(id)
	}
	lfsq.mods = append(lfsq.mods, dal.LotFileWhere.ID.IN(idsInt))
	return lfsq
}

func (lfsq *LotFileStoreQuery) LotID(id core.LotID) core.LotFileStoreQuery {
	lfsq.mods = append(lfsq.mods, dal.LotFileWhere.LotID.EQ(shared.ToNullInt(int(id))))
	return lfsq
}

func (lfsq *LotFileStoreQuery) All(ctx context.Context) (core.LotFileSlice, error) {
	executor := shared.GetExecutorOrDefault(ctx, lfsq.store.ContextExecutor)

	rows, err := dal.LotFiles(lfsq.mods...).All(ctx, executor)
	if err != nil {
		return nil, err
	}
	return lfsq.store.fromRowSlice(rows), nil
}
