package personal

import (
	"context"

	"github.com/bots-house/birzzha/core"
	"github.com/pkg/errors"
)

type CouponSpec struct {
	Code    string
	Purpose string
}

var (
	ErrCouponIsNotValidForPayment = core.NewError("coupon_is_not_valid_for_payment", "coupon is not valid for payment")
	ErrCouponExpired              = core.NewError("coupon_expired", "coupon expired")
	ErrCouponMaxUserApplyingLimit = core.NewError("coupon_max_user_applying_limit", "coupon max user applying limit")
	ErrCouponMaxApplyingLimit     = core.NewError("coupon_max_applying_limit", "coupon user applying limit")
)

type CouponInfo struct {
	Discount float64
}

func (srv *Service) ValidateCoupon(ctx context.Context, user *core.User, coupon *core.Coupon, purpose core.PaymentPurpose) error {
	if !coupon.IsConsistPurpose(purpose) {
		return ErrCouponIsNotValidForPayment
	}

	if coupon.IsExpired() {
		return ErrCouponExpired
	}

	appliesCount, err := srv.CouponApply.Query().CouponID(coupon.ID).Payments(purpose).Success().Count(ctx)
	if err != nil {
		return errors.Wrap(err, "get applies count")
	}

	if !coupon.IsCouponCanApply(appliesCount) {
		return ErrCouponMaxApplyingLimit
	}

	appliesByUserCount, err := srv.CouponApply.Query().CouponID(coupon.ID).Payments(purpose).UserID(user.ID).Success().Count(ctx)
	if err != nil {
		return errors.Wrap(err, "get applies by user count")
	}

	if !coupon.IsCouponCanApplyByUser(appliesByUserCount) {
		return ErrCouponMaxUserApplyingLimit
	}
	return nil
}

func (srv *Service) GetCoupon(ctx context.Context, user *core.User, spec *CouponSpec) (*CouponInfo, error) {
	coupon, err := srv.Coupon.Query().IsDeleted(false).Code(spec.Code).One(ctx)
	if err != nil {
		return nil, err
	}

	purpose, err := core.ParsePaymentPurpose(spec.Purpose)
	if err != nil {
		return nil, errors.Wrap(err, "parse payment purpose")
	}

	if err := srv.ValidateCoupon(ctx, user, coupon, purpose); err != nil {
		return nil, err
	}

	return &CouponInfo{
		Discount: coupon.Discount,
	}, nil
}
