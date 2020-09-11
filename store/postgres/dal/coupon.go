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
	"github.com/volatiletech/sqlboiler/v4/types"
	"github.com/volatiletech/strmangle"
)

// Coupon is an object representing the database table.
type Coupon struct {
	ID                    int               `boil:"id" json:"id" toml:"id" yaml:"id"`
	Code                  string            `boil:"code" json:"code" toml:"code" yaml:"code"`
	Discount              float64           `boil:"discount" json:"discount" toml:"discount" yaml:"discount"`
	PaymentPurposes       types.StringArray `boil:"payment_purposes" json:"payment_purposes" toml:"payment_purposes" yaml:"payment_purposes"`
	ExpireAt              null.Time         `boil:"expire_at" json:"expire_at,omitempty" toml:"expire_at" yaml:"expire_at,omitempty"`
	MaxAppliesByUserLimit null.Int          `boil:"max_applies_by_user_limit" json:"max_applies_by_user_limit,omitempty" toml:"max_applies_by_user_limit" yaml:"max_applies_by_user_limit,omitempty"`
	MaxAppliesLimit       null.Int          `boil:"max_applies_limit" json:"max_applies_limit,omitempty" toml:"max_applies_limit" yaml:"max_applies_limit,omitempty"`
	IsDeleted             bool              `boil:"is_deleted" json:"is_deleted" toml:"is_deleted" yaml:"is_deleted"`
	CreatedAt             time.Time         `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`

	R *couponR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L couponL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var CouponColumns = struct {
	ID                    string
	Code                  string
	Discount              string
	PaymentPurposes       string
	ExpireAt              string
	MaxAppliesByUserLimit string
	MaxAppliesLimit       string
	IsDeleted             string
	CreatedAt             string
}{
	ID:                    "id",
	Code:                  "code",
	Discount:              "discount",
	PaymentPurposes:       "payment_purposes",
	ExpireAt:              "expire_at",
	MaxAppliesByUserLimit: "max_applies_by_user_limit",
	MaxAppliesLimit:       "max_applies_limit",
	IsDeleted:             "is_deleted",
	CreatedAt:             "created_at",
}

// Generated where

type whereHelperint struct{ field string }

func (w whereHelperint) EQ(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint) NEQ(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint) LT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint) LTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint) GT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint) GTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperint) IN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperint) NIN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperstring) IN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperstring) NIN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelperfloat64 struct{ field string }

func (w whereHelperfloat64) EQ(x float64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperfloat64) NEQ(x float64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelperfloat64) LT(x float64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperfloat64) LTE(x float64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelperfloat64) GT(x float64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperfloat64) GTE(x float64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}
func (w whereHelperfloat64) IN(slice []float64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperfloat64) NIN(slice []float64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelpertypes_StringArray struct{ field string }

func (w whereHelpertypes_StringArray) EQ(x types.StringArray) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelpertypes_StringArray) NEQ(x types.StringArray) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelpertypes_StringArray) LT(x types.StringArray) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpertypes_StringArray) LTE(x types.StringArray) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpertypes_StringArray) GT(x types.StringArray) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpertypes_StringArray) GTE(x types.StringArray) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

type whereHelpernull_Time struct{ field string }

func (w whereHelpernull_Time) EQ(x null.Time) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Time) NEQ(x null.Time) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Time) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Time) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }
func (w whereHelpernull_Time) LT(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Time) LTE(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Time) GT(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Time) GTE(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

type whereHelpernull_Int struct{ field string }

func (w whereHelpernull_Int) EQ(x null.Int) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Int) NEQ(x null.Int) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Int) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Int) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }
func (w whereHelpernull_Int) LT(x null.Int) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Int) LTE(x null.Int) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Int) GT(x null.Int) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Int) GTE(x null.Int) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

type whereHelperbool struct{ field string }

func (w whereHelperbool) EQ(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperbool) NEQ(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperbool) LT(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperbool) LTE(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperbool) GT(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperbool) GTE(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }

type whereHelpertime_Time struct{ field string }

