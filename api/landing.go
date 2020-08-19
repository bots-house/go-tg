package api

import (
	landingops "github.com/bots-house/birzzha/api/gen/restapi/operations/landing"
	"github.com/bots-house/birzzha/api/models"
	"github.com/bots-house/birzzha/core"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"github.com/pkg/errors"
)

func (h *Handler) getReviews(params landingops.GetReviewsParams) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	result, err := h.Landing.GetReviews(ctx, int(swag.Int64Value(params.Offset)), int(swag.Int64Value(params.Limit)))
	if err != nil {
		if err2, ok := errors.Cause(err).(*core.Error); ok {
			return landingops.NewGetReviewsBadRequest().WithPayload(models.NewError(err2))
		}
		return landingops.NewGetReviewsInternalServerError().WithPayload(models.NewInternalServerError(ctx, err))
	}
	return landingops.NewGetReviewsOK().WithPayload(models.NewReviewList(h.Storage, result))
}

func (h *Handler) getLanding(params landingops.GetLandingParams) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	result, err := h.Landing.GetLanding(ctx)
	if err != nil {
		if err2, ok := errors.Cause(err).(*core.Error); ok {
			return landingops.NewGetLandingBadRequest().WithPayload(models.NewError(err2))
		}
		return landingops.NewGetLandingInternalServerError().WithPayload(models.NewInternalServerError(ctx, err))
	}
	return landingops.NewGetLandingOK().WithPayload(models.NewLanding(h.Storage, result))
}
