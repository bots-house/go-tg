package api

import (
	botops "github.com/bots-house/birzzha/api/gen/restapi/operations/bot"

	"github.com/bots-house/birzzha/pkg/log"
	"github.com/bots-house/birzzha/pkg/tg"
	"github.com/go-openapi/runtime/middleware"
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
