package models

import (
	"github.com/bots-house/birzzha/api/gen/models"
	"github.com/bots-house/birzzha/service/admin"
	"github.com/bots-house/birzzha/service/personal"
	"github.com/go-openapi/swag"
)

func NewCouponItem(in *admin.CouponItem) *models.CouponItem {
	purposes := make([]string, len(in.Purposes))
	for i, purpose := range in.Purposes {
		purposes[i] = purpose.String()
	}

	return &models.CouponItem{
		ID:                    swag.Int64(int64(in.ID)),
		Code:                  swag.String(in.Code),
		Discount:              swag.Float64(in.Discount),
		Purposes:              purposes,
		ExpireAt:              nullTimeToUnix(in.ExpireAt),
		CreatedAt:             timeToUnix(in.CreatedAt),
		AppliesCount:          swag.Int64(int64(in.AppliesCount)),
		MaxAppliesByUserLimit: nullIntToInt64(in.MaxAppliesByUserLimit),
		MaxAppliesLimit:       nullIntToInt64(in.MaxAppliesLimit),
	}
}

func newCouponItemSlice(in []*admin.CouponItem) []*models.CouponItem {
	out := make([]*models.CouponItem, len(in))
	for i, v := range in {
		out[i] = NewCouponItem(v)
	}
	return out
}

func NewCouponListItem(in *admin.CouponListItem) *models.CouponListItem {
	return &models.CouponListItem{
		Total: swag.Int64(int64(in.Total)),
		Items: newCouponItemSlice(in.Items),
	}
}

func NewCouponInfo(in *personal.CouponInfo) *models.CouponInfo {
	return &models.CouponInfo{
		Discount: swag.Float64(in.Discount),
	}
}
