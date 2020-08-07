package models

import (
	"github.com/bots-house/birzzha/api/gen/models"
	"github.com/bots-house/birzzha/pkg/health"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

func newHealthCheck(check health.Check) *models.HealthCheck {
	var errText string

	if check.Err != nil {
		errText = check.Err.Error()
	}

	took := strfmt.Duration(check.Took)

	return &models.HealthCheck{
		Ok:   swag.Bool(check.Ok()),
		Err:  errText,
		Took: &took,
	}
}

func NewHealth(health *health.Health) *models.Health {
	return &models.Health{
		Ok:       swag.Bool(health.Ok()),
		Postgres: newHealthCheck(health.Postgres),
		Redis:    newHealthCheck(health.Redis),
	}
}
