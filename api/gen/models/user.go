// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// User Объект пользователя
//
// swagger:model User
type User struct {

	// Уникальный ID пользователя в Birzzha.
	// Required: true
	ID *int64 `json:"id"`

	// Имя пользователя в Telegram
	// Required: true
	FirstName *string `json:"first_name"`

	// Фамилия пользователя в Telegram (может быть `null`)
	// Required: true
	LastName *string `json:"last_name"`

	// Path to avatar
	// Required: true
	Avatar *string `json:"avatar"`

	// True, если пользователь админ Birzzha.
	// Required: true
	IsAdmin *bool `json:"is_admin"`

	// Дата и время регистрации на бирже, в Unix-time.
	// Required: true
	JoinedAt *int64 `json:"joined_at"`

	// telegram
	// Required: true
	Telegram *UserTelegram `json:"telegram"`
}

// Validate validates this user
func (m *User) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFirstName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAvatar(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsAdmin(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateJoinedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTelegram(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *User) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *User) validateFirstName(formats strfmt.Registry) error {

	if err := validate.Required("first_name", "body", m.FirstName); err != nil {
		return err
	}

	return nil
}

func (m *User) validateLastName(formats strfmt.Registry) error {

	if err := validate.Required("last_name", "body", m.LastName); err != nil {
		return err
	}

	return nil
}

func (m *User) validateAvatar(formats strfmt.Registry) error {

	if err := validate.Required("avatar", "body", m.Avatar); err != nil {
		return err
	}

	return nil
}

func (m *User) validateIsAdmin(formats strfmt.Registry) error {

	if err := validate.Required("is_admin", "body", m.IsAdmin); err != nil {
		return err
	}

	return nil
}

func (m *User) validateJoinedAt(formats strfmt.Registry) error {

	if err := validate.Required("joined_at", "body", m.JoinedAt); err != nil {
		return err
	}

	return nil
}

func (m *User) validateTelegram(formats strfmt.Registry) error {

	if err := validate.Required("telegram", "body", m.Telegram); err != nil {
		return err
	}

	if m.Telegram != nil {
		if err := m.Telegram.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("telegram")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *User) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *User) UnmarshalBinary(b []byte) error {
	var res User
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// UserTelegram Информация о пользователе из Telegram
//
// swagger:model UserTelegram
type UserTelegram struct {

	// ID пользователя в Telegram
	// Required: true
	ID *int64 `json:"id"`

	// Username пользователя в Telegram
	// Required: true
	Username *string `json:"username"`
}

// Validate validates this user telegram
func (m *UserTelegram) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsername(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserTelegram) validateID(formats strfmt.Registry) error {

	if err := validate.Required("telegram"+"."+"id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *UserTelegram) validateUsername(formats strfmt.Registry) error {

	if err := validate.Required("telegram"+"."+"username", "body", m.Username); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UserTelegram) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserTelegram) UnmarshalBinary(b []byte) error {
	var res UserTelegram
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
