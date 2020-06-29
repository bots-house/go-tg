package api

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"github.com/pkg/errors"

	catalogops "github.com/bots-house/birzzha/api/gen/restapi/operations/catalog"
	"github.com/bots-house/birzzha/api/models"
	"github.com/bots-house/birzzha/core"
	"github.com/bots-house/birzzha/service/catalog"
	"github.com/bots-house/birzzha/store"
)

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
		return catalogops.NewGetFilterBoundariesInternalServerError().WithPayload(models.NewInternalServerError(err))
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
		SortByType: 0,
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

func (h *Handler) resolveTelegram(params catalogops.ResolveTelegramParams) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	result, err := h.Catalog.ResolveTelegram(ctx, params.Q)
	if err != nil {
		if err2, ok := errors.Cause(err).(*core.Error); ok {
			return catalogops.NewResolveTelegramBadRequest().WithPayload(models.NewError(err2))
		}
		return catalogops.NewResolveTelegramBadRequest().WithPayload(models.NewInternalServerError(err))
	}

	return catalogops.NewResolveTelegramOK().WithPayload(models.NewResolveResult(result))
}

func (h *Handler) getTopics(params catalogops.GetTopicsParams) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	topics, err := h.Catalog.GetTopics(ctx)
	if err != nil {
		return catalogops.NewGetTopicsInternalServerError()
	}

	return catalogops.NewGetTopicsOK().WithPayload(models.NewTopicSlice(topics))
}
