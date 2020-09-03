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

type CouponApplyStore struct {
	boil.ContextExecutor
}

func (store *CouponApplyStore) toRow(apply *core.CouponApply) *dal.CouponApply {
	return &dal.CouponApply{
		CouponID:  int(apply.CouponID),
		PaymentID: int(apply.PaymentID),
	}
}

func (store *CouponApplyStore) fromRow(row *dal.CouponApply) *core.CouponApply {
	return &core.CouponApply{
		CouponID:  core.CouponID(row.CouponID),
		PaymentID: core.PaymentID(row.PaymentID),
	}
}

func (store *CouponApplyStore) Add(ctx context.Context, apply *core.CouponApply) error {
	row := store.toRow(apply)

	if err := row.Insert(ctx, shared.GetExecutorOrDefault(ctx, store.ContextExecutor), boil.Infer()); err != nil {
		return errors.Wrap(err, "insert query")
	}

	*apply = *store.fromRow(row)
	return nil
}

func (store *CouponApplyStore) Query() core.CouponApplyStoreQuery {
	return &CouponApplyStoreQuery{store: store}
}

type CouponApplyStoreQuery struct {
	mods  []qm.QueryMod
	store *CouponApplyStore
}

func (casq *CouponApplyStoreQuery) CouponID(ids ...core.CouponID) core.CouponApplyStoreQuery {
	idsInt := make([]int, len(ids))
	for i, id := range ids {
		idsInt[i] = int(id)
	}

	casq.mods = append(casq.mods, dal.CouponApplyWhere.CouponID.IN(idsInt))
	return casq
}

func (casq *CouponApplyStoreQuery) Payments(purpose core.PaymentPurpose) core.CouponApplyStoreQuery {
	casq.mods = append(casq.mods,
		qm.InnerJoin("payment on coupon_apply.payment_id = payment.id"),
		qm.Where("payment.purpose = ?", purpose.String()),
	)
	return casq
}

func (casq *CouponApplyStoreQuery) UserID(id core.UserID) core.CouponApplyStoreQuery {
	casq.mods = append(casq.mods,
		qm.Where("payment.payer_id = ?", int(id)),
	)
	return casq
}

func (casq *CouponApplyStoreQuery) Success() core.CouponApplyStoreQuery {
	casq.mods = append(casq.mods,
		qm.Where("payment.status = ?", dal.PaymentStatusSuccess),
	)
	return casq
}

func (casq *CouponApplyStoreQuery) One(ctx context.Context) (*core.CouponApply, error) {
	executor := shared.GetExecutorOrDefault(ctx, casq.store.ContextExecutor)

	row, err := dal.CouponApplies(casq.mods...).One(ctx, executor)
	if err == sql.ErrNoRows {
		return nil, core.ErrCouponApplyNotFound
	} else if err != nil {
		return nil, err
	}

	return casq.store.fromRow(row), nil
}

func (casq *CouponApplyStoreQuery) Count(ctx context.Context) (int, error) {
	executor := shared.GetExecutorOrDefault(ctx, casq.store.ContextExecutor)

	count, err := dal.CouponApplies(casq.mods...).Count(ctx, executor)

	return int(count), err
}
