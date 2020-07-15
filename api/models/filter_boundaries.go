package models

import (
	"github.com/go-openapi/swag"

	"github.com/bots-house/birzzha/api/gen/models"
	"github.com/bots-house/birzzha/core"
)

func NewFilterBoundaries(in *core.LotFilterBoundaries) *models.LotFilterBoundaries {
	return &models.LotFilterBoundaries{
		MembersCount: &models.IntBoundaries{
			From: swag.Int64(int64(in.MembersCountMin)),
			To:   swag.Int64(int64(in.MembersCountMax)),
		},
		Price: &models.IntBoundaries{
			From: swag.Int64(int64(in.PriceMin)),
			To:   swag.Int64(int64(in.PriceMax)),
		},
		PricePerMember: &models.FloatBoundaries{
			From: swag.Float64(in.PricePerMemberMin),
			To:   swag.Float64(in.PricePerMemberMax),
		},
		PricePerView: &models.FloatBoundaries{
			From: swag.Float64(in.PricePerViewMin),
			To:   swag.Float64(in.PricePerViewMax),
		},
		DailyCoverage: &models.IntBoundaries{
			From: swag.Int64(int64(in.DailyCoverageMin)),
			To:   swag.Int64(int64(in.DailyCoverageMax)),
		},
		MonthlyIncome: &models.IntBoundaries{
			From: swag.Int64(int64(in.MonthlyIncomeMin)),
			To:   swag.Int64(int64(in.MonthlyIncomeMax)),
		},
		PaybackPeriod: &models.FloatBoundaries{
			From: swag.Float64(in.PaybackPeriodMin),
			To:   swag.Float64(in.PaybackPeriodMax),
		},
	}
}
