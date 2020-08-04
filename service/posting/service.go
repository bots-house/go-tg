package posting

import (
	"github.com/bots-house/birzzha/core"
	tgbotapi "github.com/bots-house/telegram-bot-api"
)

type Service struct {
	Lot      core.LotStore
	Settings core.SettingsStore
	Topic    core.TopicStore
	User     core.UserStore
	TgClient *tgbotapi.BotAPI
}
