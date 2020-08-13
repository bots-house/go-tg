package models

import (
	"github.com/Rhymond/go-money"
	"github.com/bots-house/birzzha/api/gen/models"
	"github.com/bots-house/birzzha/core"
	"github.com/bots-house/birzzha/service/admin"
	"github.com/go-openapi/swag"
)

func NewSettings(in *admin.FullSettings) *models.AdminSettings {
	return &models.AdminSettings{
		Prices: &models.AdminSettingsPrices{
			Application: newMoney(in.Prices.Application),
			Change:      newMoney(in.Prices.Change),
			Cashier:     swag.String(in.CashierUsername),
		},
		Garant: NewSettingsGarant(in.Settings),
		Channel: &models.AdminSettingsChannel{
			ID:             swag.Int64(in.Channel.PrivateID),
			JoinLink:       swag.String(in.Channel.PrivateLink),
			PublicUsername: swag.String(in.Channel.PublicUsername),
		},
		Topics:             NewAdminFullTopicSlice(in.Topics),
		LotCanceledReasons: NewLotCanceledReasonSlice(in.CanceledReasons),
		Landing:            NewAdminLanding(in.Landing),
	}
}

func NewSettingsPrices(in *core.Settings) *models.AdminSettingsPrices {
	return &models.AdminSettingsPrices{
		Application: newMoney(in.Prices.Application),
		Change:      newMoney(in.Prices.Change),
		Cashier:     swag.String(in.CashierUsername),
	}
}

func NewSettingsChannel(in *core.Settings) *models.AdminSettingsChannel {
	return &models.AdminSettingsChannel{
		ID:             swag.Int64(in.Channel.PrivateID),
		JoinLink:       swag.String(in.Channel.PrivateLink),
		PublicUsername: swag.String(in.Channel.PublicUsername),
	}
}

func NewSettingsGarant(in *core.Settings) *models.AdminSettingsGarant {
	return &models.AdminSettingsGarant{
		Name:                           swag.String(in.Garant.Name),
		Username:                       swag.String(in.Garant.Username),
		ReviewsChannel:                 swag.String(in.Garant.ReviewsChannel),
		AvatarURL:                      nullStringToString(in.Garant.AvatarURL),
		PercentageDeal:                 swag.Float64(in.Garant.PercentageDeal),
		PercentageDealOfDiscountPeriod: nullFloat64ToFloat64(in.Garant.PercentageDealDiscountPeriod),
	}
}

func ToMoney(in *models.Money) *money.Money {
	return money.New(
		int64(swag.Float64Value(in.Amount)*100.0),
		swag.StringValue(in.Currency),
	)
}
