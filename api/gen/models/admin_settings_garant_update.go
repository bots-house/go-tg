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

// AdminSettingsGarantUpdate admin settings garant update
//
// swagger:model AdminSettingsGarantUpdate
type AdminSettingsGarantUpdate struct {

	// имя гаранта
	// Required: true
	Name *string `json:"name"`

	// username гаранта
	// Required: true
	Username *string `json:"username"`

	// username канала с отызывами
	// Required: true
	ReviewsChannel *string `json:"reviews_channel"`

	// процент от сделки
	// Required: true
	PercentageDeal *float64 `json:"percentage_deal"`

	// процент от сделки (на период скидки)
	PercentageDealOfDiscountPeriod float64 `json:"percentage_deal_of_discount_period,omitempty"`

	// ссылка на аватарку гаранта (может быть null)
	AvatarURL string `json:"avatar_url,omitempty"`
}

// Validate validates this admin settings garant update
func (m *AdminSettingsGarantUpdate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsername(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReviewsChannel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePercentageDeal(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AdminSettingsGarantUpdate) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *AdminSettingsGarantUpdate) validateUsername(formats strfmt.Registry) error {

	if err := validate.Required("username", "body", m.Username); err != nil {
		return err
	}

	return nil
}

func (m *AdminSettingsGarantUpdate) validateReviewsChannel(formats strfmt.Registry) error {

	if err := validate.Required("reviews_channel", "body", m.ReviewsChannel); err != nil {
		return err
	}

	return nil
}

func (m *AdminSettingsGarantUpdate) validatePercentageDeal(formats strfmt.Registry) error {

	if err := validate.Required("percentage_deal", "body", m.PercentageDeal); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AdminSettingsGarantUpdate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AdminSettingsGarantUpdate) UnmarshalBinary(b []byte) error {
	var res AdminSettingsGarantUpdate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
