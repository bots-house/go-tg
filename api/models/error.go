package models

import (
	"context"

	"github.com/bots-house/birzzha/api/gen/models"
	"github.com/bots-house/birzzha/core"
	"github.com/bots-house/birzzha/pkg/log"
	"github.com/go-openapi/swag"
)

const internalErr = "internal_server_error"

func NewError(err *core.Error) *models.Error {
	return &models.Error{
		Code:        swag.String(err.Code),
		Description: swag.String(err.Description),
	}
}

func NewInternalServerError(ctx context.Context, err error) *models.Error {
	log.Error(ctx, internalErr, "error", err)
	return NewError(&core.Error{
		Code:        internalErr,
		Description: err.Error(),
	})
}
