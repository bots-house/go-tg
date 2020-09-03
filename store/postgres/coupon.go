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

type CouponStore struct {
	boil.ContextExecutor
}

func (store *CouponStore) toRow(coupon *core.Coupon) *dal.Coupon {
	purposes := make([]string, len(coupon.PaymentPurposes))
	for i, purpose := range coupon.PaymentPurposes {
		purposes[i] = purpose.String()
	}

	return &dal.Coupon{
		ID:                    int(coupon.ID),
		Code:                  coupon.Code,
		Discount:              coupon.Discount,
		PaymentPurposes:       purposes,
		ExpireAt:              coupon.ExpireAt,
		MaxAppliesByUserLimit: coupon.MaxAppliesByUserLimit,
		MaxAppliesLimit:       coupon.MaxAppliesLimit,
		IsDeleted:             coupon.IsDeleted,
		CreatedAt:             coupon.CreatedAt,
	}
}

func (store *CouponStore) fromRow(row *dal.Coupon) (*core.Coupon, error) {
	purposes := make([]core.PaymentPurpose, len(row.PaymentPurposes))
	for i, purpose := range row.PaymentPurposes {
		var err error
		purposes[i], err = core.ParsePaymentPurpose(purpose)
		if err != nil {
			return nil, errors.Wrap(err, "parse payment purpose")
		}
	}

	return &core.Coupon{
		ID:                    core.CouponID(row.ID),
		Code:                  row.Code,
		Discount:              row.Discount,
		PaymentPurposes:       purposes,
		ExpireAt:              row.ExpireAt,
		MaxAppliesByUserLimit: row.MaxAppliesByUserLimit,
		MaxAppliesLimit:       row.MaxAppliesLimit,
		IsDeleted:             row.IsDeleted,
		CreatedAt:             row.CreatedAt,
	}, nil
}

func (store *CouponStore) fromRowSlice(rows dal.CouponSlice) (core.CouponSlice, error) {
	out := make(core.CouponSlice, len(rows))
	for i, row := range rows {
		var err error
		out[i], err = store.fromRow(row)
		if err != nil {
			return nil, err
		}
	}
	return out, nil
}

func (store *CouponStore) Add(ctx context.Context, coupon *core.Coupon) error {
	row := store.toRow(coupon)

	if err := row.Insert(ctx, shared.GetExecutorOrDefault(ctx, store.ContextExecutor), boil.Infer()); err != nil {
		return errors.Wrap(err, "insert query")
	}

	c, err := store.fromRow(row)
	if err != nil {
		return err
	}

	*coupon = *c
	return nil
}

func (store *CouponStore) Update(ctx context.Context, coupon *core.Coupon) error {
	row := store.toRow(coupon)

	n, err := row.Update(ctx, shared.GetExecutorOrDefault(ctx, store.ContextExecutor), boil.Infer())
	if err != nil {
		return errors.Wrap(err, "update query")
	}
	if n == 0 {
		return core.ErrCouponNotFound
	}

	return nil
}

func (store *CouponStore) Query() core.CouponStoreQuery {
	return &CouponStoreQuery{store: store}
}

type CouponStoreQuery struct {
	mods  []qm.QueryMod
	store *CouponStore
}

func (csq *CouponStoreQuery) ID(ids ...core.CouponID) core.CouponStoreQuery {
	idsInt := make([]int, len(ids))
	for i, id := range ids {
		idsInt[i] = int(id)
	}

	csq.mods = append(csq.mods, dal.CouponWhere.ID.IN(idsInt))
	return csq
}

func (csq *CouponStoreQuery) Code(code string) core.CouponStoreQuery {
	csq.mods = append(csq.mods, dal.CouponWhere.Code.EQ(code))
	return csq
}

func (csq *CouponStoreQuery) Offset(v int) core.CouponStoreQuery {
	csq.mods = append(csq.mods, qm.Offset(v))
	return csq
}

func (csq *CouponStoreQuery) Limit(v int) core.CouponStoreQuery {
	csq.mods = append(csq.mods, qm.Limit(v))
	return csq
}

func (csq *CouponStoreQuery) IsDeleted(isDeleted bool) core.CouponStoreQuery {
	csq.mods = append(csq.mods, dal.CouponWhere.IsDeleted.EQ(isDeleted))
	return csq
}

func (csq *CouponStoreQuery) One(ctx context.Context) (*core.Coupon, error) {
	executor := shared.GetExecutorOrDefault(ctx, csq.store.ContextExecutor)

	row, err := dal.Coupons(csq.mods...).One(ctx, executor)
	if err == sql.ErrNoRows {
		return nil, core.ErrCouponNotFound
	} else if err != nil {
		return nil, err
	}

	return csq.store.fromRow(row)
}

func (csq *CouponStoreQuery) All(ctx context.Context) (core.CouponSlice, error) {
	executor := shared.GetExecutorOrDefault(ctx, csq.store.ContextExecutor)
	rows, err := dal.Coupons(csq.mods...).All(ctx, executor)
	if err != nil {
		return nil, err
	}

	return csq.store.fromRowSlice(rows)
}

func (csq *CouponStoreQuery) Count(ctx context.Context) (int, error) {
	executor := shared.GetExecutorOrDefault(ctx, csq.store.ContextExecutor)

	count, err := dal.Coupons(csq.mods...).Count(ctx, executor)

	return int(count), err
}
