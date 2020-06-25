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

// LotTopic is an object representing the database table.
type LotTopic struct {
	ID      int `boil:"id" json:"id" toml:"id" yaml:"id"`
	LotID   int `boil:"lot_id" json:"lot_id" toml:"lot_id" yaml:"lot_id"`
	TopicID int `boil:"topic_id" json:"topic_id" toml:"topic_id" yaml:"topic_id"`

	R *lotTopicR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L lotTopicL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var LotTopicColumns = struct {
	ID      string
	LotID   string
	TopicID string
}{
	ID:      "id",
	LotID:   "lot_id",
	TopicID: "topic_id",
}

// Generated where

var LotTopicWhere = struct {
	ID      whereHelperint
	LotID   whereHelperint
	TopicID whereHelperint
}{
	ID:      whereHelperint{field: "\"lot_topic\".\"id\""},
	LotID:   whereHelperint{field: "\"lot_topic\".\"lot_id\""},
	TopicID: whereHelperint{field: "\"lot_topic\".\"topic_id\""},
}

// LotTopicRels is where relationship names are stored.
var LotTopicRels = struct {
	Lot   string
	Topic string
}{
	Lot:   "Lot",
	Topic: "Topic",
}

// lotTopicR is where relationships are stored.
type lotTopicR struct {
	Lot   *Lot   `boil:"Lot" json:"Lot" toml:"Lot" yaml:"Lot"`
	Topic *Topic `boil:"Topic" json:"Topic" toml:"Topic" yaml:"Topic"`
}

// NewStruct creates a new relationship struct
func (*lotTopicR) NewStruct() *lotTopicR {
	return &lotTopicR{}
}

// lotTopicL is where Load methods for each relationship are stored.
type lotTopicL struct{}

var (
	lotTopicAllColumns            = []string{"id", "lot_id", "topic_id"}
	lotTopicColumnsWithoutDefault = []string{"lot_id", "topic_id"}
	lotTopicColumnsWithDefault    = []string{"id"}
	lotTopicPrimaryKeyColumns     = []string{"id"}
)

type (
	// LotTopicSlice is an alias for a slice of pointers to LotTopic.
	// This should generally be used opposed to []LotTopic.
	LotTopicSlice []*LotTopic

	lotTopicQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	lotTopicType                 = reflect.TypeOf(&LotTopic{})
	lotTopicMapping              = queries.MakeStructMapping(lotTopicType)
	lotTopicPrimaryKeyMapping, _ = queries.BindMapping(lotTopicType, lotTopicMapping, lotTopicPrimaryKeyColumns)
	lotTopicInsertCacheMut       sync.RWMutex
	lotTopicInsertCache          = make(map[string]insertCache)
	lotTopicUpdateCacheMut       sync.RWMutex
	lotTopicUpdateCache          = make(map[string]updateCache)
	lotTopicUpsertCacheMut       sync.RWMutex
	lotTopicUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single lotTopic record from the query.
func (q lotTopicQuery) One(ctx context.Context, exec boil.ContextExecutor) (*LotTopic, error) {
	o := &LotTopic{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "dal: failed to execute a one query for lot_topic")
	}

	return o, nil
}

// All returns all LotTopic records from the query.
func (q lotTopicQuery) All(ctx context.Context, exec boil.ContextExecutor) (LotTopicSlice, error) {
	var o []*LotTopic

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "dal: failed to assign all query results to LotTopic slice")
	}

	return o, nil
}

// Count returns the count of all LotTopic records in the query.
func (q lotTopicQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "dal: failed to count lot_topic rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q lotTopicQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "dal: failed to check if lot_topic exists")
	}

	return count > 0, nil
}

// Lot pointed to by the foreign key.
func (o *LotTopic) Lot(mods ...qm.QueryMod) lotQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.LotID),
	}

	queryMods = append(queryMods, mods...)

	query := Lots(queryMods...)
	queries.SetFrom(query.Query, "\"lot\"")

	return query
}

