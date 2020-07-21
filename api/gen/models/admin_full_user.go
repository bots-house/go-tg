// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AdminFullUser admin full user
//
// swagger:model AdminFullUser
type AdminFullUser struct {

	// ID пользователя.
	// Required: true
	ID *int64 `json:"id"`

	// Telegram ID пользователя.
	// Required: true
	TelegramID *int64 `json:"telegram_id"`

	// Ссылка на аватарку пользователя.
	// Required: true
	Avatar *string `json:"avatar"`

	// Полное имя пользователя.
	// Required: true
	FullName *string `json:"full_name"`

	// Юзернейм пользователя.
	// Required: true
	Username *string `json:"username"`

	// Ссылка на профиль пользователя в Telegram.
	// В случае если нет username, пользователя сначала перенаправит в бота,
	// а бот уже даст специальную ссылку которая работает только в Telegram.
	//
	// Required: true
	Link *string `json:"link"`

	// Количество лотов пользователя.
	// Required: true
	Lots *int64 `json:"lots"`

	// Админ ли пользователь.
	// Required: true
	IsAdmin *bool `json:"is_admin"`

	// Метод регистрации пользователя.
	// Required: true
	// Enum: [site bot]
	JoinedFrom *string `json:"joined_from"`

	// Дата регистрации пользователя.
	// Required: true
	JoinedAt *int64 `json:"joined_at"`

	// Дата обновления данных пользователя.
	// Required: true
	UpdatedAt *int64 `json:"updated_at"`
}

// Validate validates this admin full user
func (m *AdminFullUser) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTelegramID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAvatar(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFullName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsername(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLink(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLots(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsAdmin(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateJoinedFrom(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateJoinedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AdminFullUser) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *AdminFullUser) validateTelegramID(formats strfmt.Registry) error {

	if err := validate.Required("telegram_id", "body", m.TelegramID); err != nil {
		return err
	}

	return nil
}

func (m *AdminFullUser) validateAvatar(formats strfmt.Registry) error {

	if err := validate.Required("avatar", "body", m.Avatar); err != nil {
		return err
	}

	return nil
}

func (m *AdminFullUser) validateFullName(formats strfmt.Registry) error {

	if err := validate.Required("full_name", "body", m.FullName); err != nil {
		return err
	}

	return nil
}

func (m *AdminFullUser) validateUsername(formats strfmt.Registry) error {

	if err := validate.Required("username", "body", m.Username); err != nil {
		return err
	}

	return nil
}

func (m *AdminFullUser) validateLink(formats strfmt.Registry) error {

	if err := validate.Required("link", "body", m.Link); err != nil {
		return err
	}

	return nil
}

func (m *AdminFullUser) validateLots(formats strfmt.Registry) error {

	if err := validate.Required("lots", "body", m.Lots); err != nil {
		return err
	}

	return nil
}

func (m *AdminFullUser) validateIsAdmin(formats strfmt.Registry) error {

	if err := validate.Required("is_admin", "body", m.IsAdmin); err != nil {
		return err
	}

	return nil
}

var adminFullUserTypeJoinedFromPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["site","bot"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		adminFullUserTypeJoinedFromPropEnum = append(adminFullUserTypeJoinedFromPropEnum, v)
	}
}

const (

	// AdminFullUserJoinedFromSite captures enum value "site"
	AdminFullUserJoinedFromSite string = "site"

	// AdminFullUserJoinedFromBot captures enum value "bot"
	AdminFullUserJoinedFromBot string = "bot"
)

// prop value enum
func (m *AdminFullUser) validateJoinedFromEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, adminFullUserTypeJoinedFromPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AdminFullUser) validateJoinedFrom(formats strfmt.Registry) error {

	if err := validate.Required("joined_from", "body", m.JoinedFrom); err != nil {
		return err
	}

	// value enum
	if err := m.validateJoinedFromEnum("joined_from", "body", *m.JoinedFrom); err != nil {
		return err
	}

	return nil
}

func (m *AdminFullUser) validateJoinedAt(formats strfmt.Registry) error {

	if err := validate.Required("joined_at", "body", m.JoinedAt); err != nil {
		return err
	}

	return nil
}

func (m *AdminFullUser) validateUpdatedAt(formats strfmt.Registry) error {

	if err := validate.Required("updated_at", "body", m.UpdatedAt); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AdminFullUser) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AdminFullUser) UnmarshalBinary(b []byte) error {
	var res AdminFullUser
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
