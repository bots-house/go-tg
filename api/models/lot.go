package models

import (
	"github.com/go-openapi/swag"

	"github.com/bots-house/birzzha/api/gen/models"
	"github.com/bots-house/birzzha/core"
	"github.com/bots-house/birzzha/pkg/storage"
	"github.com/bots-house/birzzha/service/catalog"
)

func newLotPrice(in core.LotPrice) *models.LotPrice {
	return &models.LotPrice{
		Current:   swag.Int64(int64(in.Current)),
		Previous:  nullIntToInt64(in.Previous),
		IsBargain: swag.Bool(in.IsBargain),
	}
}

func newLotMetrics(in core.LotMetrics) *models.LotMetrics {
	return &models.LotMetrics{
		MembersCount:   swag.Int64(int64(in.MembersCount)),
		DailyCoverage:  swag.Int64(int64(in.DailyCoverage)),
		MonthlyIncome:  nullIntToInt64(in.MonthlyIncome),
		PricePerMember: swag.Float64(in.PricePerMember),
		PricePerView:   swag.Float64(in.PricePerView),
		PaybackPeriod:  nullFloat64ToFloat64(in.PaybackPeriod),
	}
}

func newLotExtraResourceSlice(in []core.LotExtraResource) []*models.LotExtraResource {
	result := make([]*models.LotExtraResource, len(in))

	for i, v := range in {
		result[i] = &models.LotExtraResource{URL: v.URL}
	}

	return result
}

func NewOwnedLot(s storage.Storage, in *catalog.OwnedLot) *models.OwnedLot {

	var avatar string

	if in.Avatar.Valid {
		avatar = s.PublicURL(in.Avatar.String)
	}

	return &models.OwnedLot{
		ID:          swag.Int64(int64(in.ID)),
		ExternalID:  swag.Int64(in.ExternalID),
		Name:        swag.String(in.Name),
		Avatar:      swag.String(avatar),
		Link:        swag.String(in.Link()),
		Bio:         nullStringToString(in.Bio),
		Price:       newLotPrice(in.Price),
		Comment:     swag.String(in.Comment),
		Metrics:     newLotMetrics(in.Metrics),
		CreatedAt:   timeToUnix(in.CreatedAt),
		PaidAt:      nullTimeToUnix(in.PaidAt),
		ApprovedAt:  nullTimeToUnix(in.ApprovedAt),
		PublishedAt: nullTimeToUnix(in.PublishedAt),
		Topics: NewTopicSlice(in.Topics),
		Extra: newLotExtraResourceSlice(in.ExtraResources),
	}
}
