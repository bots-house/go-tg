package models

import (
	"time"

	"github.com/bots-house/birzzha/api/gen/models"
	"github.com/bots-house/birzzha/core"
	"github.com/bots-house/birzzha/pkg/storage"
	"github.com/bots-house/birzzha/service/admin"
	"github.com/go-openapi/swag"
	"github.com/volatiletech/null/v8"
)

func nullStringToString(v null.String) *string {
	if v.Valid {
		return &v.String
	}
	return nil
}

func nullBoolToBool(v null.Bool) *bool {
	if v.Valid {
		return &v.Bool
	}
	return nil
}

func nullIntToInt64(v null.Int) *int64 {
	if v.Valid {
		x := int64(v.Int)
		return &x
	}
	return nil
}

func nullFloat64ToFloat64(v null.Float64) *float64 {
	if v.Valid {
		return &v.Float64
	}
	return nil
}

func timeToUnix(v time.Time) *int64 {
	x := v.Unix()
	return &x
}

func nullTimeToUnix(v null.Time) *int64 {
	if v.Valid {
		x := v.Time.Unix()
		return &x
	}
	return nil
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
		JoinedAt:  timeToUnix(user.JoinedAt),
	}

	if user.Avatar.Valid {
		result.Avatar = swag.String(s.PublicURL(user.Avatar.String))
	}

	return result
}

func NewAdminFullUser(s storage.Storage, user *admin.FullUser) *models.AdminFullUser {
	fullUser := &models.AdminFullUser{
		ID:         swag.Int64(int64(user.ID)),
		TelegramID: swag.Int64(int64(user.Telegram.ID)),
		FullName:   swag.String(user.Name()),
		Lots:       swag.Int64(int64(user.Lots)),
		IsAdmin:    swag.Bool(user.IsAdmin),
		Link:       swag.String(user.Link),
		JoinedFrom: swag.String(user.JoinedFrom.String()),
		JoinedAt:   timeToUnix(user.JoinedAt),
		Username:   nullStringToString(user.Telegram.Username),
		UpdatedAt:  nullTimeToUnix(user.UpdatedAt),
	}

	if user.Avatar.Valid {
		fullUser.Avatar = swag.String(s.PublicURL(user.Avatar.String))
	}
	return fullUser
}

func newAdminFullUserSlice(s storage.Storage, users []*admin.FullUser) []*models.AdminFullUser {
	result := make([]*models.AdminFullUser, len(users))
	for i, user := range users {
		result[i] = NewAdminFullUser(s, user)
	}
	return result
}

func NewAdminFullUserList(s storage.Storage, in *admin.FullUserList) *models.AdminFullUserList {
	return &models.AdminFullUserList{
		Total: swag.Int64(int64(in.Total)),
		Items: newAdminFullUserSlice(s, in.Items),
	}
}
