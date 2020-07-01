package api

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"github.com/pkg/errors"

	"github.com/bots-house/birzzha/api/authz"
	personalops "github.com/bots-house/birzzha/api/gen/restapi/operations/personal_area"
	"github.com/bots-house/birzzha/api/models"
	"github.com/bots-house/birzzha/core"
	"github.com/bots-house/birzzha/service/personal"
)

func (h *Handler) createLot(params personalops.CreateLotParams, identity *authz.Identity) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	payload := params.Payload

	topics := make([]core.TopicID, len(payload.Topics))
	for i, v := range payload.Topics {
		topics[i] = core.TopicID(v)
	}

	in := &personal.LotInput{
		Query:         swag.StringValue(payload.Query),
		TelegramID:    swag.Int64Value(payload.TelegramID),
		TopicIDs:      topics,
		Price:         int(swag.Int64Value(payload.Price)),
		IsBargain:     swag.BoolValue(payload.IsBargain),
		MonthlyIncome: int(swag.Int64Value(payload.MonthlyIncome)),
		Comment:       swag.StringValue(payload.Comment),
		Extra:         payload.Extra,
	}

	lot, err := h.Personal.AddLot(ctx, identity.User, in)
	if err != nil {
		if err2, ok := errors.Cause(err).(*core.Error); ok {
			return personalops.NewCreateLotBadRequest().WithPayload(models.NewError(err2))
		}
		return personalops.NewCreateLotInternalServerError().WithPayload(models.NewInternalServerError(err))
	}

	return personalops.NewCreateLotCreated().WithPayload(models.NewOwnedLot(h.Storage, lot))
}

func (h *Handler) getApplicationInvoice(params personalops.GetApplicationInoviceParams, identity *authz.Identity) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	id := core.LotID(params.ID)

	invoice, err := h.Personal.GetApplicationInvoice(ctx, identity.User, id)
	if err != nil {
		if err2, ok := errors.Cause(err).(*core.Error); ok {
			return personalops.NewGetApplicationInoviceBadRequest().WithPayload(models.NewError(err2))
		}
		return personalops.NewGetApplicationInoviceInternalServerError().WithPayload(models.NewInternalServerError(err))
	}

	return personalops.NewGetApplicationInoviceOK().WithPayload(models.NewApplicationInvoice(h.Storage, invoice))
}

func (h *Handler) createApplicationPayment(params personalops.CreateApplicationPaymentParams, identity *authz.Identity) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	id := core.LotID(params.ID)

	invoice, err := h.Personal.CreateApplicationPayment(ctx, identity.User, id, params.Gateway)
	if err != nil {
		if err2, ok := errors.Cause(err).(*core.Error); ok {
			return personalops.NewCreateApplicationPaymentBadRequest().WithPayload(models.NewError(err2))
		}
		return personalops.NewCreateApplicationPaymentInternalServerError().WithPayload(models.NewInternalServerError(err))
	}

	return personalops.NewCreateApplicationPaymentCreated().WithPayload(models.NewPaymentForm(invoice))
}

func (h *Handler) getPaymentStatus(params personalops.GetPaymentStatusParams, identity *authz.Identity) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	id := core.PaymentID(params.ID)

	status, err := h.Personal.GetPaymentStatus(ctx, identity.User, id)
	if err != nil {
		if err2, ok := errors.Cause(err).(*core.Error); ok {
			return personalops.NewGetPaymentStatusBadRequest().WithPayload(models.NewError(err2))
		}
		return personalops.NewGetPaymentStatusInternalServerError().WithPayload(models.NewInternalServerError(err))
	}

	return personalops.NewGetPaymentStatusOK().WithPayload(models.NewPaymentStatus(status))
}

func (h *Handler) getOwnedLots(params personalops.GetUserLotsParams, identity *authz.Identity) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	lots, err := h.Personal.GetLots(ctx, identity.User)
	if err != nil {
		if err2, ok := errors.Cause(err).(*core.Error); ok {
			return personalops.NewGetUserLotsBadRequest().WithPayload(models.NewError(err2))
		}
		return personalops.NewGetUserLotsInternalServerError().WithPayload(models.NewInternalServerError(err))
	}

	return personalops.NewGetUserLotsOK().WithPayload(models.NewOwnedLotSlice(h.Storage, lots))

}
