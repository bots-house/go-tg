package models

import "github.com/bots-house/birzzha/api/gen/models"

func NewUnitPaySuccessResp(msg string) *models.UnitPaySuccessResp {
	return &models.UnitPaySuccessResp{Result: &models.UnitPaySuccessRespResult{
		Message: &msg,
	}}
}
