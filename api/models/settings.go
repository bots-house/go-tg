package models

import (
	"github.com/Rhymond/go-money"
	"github.com/bots-house/birzzha/api/gen/models"
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
		Channel: &models.AdminSettingsChannel{
			ID:             swag.Int64(in.Channel.PrivateID),
			JoinLink:       swag.String(in.Channel.PrivateLink),
			PublicUsername: swag.String(in.Channel.PublicUsername),
		},
		Topics:             NewTopicSlice(in.Topics),
		LotCanceledReasons: NewLotCanceledReasonSlice(in.CanceledReasons),
	}
}

func ToMoney(in *models.Money) *money.Money {
	return money.New(
		int64(swag.Float64Value(in.Amount)*100.0),
		swag.StringValue(in.Currency),
	)
}
