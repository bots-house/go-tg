package views

import (
	"time"

	"github.com/bots-house/birzzha/core"
	"github.com/bots-house/birzzha/pkg/kv"
	"github.com/bots-house/birzzha/store"
)

type Service struct {
	Lot                core.LotStore
	Txier              store.Txier
	KV                 kv.Store
	SiteViewExpiration time.Duration
}
