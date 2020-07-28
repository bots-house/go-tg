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
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Landing is an object representing the database table.
type Landing struct {
	ID                        int `boil:"id" json:"id" toml:"id" yaml:"id"`
	UniqueUsersPerMonthActual int `boil:"unique_users_per_month_actual" json:"unique_users_per_month_actual" toml:"unique_users_per_month_actual" yaml:"unique_users_per_month_actual"`
	UniqueUsersPerMonthShift  int `boil:"unique_users_per_month_shift" json:"unique_users_per_month_shift" toml:"unique_users_per_month_shift" yaml:"unique_users_per_month_shift"`
	AvgSiteReachActual        int `boil:"avg_site_reach_actual" json:"avg_site_reach_actual" toml:"avg_site_reach_actual" yaml:"avg_site_reach_actual"`
	AvgSiteReachShift         int `boil:"avg_site_reach_shift" json:"avg_site_reach_shift" toml:"avg_site_reach_shift" yaml:"avg_site_reach_shift"`
	AvgChannelReachActual     int `boil:"avg_channel_reach_actual" json:"avg_channel_reach_actual" toml:"avg_channel_reach_actual" yaml:"avg_channel_reach_actual"`
	AvgChannelReachShift      int `boil:"avg_channel_reach_shift" json:"avg_channel_reach_shift" toml:"avg_channel_reach_shift" yaml:"avg_channel_reach_shift"`

	R *landingR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L landingL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var LandingColumns = struct {
	ID                        string
	UniqueUsersPerMonthActual string
	UniqueUsersPerMonthShift  string
	AvgSiteReachActual        string
	AvgSiteReachShift         string
	AvgChannelReachActual     string
	AvgChannelReachShift      string
}{
	ID:                        "id",
	UniqueUsersPerMonthActual: "unique_users_per_month_actual",
	UniqueUsersPerMonthShift:  "unique_users_per_month_shift",
	AvgSiteReachActual:        "avg_site_reach_actual",
	AvgSiteReachShift:         "avg_site_reach_shift",
	AvgChannelReachActual:     "avg_channel_reach_actual",
	AvgChannelReachShift:      "avg_channel_reach_shift",
}

// Generated where

var LandingWhere = struct {
	ID                        whereHelperint
	UniqueUsersPerMonthActual whereHelperint
	UniqueUsersPerMonthShift  whereHelperint
	AvgSiteReachActual        whereHelperint
	AvgSiteReachShift         whereHelperint
	AvgChannelReachActual     whereHelperint
	AvgChannelReachShift      whereHelperint
}{
	ID:                        whereHelperint{field: "\"landing\".\"id\""},
	UniqueUsersPerMonthActual: whereHelperint{field: "\"landing\".\"unique_users_per_month_actual\""},
	UniqueUsersPerMonthShift:  whereHelperint{field: "\"landing\".\"unique_users_per_month_shift\""},
	AvgSiteReachActual:        whereHelperint{field: "\"landing\".\"avg_site_reach_actual\""},
	AvgSiteReachShift:         whereHelperint{field: "\"landing\".\"avg_site_reach_shift\""},
	AvgChannelReachActual:     whereHelperint{field: "\"landing\".\"avg_channel_reach_actual\""},
	AvgChannelReachShift:      whereHelperint{field: "\"landing\".\"avg_channel_reach_shift\""},
}

// LandingRels is where relationship names are stored.
var LandingRels = struct {
}{}

// landingR is where relationships are stored.
type landingR struct {
}

// NewStruct creates a new relationship struct
func (*landingR) NewStruct() *landingR {
	return &landingR{}
}

// landingL is where Load methods for each relationship are stored.
type landingL struct{}

var (
	landingAllColumns            = []string{"id", "unique_users_per_month_actual", "unique_users_per_month_shift", "avg_site_reach_actual", "avg_site_reach_shift", "avg_channel_reach_actual", "avg_channel_reach_shift"}
	landingColumnsWithoutDefault = []string{"id", "unique_users_per_month_actual", "unique_users_per_month_shift", "avg_site_reach_actual", "avg_site_reach_shift", "avg_channel_reach_actual", "avg_channel_reach_shift"}
	landingColumnsWithDefault    = []string{}
	landingPrimaryKeyColumns     = []string{"id"}
)

type (
	// LandingSlice is an alias for a slice of pointers to Landing.
	// This should generally be used opposed to []Landing.
	LandingSlice []*Landing

	landingQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	landingType                 = reflect.TypeOf(&Landing{})
	landingMapping              = queries.MakeStructMapping(landingType)
	landingPrimaryKeyMapping, _ = queries.BindMapping(landingType, landingMapping, landingPrimaryKeyColumns)
	landingInsertCacheMut       sync.RWMutex
	landingInsertCache          = make(map[string]insertCache)
	landingUpdateCacheMut       sync.RWMutex
	landingUpdateCache          = make(map[string]updateCache)
	landingUpsertCacheMut       sync.RWMutex
	landingUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single landing record from the query.
func (q landingQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Landing, error) {
	o := &Landing{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "dal: failed to execute a one query for landing")
	}

	return o, nil
}

// All returns all Landing records from the query.
func (q landingQuery) All(ctx context.Context, exec boil.ContextExecutor) (LandingSlice, error) {
	var o []*Landing

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "dal: failed to assign all query results to Landing slice")
	}

	return o, nil
}

