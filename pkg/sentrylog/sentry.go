package sentrylog

import (
	"context"
	"time"

	"github.com/bots-house/birzzha/core"

	"github.com/bots-house/birzzha/pkg/log"
	"github.com/getsentry/sentry-go"
	"github.com/pkg/errors"
)

const sentryFlushTimeout = time.Second * 5

func Init(ctx context.Context, revision, env, dsn string) (func(), error) {
	if core.EnvLocal == env {
		log.Warn(ctx, "sentry are not initialized")
		return func() {}, nil
	}
	if err := sentry.Init(sentry.ClientOptions{
		Dsn:              dsn,
		AttachStacktrace: true,
		Release:          revision,
		Environment:      env,
	}); err != nil {
		return nil, errors.Wrap(err, "init sentry")
	}
	log.Info(ctx, "sentry initialized successfully", "dsn", dsn)

	return func() {
		if !sentry.Flush(sentryFlushTimeout) {
			log.Error(ctx, "timeout was reached. Some sentry events may not have been sent", "error", errors.New("sentry flush timeout"))
		}
	}, nil
}
