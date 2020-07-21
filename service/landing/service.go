package landing

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/bots-house/birzzha/core"
	"github.com/bots-house/birzzha/pkg/tg"
	"github.com/pkg/errors"
)

type Service struct {
	Review   core.ReviewStore
	Settings core.SettingsStore
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
	Channel Channel
	Stats   Stats
	Reviews *ReviewList
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

	now := time.Now()
	result, err := srv.Resolver.ResolveByID(ctx, settings.Channel.PrivateID)
	if err != nil {
		return nil, errors.Wrap(err, "resolve by id")
	}
	fmt.Println(time.Since(now))

	if result.Channel == nil {
		return nil, ErrChannelNotFound
	}

	return &Landing{
		Stats: Stats{
			UniqueVisitorsPerMonth: rand.Intn(10000) + 1,
			AvgLotChannelReach:     rand.Intn(10000) + 1,
			AvgLotSiteReach:        rand.Intn(10000) + 1,
		},
		Reviews: reviews,
		Channel: Channel{
			Title:        result.Channel.Name,
			MembersCount: result.Channel.MembersCount,
			JoinLink:     settings.Channel.PrivateLink,
			Avatar:       result.Channel.Avatar,
			Username:     settings.Channel.PublicUsername,
		},
	}, nil
}
