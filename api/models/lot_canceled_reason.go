package models

import (
	"github.com/bots-house/birzzha/api/gen/models"
	"github.com/bots-house/birzzha/core"
	"github.com/go-openapi/swag"
)

func NewLotCanceledReason(lcr *core.LotCanceledReason) *models.LotCanceledReason {
	return &models.LotCanceledReason{
		ID:       swag.Int64(int64(lcr.ID)),
		Why:      swag.String(lcr.Why),
		IsPublic: swag.Bool(lcr.IsPublic),
	}
}

func NewLotCanceledReasonSlice(lcrs []*core.LotCanceledReason) []*models.LotCanceledReason {
	res := make([]*models.LotCanceledReason, len(lcrs))

	for i, lcr := range lcrs {
		res[i] = NewLotCanceledReason(lcr)
	}

	return res
}
