package core

import (
	"context"

	"github.com/Rhymond/go-money"
	"github.com/volatiletech/null/v8"
)

type SettingsPrices struct {
	Application *money.Money
	Change      *money.Money
}

type SettingsChannel struct {
	PublicUsername string
	PrivateLink    string
	PrivateID      int64
}

type Settings struct {
	Prices          SettingsPrices
	Channel         SettingsChannel
	CashierUsername string
	UpdatedAt       null.Time
	UpdatedBy       UserID
}

var DefaultSettings = &Settings{
	Prices: SettingsPrices{
		Application: money.New(800*100, "RUB"),
		Change:      money.New(200*100, "RUB"),
	},
	CashierUsername: "alexxdd",
	Channel: SettingsChannel{
		PublicUsername: "birzzha",
		PrivateLink:    "https://t.me/joinchat/AAAAAENM1m0f_WHVNXjP4w",
		PrivateID:      -1001129109101,
	},
}

type SettingsStore interface {
	Get(ctx context.Context) (*Settings, error)
	Update(ctx context.Context, settings *Settings) error
}
