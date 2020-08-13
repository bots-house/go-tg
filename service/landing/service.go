package landing

import (
	"context"

	"github.com/Rhymond/go-money"
	"github.com/bots-house/birzzha/core"
	"github.com/bots-house/birzzha/pkg/tg"
	"github.com/pkg/errors"
)

type Service struct {
	Review   core.ReviewStore
	Settings core.SettingsStore
	Landing  core.LandingStore

	Resolver tg.Resolver
}

type Channel struct {
	Title        string
	MembersCount int
	JoinLink     string
	Avatar       string
	Username     string
}

type Stats struct {
	UniqueVisitorsPerMonth int
	AvgLotSiteReach        int
	AvgLotChannelReach     int
}

type Landing struct {
	Garant           *core.Garant
	Channel          Channel
	ApplicationPrice *money.Money
	Stats            Stats
	Reviews          *ReviewList
}

var (
	ErrChannelNotFound = core.NewError("channel_not_found", "channel not found")
)

func (srv *Service) GetLanding(ctx context.Context) (*Landing, error) {
	settings, err := srv.Settings.Get(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get settings")
	}

	reviews, err := srv.GetReviews(ctx, 0, 0)
	if err != nil {
		return nil, errors.Wrap(err, "get reviews")
	}

	result, err := srv.Resolver.ResolveByID(ctx, settings.Channel.PrivateID)
	if err != nil {
		return nil, errors.Wrap(err, "resolve by id")
	}

	if result.Channel == nil {
		return nil, ErrChannelNotFound
	}

	landing, err := srv.Landing.Get(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get landing")
	}
	return &Landing{
		Stats: Stats{
			UniqueVisitorsPerMonth: landing.UniqueUsersPerMonth(),
			AvgLotChannelReach:     landing.AvgChannelReach(),
			AvgLotSiteReach:        landing.AvgSiteReach(),
		},
		Garant: &core.Garant{
			Name:                         settings.Garant.Name,
			Username:                     settings.Garant.Username,
			ReviewsChannel:               settings.Garant.ReviewsChannel,
			PercentageDeal:               settings.Garant.PercentageDeal,
			PercentageDealDiscountPeriod: settings.Garant.PercentageDealDiscountPeriod,
			AvatarURL:                    settings.Garant.AvatarURL,
		},
		ApplicationPrice: settings.Prices.Application,
		Reviews:          reviews,
		Channel: Channel{
			Title:        result.Channel.Name,
			MembersCount: result.Channel.MembersCount,
			JoinLink:     settings.Channel.PrivateLink,
			Avatar:       result.Channel.Avatar,
			Username:     settings.Channel.PublicUsername,
		},
	}, nil
}
