package models

import (
	"github.com/bots-house/birzzha/api/gen/models"
	"github.com/bots-house/birzzha/core"
	"github.com/go-openapi/swag"
)

func NewError(err *core.Error) *models.Error {
	return &models.Error{
		Code:        swag.String(err.Code),
		Description: swag.String(err.Description),
	}
}

func NewInternalServerError(err error) *models.Error {
	return NewError(&core.Error{
		Code:        "internal_server_error",
		Description: err.Error(),
	})
}
