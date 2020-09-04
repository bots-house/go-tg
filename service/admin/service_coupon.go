package admin

import (
	"context"
	"time"

	"github.com/bots-house/birzzha/core"
	"github.com/pkg/errors"
	"github.com/volatiletech/null/v8"
)

type CouponInput struct {
	Code                  string
	Discount              float64
	Purposes              []string
	ExpireAt              null.Time
	MaxAppliesByUserLimit null.Int
	MaxAppliesLimit       null.Int
}

type CouponItem struct {
	ID                    core.CouponID
	Code                  string
	Discount              float64
	Purposes              []core.PaymentPurpose
	ExpireAt              null.Time
	CreatedAt             time.Time
	MaxAppliesByUserLimit null.Int
	MaxAppliesLimit       null.Int
	AppliesCount          int
}

type CouponListItem struct {
	Total int
	Items []*CouponItem
}

var (
	ErrCouponWithThisCodeAlreadyExist      = core.NewError("coupon_with_this_code_already_exist", "coupon with this code already exist")
	ErrCouponDiscountMustBeGreaterThanZero = core.NewError("coupon_discount_must_be_greater_than_zero", "coupon discount must be greater than zero")
	ErrCouponDiscountMustBeSmaller         = core.NewError("coupon_discount_must_be_smaller", "coupon discount must be smaller")
)

func (srv *Service) IsExistCoupon(ctx context.Context, code string) error {
	_, err := srv.Coupon.Query().IsDeleted(false).Code(code).One(ctx)
	if err != core.ErrCouponNotFound && err != nil {
		return errors.Wrap(err, "get coupon")
	} else if err == nil {
		return ErrCouponWithThisCodeAlreadyExist
	}
	return nil
}

func (srv *Service) CreateCoupon(ctx context.Context, user *core.User, in *CouponInput) (*CouponItem, error) {
	if err := srv.IsAdmin(user); err != nil {
		return nil, err
	}

	if in.Discount <= 0 {
		return nil, ErrCouponDiscountMustBeGreaterThanZero
	}

	if in.Discount > 1 {
		return nil, ErrCouponDiscountMustBeSmaller
	}

	if err := srv.IsExistCoupon(ctx, in.Code); err != nil {
		return nil, err
	}

	purposes := make([]core.PaymentPurpose, len(in.Purposes))
	for i, purpose := range in.Purposes {
		var err error
		purposes[i], err = core.ParsePaymentPurpose(purpose)
		if err != nil {
			return nil, errors.Wrap(err, "parse payment purpose")
		}
	}

	coupon := core.NewCoupon(
		in.Discount,
		purposes,
		in.ExpireAt,
		in.MaxAppliesByUserLimit,
		in.MaxAppliesLimit,
		in.Code,
	)

	if err := srv.Coupon.Add(ctx, coupon); err != nil {
		return nil, errors.Wrap(err, "create coupon")
	}

	applies, err := srv.CouponApply.Query().CouponID(coupon.ID).Count(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get applies count")
	}

	return &CouponItem{
		ID:                    coupon.ID,
		Code:                  coupon.Code,
		Discount:              coupon.Discount,
		Purposes:              coupon.PaymentPurposes,
		ExpireAt:              coupon.ExpireAt,
		CreatedAt:             coupon.CreatedAt,
		MaxAppliesByUserLimit: coupon.MaxAppliesByUserLimit,
		MaxAppliesLimit:       coupon.MaxAppliesLimit,
		AppliesCount:          applies,
	}, nil
}

func (srv *Service) UpdateCoupon(ctx context.Context, user *core.User, id core.CouponID, in *CouponInput) (*CouponItem, error) {
	if err := srv.IsAdmin(user); err != nil {
		return nil, err
	}

	if in.Discount <= 0 {
		return nil, ErrCouponDiscountMustBeGreaterThanZero
	}

	if in.Discount > 1 {
		return nil, ErrCouponDiscountMustBeSmaller
	}

	coupon, err := srv.Coupon.Query().IsDeleted(false).ID(id).One(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get coupon")
	}

	c, err := srv.Coupon.Query().IsDeleted(false).Code(in.Code).One(ctx)
	if err != core.ErrCouponNotFound && err != nil {
		return nil, errors.Wrap(err, "get coupon")
	} else if err == nil {
		if c.ID != coupon.ID {
			return nil, ErrCouponWithThisCodeAlreadyExist
		}
	}

	purposes := make([]core.PaymentPurpose, len(in.Purposes))
	for i, purpose := range in.Purposes {
		var err error
		purposes[i], err = core.ParsePaymentPurpose(purpose)
		if err != nil {
			return nil, errors.Wrap(err, "parse payment purpose")
		}
	}

	coupon.Code = in.Code
	coupon.Discount = in.Discount
	coupon.PaymentPurposes = purposes
	coupon.ExpireAt = in.ExpireAt
	coupon.MaxAppliesByUserLimit = in.MaxAppliesByUserLimit
	coupon.MaxAppliesLimit = in.MaxAppliesLimit

	if err := srv.Coupon.Update(ctx, coupon); err != nil {
		return nil, errors.Wrap(err, "update coupon")
	}

	return srv.newCouponItem(ctx, coupon)
}

func (srv *Service) DeleteCoupon(ctx context.Context, user *core.User, id core.CouponID) error {
	if err := srv.IsAdmin(user); err != nil {
		return err
	}

	coupon, err := srv.Coupon.Query().IsDeleted(false).ID(id).One(ctx)
	if err != nil {
		return errors.Wrap(err, "get coupon")
	}

	coupon.IsDeleted = true
	if err := srv.Coupon.Update(ctx, coupon); err != nil {
		return errors.Wrap(err, "update coupon")
	}
	return nil
}

func (srv *Service) newCouponItem(ctx context.Context, coupon *core.Coupon) (*CouponItem, error) {
	applies, err := srv.CouponApply.Query().CouponID(coupon.ID).Count(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get applies count")
	}

	return &CouponItem{
		ID:                    coupon.ID,
		Code:                  coupon.Code,
		Discount:              coupon.Discount,
		Purposes:              coupon.PaymentPurposes,
		ExpireAt:              coupon.ExpireAt,
		CreatedAt:             coupon.CreatedAt,
		MaxAppliesByUserLimit: coupon.MaxAppliesByUserLimit,
		MaxAppliesLimit:       coupon.MaxAppliesLimit,
		AppliesCount:          applies,
	}, nil
}

func (srv *Service) newCouponItemSlice(ctx context.Context, coupons core.CouponSlice) ([]*CouponItem, error) {
	out := make([]*CouponItem, len(coupons))
	for i, coupon := range coupons {
		var err error
		out[i], err = srv.newCouponItem(ctx, coupon)
		if err != nil {
			return nil, err
		}
	}
	return out, nil
}

func (srv *Service) GetCoupons(ctx context.Context, user *core.User, limit int, offset int) (*CouponListItem, error) {
	if err := srv.IsAdmin(user); err != nil {
		return nil, err
	}

	coupons, err := srv.Coupon.Query().IsDeleted(false).Limit(limit).Offset(offset).All(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get coupons")
	}

	couponsCount, err := srv.Coupon.Query().IsDeleted(false).Count(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get coupons count")
	}

	couponItemSlice, err := srv.newCouponItemSlice(ctx, coupons)
	if err != nil {
		return nil, err
	}

	return &CouponListItem{
		Total: couponsCount,
		Items: couponItemSlice,
	}, nil
}
