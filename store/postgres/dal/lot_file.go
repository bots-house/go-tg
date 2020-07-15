// Code generated by SQLBoiler 4.1.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package dal

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// LotFile is an object representing the database table.
type LotFile struct {
	ID        int       `boil:"id" json:"id" toml:"id" yaml:"id"`
	LotID     null.Int  `boil:"lot_id" json:"lot_id,omitempty" toml:"lot_id" yaml:"lot_id,omitempty"`
	Name      string    `boil:"name" json:"name" toml:"name" yaml:"name"`
	Size      int       `boil:"size" json:"size" toml:"size" yaml:"size"`
	MimeType  string    `boil:"mime_type" json:"mime_type" toml:"mime_type" yaml:"mime_type"`
	Path      string    `boil:"path" json:"path" toml:"path" yaml:"path"`
	CreatedAt time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`

	R *lotFileR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L lotFileL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var LotFileColumns = struct {
	ID        string
	LotID     string
	Name      string
	Size      string
	MimeType  string
	Path      string
	CreatedAt string
}{
	ID:        "id",
	LotID:     "lot_id",
	Name:      "name",
	Size:      "size",
	MimeType:  "mime_type",
	Path:      "path",
	CreatedAt: "created_at",
}

// Generated where

var LotFileWhere = struct {
	ID        whereHelperint
	LotID     whereHelpernull_Int
	Name      whereHelperstring
	Size      whereHelperint
	MimeType  whereHelperstring
	Path      whereHelperstring
	CreatedAt whereHelpertime_Time
}{
	ID:        whereHelperint{field: "\"lot_file\".\"id\""},
	LotID:     whereHelpernull_Int{field: "\"lot_file\".\"lot_id\""},
	Name:      whereHelperstring{field: "\"lot_file\".\"name\""},
	Size:      whereHelperint{field: "\"lot_file\".\"size\""},
	MimeType:  whereHelperstring{field: "\"lot_file\".\"mime_type\""},
	Path:      whereHelperstring{field: "\"lot_file\".\"path\""},
	CreatedAt: whereHelpertime_Time{field: "\"lot_file\".\"created_at\""},
}

// LotFileRels is where relationship names are stored.
var LotFileRels = struct {
	Lot string
}{
	Lot: "Lot",
}

// lotFileR is where relationships are stored.
type lotFileR struct {
	Lot *Lot `boil:"Lot" json:"Lot" toml:"Lot" yaml:"Lot"`
}

// NewStruct creates a new relationship struct
func (*lotFileR) NewStruct() *lotFileR {
	return &lotFileR{}
}

// lotFileL is where Load methods for each relationship are stored.
type lotFileL struct{}

var (
	lotFileAllColumns            = []string{"id", "lot_id", "name", "size", "mime_type", "path", "created_at"}
	lotFileColumnsWithoutDefault = []string{"lot_id", "name", "size", "mime_type", "path", "created_at"}
	lotFileColumnsWithDefault    = []string{"id"}
	lotFilePrimaryKeyColumns     = []string{"id"}
)

type (
	// LotFileSlice is an alias for a slice of pointers to LotFile.
	// This should generally be used opposed to []LotFile.
	LotFileSlice []*LotFile

	lotFileQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	lotFileType                 = reflect.TypeOf(&LotFile{})
	lotFileMapping              = queries.MakeStructMapping(lotFileType)
	lotFilePrimaryKeyMapping, _ = queries.BindMapping(lotFileType, lotFileMapping, lotFilePrimaryKeyColumns)
	lotFileInsertCacheMut       sync.RWMutex
	lotFileInsertCache          = make(map[string]insertCache)
	lotFileUpdateCacheMut       sync.RWMutex
	lotFileUpdateCache          = make(map[string]updateCache)
	lotFileUpsertCacheMut       sync.RWMutex
	lotFileUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single lotFile record from the query.
func (q lotFileQuery) One(ctx context.Context, exec boil.ContextExecutor) (*LotFile, error) {
	o := &LotFile{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "dal: failed to execute a one query for lot_file")
	}

	return o, nil
}

// All returns all LotFile records from the query.
func (q lotFileQuery) All(ctx context.Context, exec boil.ContextExecutor) (LotFileSlice, error) {
	var o []*LotFile

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "dal: failed to assign all query results to LotFile slice")
	}

	return o, nil
}

// Count returns the count of all LotFile records in the query.
func (q lotFileQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "dal: failed to count lot_file rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q lotFileQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "dal: failed to check if lot_file exists")
	}

	return count > 0, nil
}

// Lot pointed to by the foreign key.
func (o *LotFile) Lot(mods ...qm.QueryMod) lotQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.LotID),
	}

	queryMods = append(queryMods, mods...)

	query := Lots(queryMods...)
	queries.SetFrom(query.Query, "\"lot\"")

	return query
}

// LoadLot allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (lotFileL) LoadLot(ctx context.Context, e boil.ContextExecutor, singular bool, maybeLotFile interface{}, mods queries.Applicator) error {
	var slice []*LotFile
	var object *LotFile

	if singular {
		object = maybeLotFile.(*LotFile)
	} else {
		slice = *maybeLotFile.(*[]*LotFile)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &lotFileR{}
		}
		if !queries.IsNil(object.LotID) {
			args = append(args, object.LotID)
		}

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &lotFileR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.LotID) {
					continue Outer
				}
			}

			if !queries.IsNil(obj.LotID) {
				args = append(args, obj.LotID)
			}

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`lot`),
		qm.WhereIn(`lot.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Lot")
	}

	var resultSlice []*Lot
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Lot")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for lot")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for lot")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Lot = foreign
		if foreign.R == nil {
			foreign.R = &lotR{}
		}
		foreign.R.LotFiles = append(foreign.R.LotFiles, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.LotID, foreign.ID) {
				local.R.Lot = foreign
				if foreign.R == nil {
					foreign.R = &lotR{}
				}
				foreign.R.LotFiles = append(foreign.R.LotFiles, local)
				break
			}
		}
	}

	return nil
}

