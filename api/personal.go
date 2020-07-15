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

	files := make([]core.LotFileID, len(payload.Files))
	for i, v := range payload.Files {
		files[i] = core.LotFileID(v)
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
		Files:         files,
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

func (h *Handler) getLotCanceledReasons(params personalops.GetLotCanceledReasonsParams) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	lcrs, err := h.Personal.GetLotCanceledReasons(ctx)
	if err != nil {
		if err2, ok := errors.Cause(err).(*core.Error); ok {
			return personalops.NewGetLotCanceledReasonsBadRequest().WithPayload(models.NewError(err2))
		}
		return personalops.NewGetLotCanceledReasonsInternalServerError().WithPayload(models.NewInternalServerError(err))
	}

	return personalops.NewGetLotCanceledReasonsOK().WithPayload(models.NewLotCanceledReasonSlice(lcrs))
}

func (h *Handler) cancelLot(params personalops.CancelLotParams, identity *authz.Identity) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	user := identity.GetUser()
	lotID := core.LotID(params.ID)
	reasonID := core.LotCanceledReasonID(params.ReasonID)

	err := h.Personal.CancelLot(ctx, user, lotID, reasonID)
	if err != nil {
		if err2, ok := errors.Cause(err).(*core.Error); ok {
			return personalops.NewCancelLotBadRequest().WithPayload(models.NewError(err2))
		}
		return personalops.NewCancelLotInternalServerError().WithPayload(models.NewInternalServerError(err))
	}

	return personalops.NewCancelLotOK()
}

func (h *Handler) uploadLotFile(params personalops.UploadLotFileParams, identity *authz.Identity) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	defer params.File.Close()
	result, err := h.Personal.UploadLotFile(
		ctx,
		params.File,
		params.HTTPRequest.MultipartForm.File["file"][0].Filename,
		params.HTTPRequest.MultipartForm.File["file"][0].Size,
		params.HTTPRequest.MultipartForm.File["file"][0].Header["Content-Type"][0],
	)
	if err != nil {
		if err2, ok := errors.Cause(err).(*core.Error); ok {
			return personalops.NewUploadLotFileBadRequest().WithPayload(models.NewError(err2))
		}
		return personalops.NewUploadLotFileInternalServerError().WithPayload(models.NewInternalServerError(err))
	}
	return personalops.NewUploadLotFileCreated().WithPayload(models.NewUploadedLotFile(h.Storage, result))
}
