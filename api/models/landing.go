package models

import (
	"github.com/bots-house/birzzha/api/gen/models"
	"github.com/bots-house/birzzha/core"
	"github.com/bots-house/birzzha/pkg/storage"
	"github.com/bots-house/birzzha/service/landing"
	"github.com/go-openapi/swag"
)

func NewLanding(s storage.Storage, in *landing.Landing) *models.Landing {
	landing := &models.Landing{
		Garant: &models.AdminSettingsGarant{
			Name:                           swag.String(in.Garant.Name),
			Username:                       swag.String(in.Garant.Username),
			ReviewsChannel:                 swag.String(in.Garant.ReviewsChannel),
			PercentageDeal:                 swag.Float64(in.Garant.PercentageDeal),
			PercentageDealOfDiscountPeriod: in.Garant.PercentageDealDiscountPeriod.Ptr(),
			AvatarURL:                      in.Garant.AvatarURL.Ptr(),
		},
		Stats: &models.LandingStats{
			UniqueVisitorsPerMonth: swag.Int64(int64(in.Stats.UniqueVisitorsPerMonth)),
			AvgLotChannelReach:     swag.Int64(int64(in.Stats.AvgLotChannelReach)),
			AvgLotSiteReach:        swag.Int64(int64(in.Stats.AvgLotSiteReach)),
		},
		Application: newMoney(in.ApplicationPrice),
		Reviews:     NewReviewList(s, in.Reviews),
	}

	if in.Channel != nil {
		landing.Channel = &models.LandingChannel{
			Title:          swag.String(in.Channel.Title),
			MembersCount:   swag.Int64(int64(in.Channel.MembersCount)),
			JoinLink:       swag.String(in.Channel.JoinLink),
			Avatar:         swag.String(in.Channel.Avatar),
			PublicUsername: swag.String(in.Channel.Username),
		}
	}
	return landing
}

func NewAdminLanding(in *core.Landing) *models.AdminSettingsLanding {
	return &models.AdminSettingsLanding{
		UniqueUsersPerMonthActual: swag.Int64(int64(in.UniqueUsersPerMonthActual)),
		UniqueUsersPerMonthShift:  swag.Int64(int64(in.UniqueUsersPerMonthShift)),
		AvgChannelReachShift:      swag.Int64(int64(in.AvgChannelReachShift)),
		AvgChannelReachActual:     swag.Int64(int64(in.AvgChannelReachActual)),
		AvgSiteReachActual:        swag.Int64(int64(in.AvgSiteReachActual)),
		AvgSiteReachShift:         swag.Int64(int64(in.AvgChannelReachShift)),
	}
}
