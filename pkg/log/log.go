package log

import (
	"context"
	"os"
	"time"

	"github.com/bots-house/birzzha/core"
	"github.com/getsentry/sentry-go"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/go-kit/kit/log/term"
)

func NewLogger(debug bool, colors bool) log.Logger {
	var logger log.Logger

	writer := log.NewSyncWriter(os.Stdout)
	if colors {
		logger = term.NewLogger(writer, log.NewLogfmtLogger, loggerColorFn)
	} else {
		logger = log.NewLogfmtLogger(writer)
	}

	logger = log.With(logger, "ts", log.TimestampFormat(time.Now, "15:04:05"))

	if !debug {
		logger = level.NewFilter(logger, level.AllowInfo())
	} else {
		logger = level.NewFilter(logger, level.AllowDebug())
	}

	return logger
}

// With creates a context with new key values.
func With(ctx context.Context, kvs ...interface{}) context.Context {
	logger := GetLogger(ctx)
	logger = log.With(logger, kvs...)

	return WithLogger(ctx, logger)
}

// WithPrefix create a context with prefix.
func WithPrefix(ctx context.Context, kvs ...interface{}) context.Context {
	logger := GetLogger(ctx)
	logger = log.WithPrefix(logger, kvs...)

	return WithLogger(ctx, logger)
}

// Log message.
func Log(ctx context.Context, msg string, kvs ...interface{}) {
	kvs = append([]interface{}{
		"msg",
		msg,
	}, kvs...)

	_ = GetLogger(ctx).Log(kvs...)
}

// Debug message
func Debug(ctx context.Context, msg string, kvs ...interface{}) {
	kvs = append([]interface{}{
		"msg",
		msg,
	}, kvs...)

	_ = level.Debug(GetLogger(ctx)).Log(kvs...)
}

// Info message
func Info(ctx context.Context, msg string, kvs ...interface{}) {
	kvs = append([]interface{}{
		"msg",
		msg,
	}, kvs...)

	_ = level.Info(GetLogger(ctx)).Log(kvs...)
}

// Error message
func Error(ctx context.Context, msg string, kvs ...interface{}) {
	kvs = append([]interface{}{
		"msg",
		msg,
	}, kvs...)
	_ = level.Error(GetLogger(ctx)).Log(kvs...)
	captureException(ctx, msg, kvs)
}

// Warn message
func Warn(ctx context.Context, msg string, kvs ...interface{}) {
	kvs = append([]interface{}{
		"msg",
		msg,
	}, kvs...)

	_ = level.Warn(GetLogger(ctx)).Log(kvs...)
}

func captureException(ctx context.Context, msg string, kvs []interface{}) {
	if !sentry.HasHubOnContext(ctx) {
		return
	}
	hub := sentry.GetHubFromContext(ctx)
	module, ok := ctx.Value(core.ModuleName).(string)
	if !ok {
		module = core.Unknown
	}
	hub.WithScope(func(scope *sentry.Scope) {
		scope.SetTag(string(core.ModuleName), module)
		for i := 0; i < len(kvs); i += 2 {
			if key, ok := kvs[i].(string); ok {
				switch value := kvs[i+1].(type) {
				case error:
					scope.SetExtra(key, value.Error())
					msg += " " + value.Error()
				default:
					scope.SetExtra(key, kvs[i+1])
				}
			}
		}
		scope.SetLevel(sentry.LevelError)
		hub.CaptureMessage(msg)
	})
}
