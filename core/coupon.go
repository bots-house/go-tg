package core

import (
	"context"
	"time"

	"github.com/volatiletech/null/v8"
)

type CouponID int

type Coupon struct {
	// Unique ID of coupon
	ID CouponID

	// String sequence
	Code string

	// Discout (example: 0.1 = 10%)
	Discount float64

	// Payment purposes
	PaymentPurposes []PaymentPurpose

	// When coupon can be expire
	ExpireAt null.Time

	// Max applies for unique user
	MaxAppliesByUserLimit null.Int

	// Max applies for all users
	MaxAppliesLimit null.Int

	// Is deleted
	IsDeleted bool

	// When coupon was created
	CreatedAt time.Time
}

func (c Coupon) IsExpired() bool {
	if c.ExpireAt.Valid {
		return c.ExpireAt.Time.Unix() < time.Now().Unix()
	}
	return false
}

func (c Coupon) IsFullDiscounted() bool {
	return c.Discount == 1
}

func (c Coupon) IsConsistPurpose(purpose PaymentPurpose) bool {
	for _, p := range c.PaymentPurposes {
		if purpose == p {
			return true
		}
	}
	return false
}

func (c Coupon) IsCouponCanApply(applies int) bool {
	if c.MaxAppliesLimit.Valid {
		return c.MaxAppliesLimit.Int > applies
	}
	return true
}

func (c Coupon) IsCouponCanApplyByUser(applies int) bool {
	if c.MaxAppliesByUserLimit.Valid {
		return c.MaxAppliesByUserLimit.Int > applies
	}
	return true
}

type CouponSlice []*Coupon

func NewCoupon(
	discount float64,
	paymentPurposes []PaymentPurpose,
	expireAt null.Time,
	maxAppliesByUserLimit null.Int,
	maxAppliesLimit null.Int,
	code string,
) *Coupon {
	return &Coupon{
		Discount:              discount,
		PaymentPurposes:       paymentPurposes,
		ExpireAt:              expireAt,
		MaxAppliesByUserLimit: maxAppliesByUserLimit,
		MaxAppliesLimit:       maxAppliesLimit,
		CreatedAt:             time.Now(),
		Code:                  code,
		IsDeleted:             false,
	}
}

var (
	ErrCouponNotFound = NewError("coupon_not_found", "coupon not found")
)

type CouponStore interface {
	Add(ctx context.Context, coupon *Coupon) error
	Update(ctx context.Context, coupon *Coupon) error
	Query() CouponStoreQuery
}

type CouponStoreQuery interface {
	ID(ids ...CouponID) CouponStoreQuery
	IsDeleted(isDeleted bool) CouponStoreQuery
	Code(code string) CouponStoreQuery
	Limit(limit int) CouponStoreQuery
	Offset(offset int) CouponStoreQuery
	One(ctx context.Context) (*Coupon, error)
	All(ctx context.Context) (CouponSlice, error)
	Count(ctx context.Context) (int, error)
}
