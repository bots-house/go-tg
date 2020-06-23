package models

import (
	"github.com/go-openapi/swag"

	"github.com/bots-house/birzzha/api/gen/models"
	"github.com/bots-house/birzzha/pkg/tg"
)

func newResolveResultChannel(in *tg.Channel) *models.TelegramResolveResultChannel {
	if in == nil {
		return nil
	}

	return &models.TelegramResolveResultChannel{
		ID:           swag.Int64(in.ID),
		Name:         swag.String(in.Name),
		Avatar:       swag.String(in.Avatar),
		Description:  swag.String(in.Description),
		MembersCount: swag.Int64(int64(in.MembersCount)),
		Username:     swag.String(in.Username),
		DailyAverage: swag.Int64(int64(in.DailyCoverage)),
	}
}

func newResolveResultGroup(in *tg.Group) *models.TelegramResolveResultGroup {
	if in == nil {
		return nil
	}

	return &models.TelegramResolveResultGroup{
		ID:           swag.Int64(in.ID),
		Name:         swag.String(in.Name),
		Avatar:       swag.String(in.Avatar),
		Description:  swag.String(in.Description),
		MembersCount: swag.Int64(int64(in.MembersCount)),
		Username:     swag.String(in.Username),
	}
}

func NewResolveResult(in *tg.ResolveResult) *models.TelegramResolveResult {
	return &models.TelegramResolveResult{
		Channel: newResolveResultChannel(in.Channel),
		Group:   newResolveResultGroup(in.Group),
	}
}