func (w whereHelpertime_Time) EQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelpertime_Time) NEQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelpertime_Time) LT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpertime_Time) LTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpertime_Time) GT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpertime_Time) GTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var CouponWhere = struct {
	ID                    whereHelperint
	Code                  whereHelperstring
	Discount              whereHelperfloat64
	PaymentPurposes       whereHelpertypes_StringArray
	ExpireAt              whereHelpernull_Time
	MaxAppliesByUserLimit whereHelpernull_Int
	MaxAppliesLimit       whereHelpernull_Int
	IsDeleted             whereHelperbool
	CreatedAt             whereHelpertime_Time
}{
	ID:                    whereHelperint{field: "\"coupon\".\"id\""},
	Code:                  whereHelperstring{field: "\"coupon\".\"code\""},
	Discount:              whereHelperfloat64{field: "\"coupon\".\"discount\""},
	PaymentPurposes:       whereHelpertypes_StringArray{field: "\"coupon\".\"payment_purposes\""},
	ExpireAt:              whereHelpernull_Time{field: "\"coupon\".\"expire_at\""},
	MaxAppliesByUserLimit: whereHelpernull_Int{field: "\"coupon\".\"max_applies_by_user_limit\""},
	MaxAppliesLimit:       whereHelpernull_Int{field: "\"coupon\".\"max_applies_limit\""},
	IsDeleted:             whereHelperbool{field: "\"coupon\".\"is_deleted\""},
	CreatedAt:             whereHelpertime_Time{field: "\"coupon\".\"created_at\""},
}

// CouponRels is where relationship names are stored.
var CouponRels = struct {
	CouponApplies string
}{
	CouponApplies: "CouponApplies",
}

// couponR is where relationships are stored.
type couponR struct {
	CouponApplies CouponApplySlice `boil:"CouponApplies" json:"CouponApplies" toml:"CouponApplies" yaml:"CouponApplies"`
}

// NewStruct creates a new relationship struct
func (*couponR) NewStruct() *couponR {
	return &couponR{}
}

// couponL is where Load methods for each relationship are stored.
type couponL struct{}

var (
	couponAllColumns            = []string{"id", "code", "discount", "payment_purposes", "expire_at", "max_applies_by_user_limit", "max_applies_limit", "is_deleted", "created_at"}
	couponColumnsWithoutDefault = []string{"code", "discount", "payment_purposes", "expire_at", "max_applies_by_user_limit", "max_applies_limit", "is_deleted", "created_at"}
	couponColumnsWithDefault    = []string{"id"}
	couponPrimaryKeyColumns     = []string{"id"}
)

type (
	// CouponSlice is an alias for a slice of pointers to Coupon.
	// This should generally be used opposed to []Coupon.
	CouponSlice []*Coupon

	couponQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	couponType                 = reflect.TypeOf(&Coupon{})
	couponMapping              = queries.MakeStructMapping(couponType)
	couponPrimaryKeyMapping, _ = queries.BindMapping(couponType, couponMapping, couponPrimaryKeyColumns)
	couponInsertCacheMut       sync.RWMutex
	couponInsertCache          = make(map[string]insertCache)
	couponUpdateCacheMut       sync.RWMutex
	couponUpdateCache          = make(map[string]updateCache)
	couponUpsertCacheMut       sync.RWMutex
	couponUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single coupon record from the query.
func (q couponQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Coupon, error) {
	o := &Coupon{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "dal: failed to execute a one query for coupon")
	}

	return o, nil
}

// All returns all Coupon records from the query.
func (q couponQuery) All(ctx context.Context, exec boil.ContextExecutor) (CouponSlice, error) {
	var o []*Coupon

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "dal: failed to assign all query results to Coupon slice")
	}

	return o, nil
}

// Count returns the count of all Coupon records in the query.
func (q couponQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "dal: failed to count coupon rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q couponQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "dal: failed to check if coupon exists")
	}

	return count > 0, nil
}

// CouponApplies retrieves all the coupon_apply's CouponApplies with an executor.
func (o *Coupon) CouponApplies(mods ...qm.QueryMod) couponApplyQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"coupon_apply\".\"coupon_id\"=?", o.ID),
	)

	query := CouponApplies(queryMods...)
	queries.SetFrom(query.Query, "\"coupon_apply\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"coupon_apply\".*"})
	}

	return query
}

