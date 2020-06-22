package api

import (
	"github.com/bots-house/birzzha/api/gen/restapi/operations"
	"github.com/bots-house/birzzha/pkg/log"
	"github.com/bots-house/birzzha/pkg/tg"
	"github.com/go-openapi/runtime/middleware"
	"github.com/tomasen/realip"
)

func (h *Handler) handleBotUpdate(params operations.HandleBotUpdateParams) middleware.Responder {
	ip := realip.FromRequest(params.HTTPRequest)

	if !tg.IsAllowedIP(ip) {
		return operations.NewHandleBotUpdateUnauthorized()
	}

	ctx := params.HTTPRequest.Context()

	update := &params.Payload

	if err := h.Bot.HandleUpdate(ctx, update); err != nil {
		log.Error(ctx, "handle update failed", "update_id", update.UpdateID, "err", err)
		return operations.NewHandleBotUpdateInternalServerError()
	}

	return operations.NewHandleBotUpdateOK()
}
