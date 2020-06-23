package models

import (
	"time"

	"github.com/bots-house/birzzha/api/gen/models"
	"github.com/bots-house/birzzha/core"
	"github.com/bots-house/birzzha/pkg/storage"
	"github.com/go-openapi/swag"
	"github.com/volatiletech/null/v8"
)

func nullStringToString(v null.String) *string {
	if v.Valid {
		return &v.String
	}
	return nil
}

func timeToUnix(v time.Time) int64 {
	return v.Unix()
}

func NewUser(s storage.Storage, user *core.User) *models.User {
	result := &models.User{
		ID: swag.Int64(int64(user.ID)),
		Telegram: &models.UserTelegram{
			ID:       swag.Int64(int64(user.Telegram.ID)),
			Username: nullStringToString(user.Telegram.Username),
		},
		FirstName: swag.String(user.FirstName),
		LastName:  nullStringToString(user.LastName),
		IsAdmin:   swag.Bool(user.IsAdmin),
		JoinedAt:  swag.Int64(timeToUnix(user.JoinedAt)),
	}

	if user.Avatar.Valid {
		result.Avatar = swag.String(s.PublicURL(user.Avatar.String))
	}

	return result
}
