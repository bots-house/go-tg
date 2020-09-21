package api

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/tomasen/realip"

	botops "github.com/bots-house/birzzha/api/gen/restapi/operations/bot"
	"github.com/bots-house/birzzha/api/models"

	"github.com/bots-house/birzzha/pkg/log"
	"github.com/bots-house/birzzha/pkg/tg"
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
		return botops.NewHandleUpdateOK()
	}

	return botops.NewHandleUpdateOK()
}

func (h *Handler) getBotInfo(params botops.GetBotInfoParams) middleware.Responder {
	bot := h.Bot.Client().Self

	return botops.NewGetBotInfoOK().WithPayload(models.NewBotInfo(&bot))
}

//func (h *Handler) getBotFile(params botops.GetFileParams) middleware.Responder {
//	ctx := params.HTTPRequest.Context()
//
//	rc, err := h.BotFileProxy.Get(ctx, params.ID)
//	if err != nil {
//		if err2, ok := errors.Cause(err).(*core.Error); ok {
//			return botops.NewGetFileBadRequest().WithPayload(models.NewError(err2))
//		}
//		return botops.NewGetFileInternalServerError().WithPayload(models.NewInternalServerError(err))
//	}
//
//	return botops.NewGetFileOK().WithPayload(rc)
//}
