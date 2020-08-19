package worker

import (
	"context"
	"math"
	"strconv"
	"sync"

	"github.com/bots-house/birzzha/core"
	"github.com/bots-house/birzzha/pkg/log"
	"github.com/pkg/errors"
)

func (wrk *Worker) taskUpdateLandingAvgSiteReach(ctx context.Context, landing *core.Landing) error {
	views, err := wrk.Lot.AverageSiteViews(ctx)
	if err != nil {
		return errors.Wrap(err, "count average views")
	}

	landing.AvgSiteReachActual = int(math.Round(views))

	return nil
}

func (wrk *Worker) taskUpdateLandingUniqueUsersPerMonth(ctx context.Context, landing *core.Landing) error {
	visitors, err := wrk.SiteStat.GetUniqueVisitorsPerMonth(ctx)
	if err != nil {
		return errors.Wrap(err, "count unique visitors")
	}

	landing.UniqueUsersPerMonthActual = visitors

	return nil
}

func (wrk *Worker) taskUpdateLandingAvgChannelReach(ctx context.Context, landing *core.Landing) error {
	settings, err := wrk.Settings.Get(ctx)
	if err != nil {
		return errors.Wrap(err, "get settings")
	}

	id := core.MTProtoPrivateID(settings.Channel.PrivateID)

	stats, err := wrk.TelegramStat.Get(ctx, strconv.FormatInt(id, 10))
	if err != nil {
		return errors.Wrap(err, "get telegram stats")
	}

	landing.AvgChannelReachActual = stats.ViewsPerPostAvg

	return nil
}

func (wrk *Worker) taskUpdateLanding(ctx context.Context) error {
	ctx = log.With(ctx, "task", "update_landing")

	landing, err := wrk.Landing.Get(ctx)
	if err != nil {
		return errors.Wrap(err, "get landing")
	}

	wg := sync.WaitGroup{}

	for _, subtask := range []struct {
		Name string
		Task func(ctx context.Context, landing *core.Landing) error
	}{
		{"avg_channel_reach", wrk.taskUpdateLandingAvgChannelReach},
		{"avg_site_reach", wrk.taskUpdateLandingAvgSiteReach},
		{"unique_users_per_month", wrk.taskUpdateLandingUniqueUsersPerMonth},
	} {
		subtask := subtask

		wg.Add(1)
		go func() {
			defer wg.Done()

			log.Debug(ctx, "run subtask", "subtask", subtask.Name)
			if err := subtask.Task(ctx, landing); err != nil {
				log.Error(ctx, "subtask error", "subtask", subtask.Name, "err", err)
			}
		}()
	}

	wg.Wait()

	if err := wrk.Landing.Update(ctx, landing); err != nil {
		return errors.Wrap(err, "update landing in store")
	}

	return nil
}
