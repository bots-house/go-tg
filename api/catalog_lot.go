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
	"github.com/bots-house/birzzha/store"
)

func (h *Handler) createLot(params catalogops.CreateLotParams, identity *authz.Identity) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	payload := params.Payload

	topics := make([]core.TopicID, len(payload.Topics))
	for i, v := range payload.Topics {
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

func newTopicIDSlice(in []int64) []core.TopicID {
	result := make([]core.TopicID, len(in))

	for i, v := range in {
		result[i] = core.TopicID(v)
	}

	return result
}

func (h *Handler) getFilterBoundaries(params catalogops.GetFilterBoundariesParams) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	query := &core.LotFilterBoundariesQuery{Topics: newTopicIDSlice(params.Topics)}

	boundaries, err := h.Catalog.GetFilterBoundaries(ctx, query)
	if err != nil {
		return catalogops.NewCreateLotInternalServerError().WithPayload(models.NewInternalServerError(err))
	}

	return catalogops.NewGetFilterBoundariesOK().WithPayload(models.NewFilterBoundaries(boundaries))
}

func (h *Handler) getLots(params catalogops.GetLotsParams) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	query := &catalog.LotsQuery{
		Topics:             newTopicIDSlice(params.Topics),
		MembersCountFrom:   int(swag.Int64Value(params.MembersCountFrom)),
		MembersCountTo:     int(swag.Int64Value(params.MembersCountTo)),
		PriceFrom:          int(swag.Int64Value(params.PriceFrom)),
		PriceTo:            int(swag.Int64Value(params.PriceTo)),
		PricePerMemberFrom: swag.Float64Value(params.PricePerMemberFrom),
		PricePerMemberTo:   swag.Float64Value(params.PricePerMemberTo),
		DailyCoverageFrom:  int(swag.Int64Value(params.DailyCoverageFrom)),
		DailyCoverageTo:    int(swag.Int64Value(params.DailyCoverageTo)),
		PricePerViewFrom:   swag.Float64Value(params.PricePerViewFrom),
		PricePerViewTo:     swag.Float64Value(params.PricePerViewTo),
		MonthlyIncomeFrom:  int(swag.Int64Value(params.MonthlyIncomeFrom)),
		MonthlyIncomeTo:    int(swag.Int64Value(params.MonthlyIncomeTo)),
		PaybackPeriodFrom:  swag.Float64Value(params.PaybackPeriodFrom),
		PaybackPeriodTo:    swag.Float64Value(params.PaybackPeriodTo),
		//SortBy:             sortBy,
		SortByType:         0,
	}

	if params.SortBy != nil {
		sortBy, err := core.ParseLotField(swag.StringValue(params.SortBy))
		if err != nil {
			if err2, ok := errors.Cause(err).(*core.Error); ok {
				return catalogops.NewGetLotsBadRequest().WithPayload(models.NewError(err2))
			}
			return catalogops.NewGetLotsInternalServerError().WithPayload(models.NewInternalServerError(err))
		}

		query.SortBy = sortBy
	}

	if params.SortByType != nil {
		v := *params.SortByType
		switch v {
		case "asc":
			query.SortByType = store.SortTypeAsc
		case "desc":
			query.SortByType = store.SortTypeDesc
		}
	}

	lots, err := h.Catalog.GetLots(ctx, query)
	if err != nil {
		if err2, ok := errors.Cause(err).(*core.Error); ok {
			return catalogops.NewGetLotsBadRequest().WithPayload(models.NewError(err2))
		}
		return catalogops.NewGetLotsInternalServerError().WithPayload(models.NewInternalServerError(err))
	}

	return catalogops.NewGetLotsOK().WithPayload(models.NewItemLotSlice(h.Storage, lots))
}