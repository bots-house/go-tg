package api

import (
	catalogops "github.com/bots-house/birzzha/api/gen/restapi/operations/catalog"
	"github.com/bots-house/birzzha/api/models"
	"github.com/go-openapi/runtime/middleware"
)

func (h *Handler) getTopics(params catalogops.GetTopicsParams) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	topics, err := h.Catalog.GetTopics(ctx)
	if err != nil {
		return catalogops.NewGetTopicsInternalServerError()
	}

	return catalogops.NewGetTopicsOK().WithPayload(models.NewTopicSlice(topics))
}
