package api

import (
	healthops "github.com/bots-house/birzzha/api/gen/restapi/operations/health"
	"github.com/bots-house/birzzha/api/models"
	"github.com/go-openapi/runtime/middleware"
)

func (h *Handler) getHealth(params healthops.GetHealthParams) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	health := h.Health.Check(ctx)

	payload := models.NewHealth(health)

	if health.Ok() {
		return healthops.NewGetHealthOK().WithPayload(payload)
	}

	return healthops.NewGetHealthServiceUnavailable().WithPayload(payload)
}
