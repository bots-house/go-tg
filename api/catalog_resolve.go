package api

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/pkg/errors"

	catalogops "github.com/bots-house/birzzha/api/gen/restapi/operations/catalog"
	"github.com/bots-house/birzzha/api/models"
	"github.com/bots-house/birzzha/core"
)

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
