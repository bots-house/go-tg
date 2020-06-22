package api

import (
	"github.com/bots-house/birzzha/api/gen/models"
	botops "github.com/bots-house/birzzha/api/gen/restapi/operations/bot"

	"github.com/bots-house/birzzha/pkg/log"
	"github.com/bots-house/birzzha/pkg/tg"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"github.com/tomasen/realip"
)

func (h *Handler) handleBotUpdate(params botops.HandleUpdateParams) middleware.Responder {
	ip := realip.FromRequest(params.HTTPRequest)

	if !tg.IsAllowedIP(ip) {
		return botops.NewHandleUpdateUnauthorized()
	}

	ctx := params.HTTPRequest.Context()

	update := &params.Payload

	if err := h.Bot.HandleUpdate(ctx, update); err != nil {
		log.Error(ctx, "handle update failed", "update_id", update.UpdateID, "err", err)
		return botops.NewHandleUpdateInternalServerError()
	}

	return botops.NewHandleUpdateOK()
}

func (h *Handler) getBotInfo(params botops.GetBotInfoParams) middleware.Responder {
	bot := h.Bot.Client().Self

	return botops.NewGetBotInfoOK().WithPayload(&models.BotInfo{
		Name:         swag.String(bot.FirstName),
		Username:     swag.String(bot.UserName),
		AuthDeepLink: swag.String("login"),
	})
}
