package stat

import (
	"context"

	"github.com/bots-house/birzzha/core"
)

var (
	ErrChannelNotFound = core.NewError("channel_not_found", "channel not found")
)

type TelegramStats struct {
	ViewsPerPostAvg   int
	ViewsPerPostDaily int
}

type Telegram interface {
	Get(ctx context.Context, query string) (*TelegramStats, error)
}
