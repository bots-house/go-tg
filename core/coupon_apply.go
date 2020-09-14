package core

import "context"

type CouponApply struct {
	CouponID  CouponID
	PaymentID PaymentID
}

func NewCouponApply(
	couponID CouponID,
	paymentID PaymentID,
) *CouponApply {
	return &CouponApply{
		CouponID:  couponID,
		PaymentID: paymentID,
	}
}

var (
	ErrCouponApplyNotFound = NewError("coupon_apply_not_found", "coupon apply not found")
)

type CouponApplyStore interface {
	Add(ctx context.Context, apply *CouponApply) error
	Query() CouponApplyStoreQuery
}

type CouponApplyStoreQuery interface {
	Payments(purpose PaymentPurpose) CouponApplyStoreQuery
	CouponID(ids ...CouponID) CouponApplyStoreQuery
	UserID(id UserID) CouponApplyStoreQuery
	PaymentID(id PaymentID) CouponApplyStoreQuery

	Success() CouponApplyStoreQuery
	One(ctx context.Context) (*CouponApply, error)
	Count(ctx context.Context) (int, error)
}
