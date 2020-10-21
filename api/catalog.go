package api

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"github.com/pkg/errors"
	"github.com/tomasen/realip"

	"github.com/bots-house/birzzha/api/authz"
	catalogops "github.com/bots-house/birzzha/api/gen/restapi/operations/catalog"
	"github.com/bots-house/birzzha/api/models"
	"github.com/bots-house/birzzha/core"
	"github.com/bots-house/birzzha/pkg/log"
	"github.com/bots-house/birzzha/service/catalog"
	"github.com/bots-house/birzzha/service/views"

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
		return catalogops.NewGetFilterBoundariesInternalServerError().WithPayload(models.NewInternalServerError(ctx, err))
	}

	return catalogops.NewGetFilterBoundariesOK().WithPayload(models.NewFilterBoundaries(boundaries))
}

func (h *Handler) getLots(params catalogops.GetLotsParams, identity *authz.Identity) middleware.Responder {
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
		Offset:             int(swag.Int64Value(params.Offset)),
		Limit:              int(swag.Int64Value(params.Limit)),
		SortByType:         0,
	}

	if params.SortBy != nil {
		sortBy, err := core.ParseLotField(swag.StringValue(params.SortBy))
		if err != nil {
			if err2, ok := errors.Cause(err).(*core.Error); ok {
				return catalogops.NewGetLotsBadRequest().WithPayload(models.NewError(err2))
			}
			return catalogops.NewGetLotsInternalServerError().WithPayload(models.NewInternalServerError(ctx, err))
		}

		query.SortBy = sortBy
	}

	if params.SortByType != nil {
		v := *params.SortByType
		switch v {
		case ascQueryParam:
			query.SortByType = store.SortTypeAsc
		case descQueryParam:
			query.SortByType = store.SortTypeDesc
		}
	}

	lots, err := h.Catalog.GetLots(ctx, identity.GetUser(), query)
	if err != nil {
		if err2, ok := errors.Cause(err).(*core.Error); ok {
			return catalogops.NewGetLotsBadRequest().WithPayload(models.NewError(err2))
		}
		return catalogops.NewGetLotsInternalServerError().WithPayload(models.NewInternalServerError(ctx, err))
	}

	return catalogops.NewGetLotsOK().WithPayload(models.NewLotList(h.Storage, lots))
}

func (h *Handler) getDailyCoverage(params catalogops.GetDailyCoverageParams) middleware.Responder {
	count, err := h.Catalog.GetDailyCoverage(params.ChannelID)
	if err != nil {
		if err2, ok := errors.Cause(err).(*core.Error); ok {
			return catalogops.NewGetDailyCoverageBadRequest().WithPayload(models.NewError(err2))
		}
		return catalogops.NewGetDailyCoverageInternalServerError().WithPayload(models.NewInternalServerError(params.HTTPRequest.Context(), err))
	}

	return catalogops.NewGetDailyCoverageOK().WithPayload(models.NewDailyCoverage(int64(count)))
}

func (h *Handler) resolveTelegram(params catalogops.ResolveTelegramParams) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	result, err := h.Catalog.ResolveTelegram(ctx, params.Q)
	if err != nil {
		if err2, ok := errors.Cause(err).(*core.Error); ok {
			return catalogops.NewResolveTelegramBadRequest().WithPayload(models.NewError(err2))
		}
		return catalogops.NewResolveTelegramBadRequest().WithPayload(models.NewInternalServerError(ctx, err))
	}

	return catalogops.NewResolveTelegramOK().WithPayload(models.NewResolveResult(result))
}

func (h *Handler) getTopics(params catalogops.GetTopicsParams) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	topics, err := h.Catalog.GetTopics(ctx)
	if err != nil {
		return catalogops.NewGetTopicsInternalServerError()
	}

	return catalogops.NewGetTopicsOK().WithPayload(models.NewTopicItemSlice(topics))
}

func (h *Handler) getAnonymousID(r *http.Request) string {
	ip := realip.FromRequest(r)

	hsh := sha256.New()
	_, _ = hsh.Write([]byte(ip))

	return hex.EncodeToString(hsh.Sum(nil))
}

func (h *Handler) registerLotView(ctx context.Context, lot core.LotID, r *http.Request, identity *authz.Identity) {
	var view *views.SiteView

	if identity.IsAnonymous() {
		view = views.NewAnonymousView(lot, h.getAnonymousID(r))
	} else {
		view = views.NewAuthorizedView(lot, identity.User.ID)
	}

	if err := h.Views.RegisterSiteView(context.Background(), view); err != nil {
		log.Error(ctx, "register view", "lot_id", lot, "is_anonymous", identity.IsAnonymous(), "err", err)
	}
}

func (h *Handler) getLot(params catalogops.GetLotParams, identity *authz.Identity) middleware.Responder {
	ctx := params.HTTPRequest.Context()
	id := core.LotID(int(params.ID))

	result, err := h.Catalog.GetLot(ctx, identity.GetUser(), id)
	if err != nil {
		if err2, ok := errors.Cause(err).(*core.Error); ok {
			return catalogops.NewGetLotBadRequest().WithPayload(models.NewError(err2))
		}
		return catalogops.NewGetLotInternalServerError().WithPayload(models.NewInternalServerError(ctx, err))
	}

	go h.registerLotView(ctx, id, params.HTTPRequest, identity)

	return catalogops.NewGetLotOK().WithPayload(models.NewFullLot(h.Storage, result))
}

func (h *Handler) getSimilarLots(params catalogops.GetSimilarLotsParams) middleware.Responder {
	ctx := params.HTTPRequest.Context()
	id := core.LotID(int(params.ID))
	limit := int(swag.Int64Value(params.Limit))
	offset := int(swag.Int64Value(params.Offset))

	result, err := h.Catalog.SimilarLots(ctx, id, limit, offset)
	if err != nil {
		if err2, ok := errors.Cause(err).(*core.Error); ok {
			return catalogops.NewGetSimilarLotsBadRequest().WithPayload(models.NewError(err2))
		}
		return catalogops.NewGetSimilarLotsInternalServerError().WithPayload(models.NewInternalServerError(ctx, err))
	}
	return catalogops.NewGetSimilarLotsOK().WithPayload(models.NewLotList(h.Storage, result))
}

func (h *Handler) toggleLotFavorite(params catalogops.ToggleLotFavoriteParams, identity *authz.Identity) middleware.Responder {
	ctx := params.HTTPRequest.Context()
	id := core.LotID(int(params.ID))

	result, err := h.Catalog.ToggleLotFavorite(ctx, identity.User, id)
	if err != nil {
		if err2, ok := errors.Cause(err).(*core.Error); ok {
			return catalogops.NewToggleLotFavoriteBadRequest().WithPayload(models.NewError(err2))
		}
		return catalogops.NewToggleLotFavoriteInternalServerError().WithPayload(models.NewInternalServerError(ctx, err))
	}
	return catalogops.NewToggleLotFavoriteOK().WithPayload(models.NewLotFavoriteStatus(result))
}