// Count returns the count of all Landing records in the query.
func (q landingQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "dal: failed to count landing rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q landingQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "dal: failed to check if landing exists")
	}

	return count > 0, nil
}

// Landings retrieves all the records using an executor.
func Landings(mods ...qm.QueryMod) landingQuery {
	mods = append(mods, qm.From("\"landing\""))
	return landingQuery{NewQuery(mods...)}
}

// FindLanding retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindLanding(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*Landing, error) {
	landingObj := &Landing{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"landing\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, landingObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "dal: unable to select from landing")
	}

	return landingObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Landing) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("dal: no landing provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(landingColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	landingInsertCacheMut.RLock()
	cache, cached := landingInsertCache[key]
	landingInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			landingAllColumns,
			landingColumnsWithDefault,
			landingColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(landingType, landingMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(landingType, landingMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"landing\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"landing\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "dal: unable to insert into landing")
	}

	if !cached {
		landingInsertCacheMut.Lock()
		landingInsertCache[key] = cache
		landingInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the Landing.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Landing) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	landingUpdateCacheMut.RLock()
	cache, cached := landingUpdateCache[key]
	landingUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			landingAllColumns,
			landingPrimaryKeyColumns,
		)

		if len(wl) == 0 {
			return 0, errors.New("dal: unable to update landing, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"landing\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, landingPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(landingType, landingMapping, append(wl, landingPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "dal: unable to update landing row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dal: failed to get rows affected by update for landing")
	}

	if !cached {
		landingUpdateCacheMut.Lock()
		landingUpdateCache[key] = cache
		landingUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q landingQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "dal: unable to update all for landing")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dal: unable to retrieve rows affected for landing")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o LandingSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), landingPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"landing\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, landingPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "dal: unable to update all in landing slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dal: unable to retrieve rows affected all in update all landing")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Landing) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("dal: no landing provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(landingColumnsWithDefault, o)

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

	landingUpsertCacheMut.RLock()
	cache, cached := landingUpsertCache[key]
	landingUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			landingAllColumns,
			landingColumnsWithDefault,
			landingColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			landingAllColumns,
			landingPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("dal: unable to upsert landing, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(landingPrimaryKeyColumns))
			copy(conflict, landingPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"landing\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(landingType, landingMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(landingType, landingMapping, ret)
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
		return errors.Wrap(err, "dal: unable to upsert landing")
	}

	if !cached {
		landingUpsertCacheMut.Lock()
		landingUpsertCache[key] = cache
		landingUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single Landing record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Landing) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("dal: no Landing provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), landingPrimaryKeyMapping)
	sql := "DELETE FROM \"landing\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "dal: unable to delete from landing")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dal: failed to get rows affected by delete for landing")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q landingQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("dal: no landingQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "dal: unable to delete all from landing")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dal: failed to get rows affected by deleteall for landing")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o LandingSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), landingPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"landing\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, landingPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "dal: unable to delete all from landing slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dal: failed to get rows affected by deleteall for landing")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Landing) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindLanding(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *LandingSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := LandingSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), landingPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"landing\".* FROM \"landing\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, landingPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "dal: unable to reload all in LandingSlice")
	}

	*o = slice

	return nil
}

// LandingExists checks if the Landing row exists.
func LandingExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"landing\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "dal: unable to check if landing exists")
	}

	return exists, nil
}
