package models

import "github.com/bots-house/birzzha/api/gen/models"

func NewDailyCoverage(views int64) *models.DailyCoverage {
	return &models.DailyCoverage{Views: &views}
}
