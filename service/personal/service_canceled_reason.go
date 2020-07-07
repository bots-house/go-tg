package personal

import (
	"context"

	"github.com/bots-house/birzzha/core"
)

func (srv *Service) GetLotCanceledReasons(ctx context.Context) ([]*core.LotCanceledReason, error) {
	return srv.LotCanceledReason.Query().All(ctx)
}
