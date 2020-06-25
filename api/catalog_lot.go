package api

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"github.com/pkg/errors"

	"github.com/bots-house/birzzha/api/authz"
	catalogops "github.com/bots-house/birzzha/api/gen/restapi/operations/catalog"
	"github.com/bots-house/birzzha/api/models"
	"github.com/bots-house/birzzha/core"
	"github.com/bots-house/birzzha/service/catalog"
)

func (h *Handler) createLot(params catalogops.CreateLotParams, identity *authz.Identity) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	payload := params.Payload

	topics := make([]core.TopicID, len(payload.TopicIds))
	for i, v := range payload.TopicIds {
		topics[i] = core.TopicID(v)
	}

	in := &catalog.LotInput{
		Query:         swag.StringValue(payload.Query),
		TelegramID:    swag.Int64Value(payload.TelegramID),
		TopicIDs:      topics,
		Price:         int(swag.Int64Value(payload.Price)),
		IsBargain:     swag.BoolValue(payload.IsBargain),
		MonthlyIncome: int(swag.Int64Value(payload.MonthlyIncome)),
		Comment:       swag.StringValue(payload.Comment),
		Extra:         payload.Extra,
	}

	lot, err := h.Catalog.AddLot(ctx, identity.User, in)
	if err != nil {
		if err2, ok := errors.Cause(err).(*core.Error); ok {
			return catalogops.NewCreateLotBadRequest().WithPayload(models.NewError(err2))
		}
		return catalogops.NewCreateLotInternalServerError().WithPayload(models.NewInternalServerError(err))
	}

	return catalogops.NewCreateLotCreated().WithPayload(models.NewOwnedLot(h.Storage, lot))
}
