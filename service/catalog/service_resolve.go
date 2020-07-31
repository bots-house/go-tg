package catalog

import (
	"context"
	"strconv"
	"time"

	"github.com/bots-house/birzzha/core"

	"github.com/bots-house/birzzha/pkg/tg"
)

// ResolveTelegram to Telegram entity from @username or links (private too).
func (srv *Service) ResolveTelegram(ctx context.Context, query string) (*tg.ResolveResult, error) {
	return srv.Resolver.Resolve(ctx, query)
}

func (srv *Service) GetDailyCoverage(channel int64) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	stats, err := srv.TelegramStat.Get(ctx, strconv.FormatInt(core.MTProtoPrivateID(channel), 10))
	if err != nil {
		return 0, err
	}

	return stats.ViewsPerPostDaily, nil
}
