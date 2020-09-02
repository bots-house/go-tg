package personal

import (
	"github.com/bots-house/birzzha/core"
	"github.com/bots-house/birzzha/pkg/stat"
	"github.com/bots-house/birzzha/pkg/storage"
	"github.com/bots-house/birzzha/pkg/tg"
	"github.com/bots-house/birzzha/service/notifications"
	"github.com/bots-house/birzzha/service/payment"
	"github.com/bots-house/birzzha/store"
)

type Service struct {
	Lot               core.LotStore
	LotFavorite       core.FavoriteStore
	Resolver          tg.Resolver
	Payment           core.PaymentStore
	Txier             store.Txier
	Notify            *notifications.Notifications
	LotCanceledReason core.LotCanceledReasonStore
	LotFile           core.LotFileStore
	TelegramStat      stat.Telegram

	Storage  storage.Storage
	Settings core.SettingsStore
	Gateways *payment.GatewayRegistry
	Parser   core.LotExtraResourceParser

	AdminNotificationsChannelID int64
}
