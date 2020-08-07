package worker

import (
	"context"
	"time"

	"github.com/bots-house/birzzha/core"
	"github.com/bots-house/birzzha/pkg/log"
	"github.com/bots-house/birzzha/pkg/stat"
	"github.com/bots-house/birzzha/pkg/storage"
	"github.com/bots-house/birzzha/pkg/tg"
	"github.com/bots-house/birzzha/service/posting"
	"github.com/pkg/errors"
	"github.com/robfig/cron/v3"
)

type Worker struct {
	Landing  core.LandingStore
	Lot      core.LotStore
	Settings core.SettingsStore

	Posting *posting.Service

	Resolver tg.Resolver

	Storage storage.Storage

	SiteStat     stat.Site
	TelegramStat stat.Telegram

	Location *time.Location

	UpdateLandingSpec string
	PublishPostsSpec  string
	UpdateLotsSpec    string

	cron *cron.Cron
}

func (wrk *Worker) wrap(ctx context.Context, do func(ctx context.Context) error) func() {
	return func() {
		if err := do(ctx); err != nil {
			panic(err)
		}
	}
}

func (wrk *Worker) setupCron(ctx context.Context) error {
	logger := log.NewCronLogger(log.GetLogger(ctx))

	opts := []cron.Option{
		cron.WithChain(
			cron.DelayIfStillRunning(logger),
			cron.Recover(logger),
		),
	}

	if wrk.Location != nil {
		log.Debug(ctx, "worker time zone set", "timezone", wrk.Location)
		opts = append(opts, cron.WithLocation(wrk.Location))
	}

	crn := cron.New(opts...)

	wrk.cron = crn

	if err := wrk.setupCronJobs(ctx); err != nil {
		return errors.Wrap(err, "setup cron jobs")
	}

	return nil
}

func (wrk *Worker) setupCronJobs(ctx context.Context) error {
	for _, job := range []struct {
		Name string
		Spec string
		Do   func()
	}{
		{
			"update lots",
			wrk.UpdateLotsSpec,
			wrk.wrap(ctx, wrk.taskUpdateLotList),
		},
		{
			"update landing",
			wrk.UpdateLandingSpec,
			wrk.wrap(ctx, wrk.taskUpdateLanding),
		},
		{
			"publish posts",
			wrk.PublishPostsSpec,
			wrk.wrap(ctx, wrk.taskPublishPosts),
		},
	} {
		log.Debug(ctx, "setup worker cron job", "name", job.Name, "spec", job.Spec)

		_, err := wrk.cron.AddFunc(job.Spec, job.Do)
		if err != nil {
			return errors.Wrapf(err, "setup job %s, %s", job.Name, job.Spec)
		}
	}

	return nil
}

func (wrk *Worker) run(ctx context.Context) {
	log.Info(ctx, "start worker")
	wrk.cron.Run()
}

func (wrk *Worker) stop(ctx context.Context) {
	stopCtx := wrk.cron.Stop()

	log.Debug(ctx, "wait until worker shutdown")
	<-stopCtx.Done()
}

// Run blocking worker.
func (wrk *Worker) Run(ctx context.Context) error {
	ctx = log.WithPrefix(ctx, "scope", "worker")

	if err := wrk.setupCron(ctx); err != nil {
		return errors.Wrap(err, "setup")
	}

	// translate app context close to worker stop
	go func() {
		<-ctx.Done()
		log.Info(ctx, "shutdown worker")
		wrk.stop(ctx)
	}()

	wrk.run(ctx)

	return nil
}
