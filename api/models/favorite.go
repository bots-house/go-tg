package models

import (
	"github.com/bots-house/birzzha/api/gen/models"
	"github.com/go-openapi/swag"
)

func NewLotFavoriteStatus(in bool) *models.LotFavoriteStatus {
	return &models.LotFavoriteStatus{
		InFavorites: swag.Bool(in),
	}
}
