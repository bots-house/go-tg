package api

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/pkg/errors"

	webhookops "github.com/bots-house/birzzha/api/gen/restapi/operations/webhook"
	"github.com/bots-house/birzzha/api/models"
	"github.com/bots-house/birzzha/core"
)

func (h *Handler) handleGatewayNotification(params webhookops.HandleGatewayNotificationParams) middleware.Responder {
	gateway := h.Gateways.Get(params.Name)
	if gateway == nil {
		err := core.NewError("gateway_not_found", "gateway not found")
		return webhookops.NewHandleGatewayNotificationBadRequest().WithPayload(models.NewError(err))
	}

	ctx := params.HTTPRequest.Context()

	notify, err := gateway.ParseNotification(ctx, params.HTTPRequest)
	if err != nil {
		if err2, ok := errors.Cause(err).(*core.Error); ok {
			return webhookops.NewHandleGatewayNotificationBadRequest().WithPayload(models.NewError(err2))
		}
		return webhookops.NewHandleGatewayNotificationInternalServerError().WithPayload(models.NewInternalServerError(err))
	}

	if err := h.Personal.ProcessGatewayNotification(ctx, notify); err != nil {
		if err2, ok := errors.Cause(err).(*core.Error); ok {
			return webhookops.NewHandleGatewayNotificationBadRequest().WithPayload(models.NewError(err2))
		}
		return webhookops.NewHandleGatewayNotificationInternalServerError().WithPayload(models.NewInternalServerError(err))
	}

	return webhookops.NewHandleGatewayNotificationOK()
}
