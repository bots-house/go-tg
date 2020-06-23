package models

import (
	tgbotapi "github.com/bots-house/telegram-bot-api"
	"github.com/go-openapi/swag"

	"github.com/bots-house/birzzha/api/gen/models"
)

func NewBotInfo(bot *tgbotapi.User) *models.BotInfo {
	return &models.BotInfo{
		Name:         swag.String(bot.FirstName),
		Username:     swag.String(bot.UserName),
		AuthDeepLink: swag.String("login"),
	}
}
