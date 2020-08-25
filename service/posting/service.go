package posting

import (
	"github.com/bots-house/birzzha/core"
	"github.com/bots-house/birzzha/pkg/notifications"
	"github.com/bots-house/birzzha/store"
	tgbotapi "github.com/bots-house/telegram-bot-api"
)

type Config struct {
	Site string
}

type Service struct {
	Lot      core.LotStore
	Settings core.SettingsStore
	Topic    core.TopicStore
	User     core.UserStore
	Post     core.PostStore

	Config Config

	Txier store.Txier

	TgClient *tgbotapi.BotAPI

	UserNotification notifications.UserNotification

	SiteWithPathListChannel string
}
