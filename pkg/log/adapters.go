package log

import (
	"github.com/go-kit/kit/log/level"
	"github.com/robfig/cron/v3"
)

type cronLogger struct {
	logger Logger
}

func (cl *cronLogger) Info(msg string, keysAndValues ...interface{}) {
	kvs := append([]interface{}{
		"msg", msg,
	}, keysAndValues...)

	_ = level.Info(cl.logger).Log(kvs...)
}

func (cl *cronLogger) Error(err error, msg string, keysAndValues ...interface{}) {
	kvs := append([]interface{}{
		"msg", msg,
		"err", err,
	}, keysAndValues...)

	_ = level.Error(cl.logger).Log(kvs...)
}

func NewCronLogger(logger Logger) cron.Logger {
	return &cronLogger{logger}
}