// SetLot of the lotFile to the related item.
// Sets o.R.Lot to related.
// Adds o to related.R.LotFiles.
func (o *LotFile) SetLot(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Lot) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"lot_file\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"lot_id"}),
		strmangle.WhereClause("\"", "\"", 2, lotFilePrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	queries.Assign(&o.LotID, related.ID)
	if o.R == nil {
		o.R = &lotFileR{
			Lot: related,
		}
	} else {
		o.R.Lot = related
	}

	if related.R == nil {
		related.R = &lotR{
			LotFiles: LotFileSlice{o},
		}
	} else {
		related.R.LotFiles = append(related.R.LotFiles, o)
	}

	return nil
}

// RemoveLot relationship.
// Sets o.R.Lot to nil.
// Removes o from all passed in related items' relationships struct (Optional).
func (o *LotFile) RemoveLot(ctx context.Context, exec boil.ContextExecutor, related *Lot) error {
	var err error

	queries.SetScanner(&o.LotID, nil)
	if _, err = o.Update(ctx, exec, boil.Whitelist("lot_id")); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	if o.R != nil {
		o.R.Lot = nil
	}
	if related == nil || related.R == nil {
		return nil
	}

	for i, ri := range related.R.LotFiles {
		if queries.Equal(o.LotID, ri.LotID) {
			continue
		}

		ln := len(related.R.LotFiles)
		if ln > 1 && i < ln-1 {
			related.R.LotFiles[i] = related.R.LotFiles[ln-1]
		}
		related.R.LotFiles = related.R.LotFiles[:ln-1]
		break
	}
	return nil
}

// LotFiles retrieves all the records using an executor.
func LotFiles(mods ...qm.QueryMod) lotFileQuery {
	mods = append(mods, qm.From("\"lot_file\""))
	return lotFileQuery{NewQuery(mods...)}
}

// FindLotFile retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindLotFile(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*LotFile, error) {
	lotFileObj := &LotFile{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"lot_file\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, lotFileObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "dal: unable to select from lot_file")
	}

	return lotFileObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *LotFile) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("dal: no lot_file provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(lotFileColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	lotFileInsertCacheMut.RLock()
	cache, cached := lotFileInsertCache[key]
	lotFileInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			lotFileAllColumns,
			lotFileColumnsWithDefault,
			lotFileColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(lotFileType, lotFileMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(lotFileType, lotFileMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"lot_file\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"lot_file\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "dal: unable to insert into lot_file")
	}

	if !cached {
		lotFileInsertCacheMut.Lock()
		lotFileInsertCache[key] = cache
		lotFileInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the LotFile.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *LotFile) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	lotFileUpdateCacheMut.RLock()
	cache, cached := lotFileUpdateCache[key]
	lotFileUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			lotFileAllColumns,
			lotFilePrimaryKeyColumns,
		)

		if len(wl) == 0 {
			return 0, errors.New("dal: unable to update lot_file, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"lot_file\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, lotFilePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(lotFileType, lotFileMapping, append(wl, lotFilePrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "dal: unable to update lot_file row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dal: failed to get rows affected by update for lot_file")
	}

	if !cached {
		lotFileUpdateCacheMut.Lock()
		lotFileUpdateCache[key] = cache
		lotFileUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q lotFileQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "dal: unable to update all for lot_file")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dal: unable to retrieve rows affected for lot_file")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o LotFileSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("dal: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), lotFilePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"lot_file\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, lotFilePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "dal: unable to update all in lotFile slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dal: unable to retrieve rows affected all in update all lotFile")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *LotFile) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("dal: no lot_file provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(lotFileColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	lotFileUpsertCacheMut.RLock()
	cache, cached := lotFileUpsertCache[key]
	lotFileUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			lotFileAllColumns,
			lotFileColumnsWithDefault,
			lotFileColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			lotFileAllColumns,
			lotFilePrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("dal: unable to upsert lot_file, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(lotFilePrimaryKeyColumns))
			copy(conflict, lotFilePrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"lot_file\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(lotFileType, lotFileMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(lotFileType, lotFileMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "dal: unable to upsert lot_file")
	}

	if !cached {
		lotFileUpsertCacheMut.Lock()
		lotFileUpsertCache[key] = cache
		lotFileUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single LotFile record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *LotFile) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("dal: no LotFile provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), lotFilePrimaryKeyMapping)
	sql := "DELETE FROM \"lot_file\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "dal: unable to delete from lot_file")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dal: failed to get rows affected by delete for lot_file")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q lotFileQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("dal: no lotFileQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "dal: unable to delete all from lot_file")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dal: failed to get rows affected by deleteall for lot_file")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o LotFileSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), lotFilePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"lot_file\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, lotFilePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "dal: unable to delete all from lotFile slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dal: failed to get rows affected by deleteall for lot_file")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *LotFile) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindLotFile(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *LotFileSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := LotFileSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), lotFilePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"lot_file\".* FROM \"lot_file\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, lotFilePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "dal: unable to reload all in LotFileSlice")
	}

	*o = slice

	return nil
}

// LotFileExists checks if the LotFile row exists.
func LotFileExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"lot_file\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "dal: unable to check if lot_file exists")
	}

	return exists, nil
}
