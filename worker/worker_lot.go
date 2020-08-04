package worker

import (
	"context"
	"strconv"
	"time"

	"github.com/bots-house/birzzha/core"
	"github.com/bots-house/birzzha/pkg/log"
	"github.com/bots-house/birzzha/pkg/stat"
	"github.com/bots-house/birzzha/pkg/storage"
	"github.com/pkg/errors"
	"github.com/volatiletech/null/v8"
)

func (wrk *Worker) taskUpdateLotList(ctx context.Context) error {
	ctx = log.With(ctx, "task", "update_lot_list")
	lots, err := wrk.Lot.Query().Statuses(core.LotStatusPublished).All(ctx)
	if err != nil {
		return errors.Wrap(err, "get published lots")
	}

	for _, lot := range lots {
		lot := lot
		log.Debug(ctx, "updating lot", "lot_id", lot.ID)

		if err := wrk.taskUpdateLot(ctx, lot); err != nil {
			log.Error(ctx, "update lot task", "error", "lot_id", lot.ID)
			continue
		}
	}

	return nil
}

func (wrk *Worker) taskUpdateLot(ctx context.Context, lot *core.Lot) error {
	//must wait so we won't run out of rate limiter
	defer time.Sleep(time.Millisecond * 500)

	result, err := wrk.Resolver.ResolveByID(ctx, lot.ExternalID)
	if err != nil {
		return errors.Wrapf(err, "failed to resolve lot by external id %d", lot.ExternalID)
	}
	channel := result.Channel

	if channel.Username != "" {
		lot.Username = null.NewString(channel.Username, channel.Username != "")
		lot.JoinLink = null.NewString("", false)
	}

	wrk.addAvatar(ctx, channel.Avatar, lot)

	var dailyCoverage int
	tStat, err := wrk.TelegramStat.Get(ctx, strconv.FormatInt(core.MTProtoPrivateID(lot.ExternalID), 10))
	switch {
	case err == nil:
		dailyCoverage = tStat.ViewsPerPostDaily
	default:
		dailyCoverage = lot.Metrics.DailyCoverage
		if err.Error() != stat.ErrChannelNotFound.Error() {
			log.Error(ctx, "failed to get telegram stat while updating lot", "error", err)
		}
	}

	lot.Metrics = core.NewLotMetrics(lot.Price.Current, channel.MembersCount, dailyCoverage, lot.Metrics.MonthlyIncome)

	return errors.Wrapf(wrk.Lot.Update(ctx, lot), "update lot in store lot_id=%d and error=%v", lot.ID, err)
}

func (wrk *Worker) addAvatar(ctx context.Context, url string, lot *core.Lot) {
	if url == "" {
		return
	}
	avatar, err := wrk.Storage.AddByURL(ctx, storage.LotDir, url)
	if err != nil {
		log.Error(ctx, "add avatar by url (worker)", "error", err, "url", url, "lot_id", lot.ID)
		return
	}
	lot.Avatar = null.NewString(avatar, true)
}
