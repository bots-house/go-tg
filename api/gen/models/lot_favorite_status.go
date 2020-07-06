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

// LotFavoriteStatus Ответ после смены статуса избранности лота.
//
// swagger:model LotFavoriteStatus
type LotFavoriteStatus struct {

	// Статус избранности лота.
	// Required: true
	InFavorites *bool `json:"in_favorites"`
}

// Validate validates this lot favorite status
func (m *LotFavoriteStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInFavorites(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LotFavoriteStatus) validateInFavorites(formats strfmt.Registry) error {

	if err := validate.Required("in_favorites", "body", m.InFavorites); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LotFavoriteStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LotFavoriteStatus) UnmarshalBinary(b []byte) error {
	var res LotFavoriteStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}