// LoadCouponApplies allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (couponL) LoadCouponApplies(ctx context.Context, e boil.ContextExecutor, singular bool, maybeCoupon interface{}, mods queries.Applicator) error {
	var slice []*Coupon
	var object *Coupon

	if singular {
		object = maybeCoupon.(*Coupon)
	} else {
		slice = *maybeCoupon.(*[]*Coupon)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &couponR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &couponR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`coupon_apply`),
		qm.WhereIn(`coupon_apply.coupon_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load coupon_apply")
	}

	var resultSlice []*CouponApply
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice coupon_apply")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on coupon_apply")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for coupon_apply")
	}

	if singular {
		object.R.CouponApplies = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &couponApplyR{}
			}
			foreign.R.Coupon = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.CouponID {
				local.R.CouponApplies = append(local.R.CouponApplies, foreign)
				if foreign.R == nil {
					foreign.R = &couponApplyR{}
				}
				foreign.R.Coupon = local
				break
			}
		}
	}

	return nil
}

// AddCouponApplies adds the given related objects to the existing relationships
// of the coupon, optionally inserting them as new records.
// Appends related to o.R.CouponApplies.
// Sets related.R.Coupon appropriately.
func (o *Coupon) AddCouponApplies(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*CouponApply) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.CouponID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"coupon_apply\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"coupon_id"}),
				strmangle.WhereClause("\"", "\"", 2, couponApplyPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.CouponID = o.ID
		}
	}

	if o.R == nil {
		o.R = &couponR{
			CouponApplies: related,
		}
	} else {
		o.R.CouponApplies = append(o.R.CouponApplies, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &couponApplyR{
				Coupon: o,
			}
		} else {
			rel.R.Coupon = o
		}
	}
	return nil
}

// Coupons retrieves all the records using an executor.
func Coupons(mods ...qm.QueryMod) couponQuery {
	mods = append(mods, qm.From("\"coupon\""))
	return couponQuery{NewQuery(mods...)}
}

// FindCoupon retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindCoupon(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*Coupon, error) {
	couponObj := &Coupon{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"coupon\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, couponObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "dal: unable to select from coupon")
	}

	return couponObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Coupon) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("dal: no coupon provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(couponColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	couponInsertCacheMut.RLock()
	cache, cached := couponInsertCache[key]
	couponInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			couponAllColumns,
			couponColumnsWithDefault,
			couponColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(couponType, couponMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(couponType, couponMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"coupon\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"coupon\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "dal: unable to insert into coupon")
	}

	if !cached {
		couponInsertCacheMut.Lock()
		couponInsertCache[key] = cache
		couponInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the Coupon.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Coupon) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	couponUpdateCacheMut.RLock()
	cache, cached := couponUpdateCache[key]
	couponUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			couponAllColumns,
			couponPrimaryKeyColumns,
		)

		if len(wl) == 0 {
			return 0, errors.New("dal: unable to update coupon, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"coupon\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, couponPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(couponType, couponMapping, append(wl, couponPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "dal: unable to update coupon row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dal: failed to get rows affected by update for coupon")
	}

	if !cached {
		couponUpdateCacheMut.Lock()
		couponUpdateCache[key] = cache
		couponUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q couponQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "dal: unable to update all for coupon")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dal: unable to retrieve rows affected for coupon")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o CouponSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), couponPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"coupon\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, couponPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "dal: unable to update all in coupon slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dal: unable to retrieve rows affected all in update all coupon")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Coupon) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("dal: no coupon provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(couponColumnsWithDefault, o)

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

	couponUpsertCacheMut.RLock()
	cache, cached := couponUpsertCache[key]
	couponUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			couponAllColumns,
			couponColumnsWithDefault,
			couponColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			couponAllColumns,
			couponPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("dal: unable to upsert coupon, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(couponPrimaryKeyColumns))
			copy(conflict, couponPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"coupon\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(couponType, couponMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(couponType, couponMapping, ret)
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
		return errors.Wrap(err, "dal: unable to upsert coupon")
	}

	if !cached {
		couponUpsertCacheMut.Lock()
		couponUpsertCache[key] = cache
		couponUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single Coupon record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Coupon) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("dal: no Coupon provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), couponPrimaryKeyMapping)
	sql := "DELETE FROM \"coupon\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "dal: unable to delete from coupon")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dal: failed to get rows affected by delete for coupon")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q couponQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("dal: no couponQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "dal: unable to delete all from coupon")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dal: failed to get rows affected by deleteall for coupon")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o CouponSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), couponPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"coupon\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, couponPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "dal: unable to delete all from coupon slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dal: failed to get rows affected by deleteall for coupon")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Coupon) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindCoupon(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CouponSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := CouponSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), couponPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"coupon\".* FROM \"coupon\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, couponPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "dal: unable to reload all in CouponSlice")
	}

	*o = slice

	return nil
}

// CouponExists checks if the Coupon row exists.
func CouponExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"coupon\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "dal: unable to check if coupon exists")
	}

	return exists, nil
}