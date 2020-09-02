package core

import (
	"context"
	"fmt"

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

func (settings *SettingsChannel) PostLink(id int) string {
	return fmt.Sprintf("https://t.me/%s/%d",
		settings.PublicUsername,
		id,
	)
}

func (settings *SettingsChannel) Link() string {
	if settings.PrivateLink != "" {
		return settings.PrivateLink
	} else {
		return settings.PublicUsername
	}
}

type Garant struct {
	Name                         string
	Username                     string
	ReviewsChannel               string
	PercentageDeal               float64
	PercentageDealDiscountPeriod null.Float64
	AvatarURL                    null.String
}

type Settings struct {
	Garant          Garant
	Prices          SettingsPrices
	Channel         SettingsChannel
	CashierUsername string
	UpdatedAt       null.Time
	UpdatedBy       UserID
}

var DefaultSettings = &Settings{
	Garant: Garant{
		Name:           "Alex Dalakian",
		Username:       "alexxdd",
		ReviewsChannel: "birzzha_review",
		PercentageDeal: 4,
	},
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