// Topic pointed to by the foreign key.
func (o *LotTopic) Topic(mods ...qm.QueryMod) topicQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.TopicID),
	}

	queryMods = append(queryMods, mods...)

	query := Topics(queryMods...)
	queries.SetFrom(query.Query, "\"topic\"")

	return query
}

// LoadLot allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (lotTopicL) LoadLot(ctx context.Context, e boil.ContextExecutor, singular bool, maybeLotTopic interface{}, mods queries.Applicator) error {
	var slice []*LotTopic
	var object *LotTopic

	if singular {
		object = maybeLotTopic.(*LotTopic)
	} else {
		slice = *maybeLotTopic.(*[]*LotTopic)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &lotTopicR{}
		}
		args = append(args, object.LotID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &lotTopicR{}
			}

			for _, a := range args {
				if a == obj.LotID {
					continue Outer
				}
			}

			args = append(args, obj.LotID)

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
		foreign.R.LotTopics = append(foreign.R.LotTopics, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.LotID == foreign.ID {
				local.R.Lot = foreign
				if foreign.R == nil {
					foreign.R = &lotR{}
				}
				foreign.R.LotTopics = append(foreign.R.LotTopics, local)
				break
			}
		}
	}

	return nil
}

// LoadTopic allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (lotTopicL) LoadTopic(ctx context.Context, e boil.ContextExecutor, singular bool, maybeLotTopic interface{}, mods queries.Applicator) error {
	var slice []*LotTopic
	var object *LotTopic

	if singular {
		object = maybeLotTopic.(*LotTopic)
	} else {
		slice = *maybeLotTopic.(*[]*LotTopic)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &lotTopicR{}
		}
		args = append(args, object.TopicID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &lotTopicR{}
			}

			for _, a := range args {
				if a == obj.TopicID {
					continue Outer
				}
			}

			args = append(args, obj.TopicID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`topic`),
		qm.WhereIn(`topic.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Topic")
	}

	var resultSlice []*Topic
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Topic")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for topic")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for topic")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Topic = foreign
		if foreign.R == nil {
			foreign.R = &topicR{}
		}
		foreign.R.LotTopics = append(foreign.R.LotTopics, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.TopicID == foreign.ID {
				local.R.Topic = foreign
				if foreign.R == nil {
					foreign.R = &topicR{}
				}
				foreign.R.LotTopics = append(foreign.R.LotTopics, local)
				break
			}
		}
	}

	return nil
}

// SetLot of the lotTopic to the related item.
// Sets o.R.Lot to related.
// Adds o to related.R.LotTopics.
func (o *LotTopic) SetLot(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Lot) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"lot_topic\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"lot_id"}),
		strmangle.WhereClause("\"", "\"", 2, lotTopicPrimaryKeyColumns),
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

	o.LotID = related.ID
	if o.R == nil {
		o.R = &lotTopicR{
			Lot: related,
		}
	} else {
		o.R.Lot = related
	}

	if related.R == nil {
		related.R = &lotR{
			LotTopics: LotTopicSlice{o},
		}
	} else {
		related.R.LotTopics = append(related.R.LotTopics, o)
	}

	return nil
}

// SetTopic of the lotTopic to the related item.
// Sets o.R.Topic to related.
// Adds o to related.R.LotTopics.
func (o *LotTopic) SetTopic(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Topic) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"lot_topic\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"topic_id"}),
		strmangle.WhereClause("\"", "\"", 2, lotTopicPrimaryKeyColumns),
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

	o.TopicID = related.ID
	if o.R == nil {
		o.R = &lotTopicR{
			Topic: related,
		}
	} else {
		o.R.Topic = related
	}

	if related.R == nil {
		related.R = &topicR{
			LotTopics: LotTopicSlice{o},
		}
	} else {
		related.R.LotTopics = append(related.R.LotTopics, o)
	}

	return nil
}

// LotTopics retrieves all the records using an executor.
func LotTopics(mods ...qm.QueryMod) lotTopicQuery {
	mods = append(mods, qm.From("\"lot_topic\""))
	return lotTopicQuery{NewQuery(mods...)}
}

// FindLotTopic retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindLotTopic(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*LotTopic, error) {
	lotTopicObj := &LotTopic{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"lot_topic\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, lotTopicObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "dal: unable to select from lot_topic")
	}

	return lotTopicObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *LotTopic) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("dal: no lot_topic provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(lotTopicColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	lotTopicInsertCacheMut.RLock()
	cache, cached := lotTopicInsertCache[key]
	lotTopicInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			lotTopicAllColumns,
			lotTopicColumnsWithDefault,
			lotTopicColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(lotTopicType, lotTopicMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(lotTopicType, lotTopicMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"lot_topic\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"lot_topic\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "dal: unable to insert into lot_topic")
	}

	if !cached {
		lotTopicInsertCacheMut.Lock()
		lotTopicInsertCache[key] = cache
		lotTopicInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the LotTopic.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *LotTopic) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	lotTopicUpdateCacheMut.RLock()
	cache, cached := lotTopicUpdateCache[key]
	lotTopicUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			lotTopicAllColumns,
			lotTopicPrimaryKeyColumns,
		)

		if len(wl) == 0 {
			return 0, errors.New("dal: unable to update lot_topic, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"lot_topic\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, lotTopicPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(lotTopicType, lotTopicMapping, append(wl, lotTopicPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "dal: unable to update lot_topic row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dal: failed to get rows affected by update for lot_topic")
	}

	if !cached {
		lotTopicUpdateCacheMut.Lock()
		lotTopicUpdateCache[key] = cache
		lotTopicUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q lotTopicQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "dal: unable to update all for lot_topic")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dal: unable to retrieve rows affected for lot_topic")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o LotTopicSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), lotTopicPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"lot_topic\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, lotTopicPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "dal: unable to update all in lotTopic slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dal: unable to retrieve rows affected all in update all lotTopic")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *LotTopic) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("dal: no lot_topic provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(lotTopicColumnsWithDefault, o)

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

	lotTopicUpsertCacheMut.RLock()
	cache, cached := lotTopicUpsertCache[key]
	lotTopicUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			lotTopicAllColumns,
			lotTopicColumnsWithDefault,
			lotTopicColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			lotTopicAllColumns,
			lotTopicPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("dal: unable to upsert lot_topic, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(lotTopicPrimaryKeyColumns))
			copy(conflict, lotTopicPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"lot_topic\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(lotTopicType, lotTopicMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(lotTopicType, lotTopicMapping, ret)
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
		return errors.Wrap(err, "dal: unable to upsert lot_topic")
	}

	if !cached {
		lotTopicUpsertCacheMut.Lock()
		lotTopicUpsertCache[key] = cache
		lotTopicUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single LotTopic record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *LotTopic) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("dal: no LotTopic provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), lotTopicPrimaryKeyMapping)
	sql := "DELETE FROM \"lot_topic\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "dal: unable to delete from lot_topic")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dal: failed to get rows affected by delete for lot_topic")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q lotTopicQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("dal: no lotTopicQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "dal: unable to delete all from lot_topic")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dal: failed to get rows affected by deleteall for lot_topic")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o LotTopicSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), lotTopicPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"lot_topic\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, lotTopicPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "dal: unable to delete all from lotTopic slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dal: failed to get rows affected by deleteall for lot_topic")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *LotTopic) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindLotTopic(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *LotTopicSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := LotTopicSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), lotTopicPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"lot_topic\".* FROM \"lot_topic\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, lotTopicPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "dal: unable to reload all in LotTopicSlice")
	}

	*o = slice

	return nil
}

// LotTopicExists checks if the LotTopic row exists.
func LotTopicExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"lot_topic\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "dal: unable to check if lot_topic exists")
	}

	return exists, nil
}
