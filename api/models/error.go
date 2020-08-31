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

func NewUnitPayError(err *core.Error) *models.UnitPayErrorResp {
	return &models.UnitPayErrorResp{
		Error: &models.UnitPayErrorRespError{
			Message: swag.String(err.Description),
		},
	}
}

func NewUnitPayInternalError(ctx context.Context, err error) *models.UnitPayErrorResp {
	log.Error(ctx, internalErr, "error", err)
	return &models.UnitPayErrorResp{
		Error: &models.UnitPayErrorRespError{
			Message: swag.String(err.Error()),
		},
	}
}

func NewInternalServerError(ctx context.Context, err error) *models.Error {
	log.Error(ctx, internalErr, "error", err)
	return NewError(&core.Error{
		Code:        internalErr,
		Description: err.Error(),
	})
}
