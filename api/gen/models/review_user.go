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

// ReviewUser Пользователь который оставил отзыв.
//
// swagger:model ReviewUser
type ReviewUser struct {

	// Имя пользователя.
	// Required: true
	FirstName *string `json:"first_name"`

	// Фамилия пользователя.
	// Required: true
	LastName *string `json:"last_name"`

	// Юзернейм пользователя.
	// Required: true
	Username *string `json:"username"`

	// Ссылка на аватарку пользователя.
	// Required: true
	Avatar *string `json:"avatar"`
}

// Validate validates this review user
func (m *ReviewUser) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFirstName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsername(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAvatar(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReviewUser) validateFirstName(formats strfmt.Registry) error {

	if err := validate.Required("first_name", "body", m.FirstName); err != nil {
		return err
	}

	return nil
}

func (m *ReviewUser) validateLastName(formats strfmt.Registry) error {

	if err := validate.Required("last_name", "body", m.LastName); err != nil {
		return err
	}

	return nil
}

func (m *ReviewUser) validateUsername(formats strfmt.Registry) error {

	if err := validate.Required("username", "body", m.Username); err != nil {
		return err
	}

	return nil
}

func (m *ReviewUser) validateAvatar(formats strfmt.Registry) error {

	if err := validate.Required("avatar", "body", m.Avatar); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReviewUser) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReviewUser) UnmarshalBinary(b []byte) error {
	var res ReviewUser
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}