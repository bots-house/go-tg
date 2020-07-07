package personal

import (
	"github.com/bots-house/birzzha/core"
	"github.com/bots-house/birzzha/pkg/storage"
	"github.com/bots-house/birzzha/pkg/tg"
	"github.com/bots-house/birzzha/service/admin"
	"github.com/bots-house/birzzha/service/payment"
	"github.com/bots-house/birzzha/store"
)

type Service struct {
	Lot               core.LotStore
	Resolver          tg.Resolver
	Payment           core.PaymentStore
	Txier             store.Txier
	AdminNotify       *admin.Notifications
	LotCanceledReason core.LotCanceledReasonStore

	Storage  storage.Storage
	Settings core.SettingsStore
	Gateways *payment.GatewayRegistry
}
