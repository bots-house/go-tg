package catalog

import (
	"context"

	"github.com/bots-house/birzzha/pkg/tg"
)

// ResolveTelegram to Telegram entity from @username or links (private too).
func (srv *Service) ResolveTelegram(ctx context.Context, query string) (*tg.ResolveResult, error) {
	return srv.Resolver.Resolve(ctx, query)
}
