package api

import (
	webhookops "github.com/bots-house/birzzha/api/gen/restapi/operations/webhook"
	"github.com/bots-house/birzzha/api/models"
	"github.com/bots-house/birzzha/core"
	"github.com/go-openapi/runtime/middleware"
	"github.com/pkg/errors"
)

func (h *Handler) handleGatewayUnitpayNotification(params webhookops.HandleGatewayUnitpayNotificationParams) middleware.Responder {
	gateway := h.Gateways.Get(params.Name)
	if gateway == nil {
		err := core.NewError("gateway_not_found", "gateway not found")
		return webhookops.NewHandleGatewayNotificationBadRequest().WithPayload(models.NewError(err))
	}

	ctx := params.HTTPRequest.Context()

	notify, err := gateway.ParseNotification(ctx, params.HTTPRequest)
	if err != nil {
		if err2, ok := errors.Cause(err).(*core.Error); ok {
			return webhookops.NewHandleGatewayUnitpayNotificationBadRequest().WithPayload(models.NewUnitPayError(err2))
		}
		return webhookops.NewHandleGatewayUnitpayNotificationInternalServerError().WithPayload(models.NewUnitPayInternalError(ctx, err))
	}

	msg, err := h.Personal.ProcessUnitPayNotification(ctx, notify)
	if err != nil {
		if err2, ok := errors.Cause(err).(*core.Error); ok {
			return webhookops.NewHandleGatewayUnitpayNotificationBadRequest().WithPayload(models.NewUnitPayError(err2))
		}
		return webhookops.NewHandleGatewayUnitpayNotificationInternalServerError().WithPayload(models.NewUnitPayInternalError(ctx, err))
	}

	return webhookops.NewHandleGatewayUnitpayNotificationOK().WithPayload(models.NewUnitPaySuccessResp(msg))
}

func (h *Handler) handlerGatewayNotification(params webhookops.HandleGatewayNotificationParams) middleware.Responder {
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
		return webhookops.NewHandleGatewayNotificationInternalServerError().WithPayload(models.NewInternalServerError(ctx, err))
	}

	if err := h.Personal.ProcessGatewayNotification(ctx, notify); err != nil {
		if err2, ok := errors.Cause(err).(*core.Error); ok {
			return webhookops.NewHandleGatewayNotificationBadRequest().WithPayload(models.NewError(err2))
		}
		return webhookops.NewHandleGatewayNotificationInternalServerError().WithPayload(models.NewInternalServerError(ctx, err))
	}

	return webhookops.NewHandleGatewayNotificationOK()
}
