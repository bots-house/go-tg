package stat

import "context"

type TelegramStats struct {
	ViewsPerPostAvg   int
	ViewsPerPostDaily int
}

type Telegram interface {
	Get(ctx context.Context, query string) (*TelegramStats, error)
}
