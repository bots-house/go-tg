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

// LotMetrics lot metrics
//
// swagger:model LotMetrics
type LotMetrics struct {

	// Кол-во участников в канале
	// Required: true
	MembersCount *int64 `json:"members_count"`

	// Дневной охват
	// Required: true
	DailyCoverage *int64 `json:"daily_coverage"`

	// Прибыль в месяц
	// Required: true
	MonthlyIncome *int64 `json:"monthly_income"`

	// Цена за подписчика
	// Required: true
	PricePerMember *float64 `json:"price_per_member"`

	// Цена за просмотр
	// Required: true
	PricePerView *float64 `json:"price_per_view"`

	// Окупаемость
	// Required: true
	PaybackPeriod *float64 `json:"payback_period"`
}

// Validate validates this lot metrics
func (m *LotMetrics) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMembersCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDailyCoverage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMonthlyIncome(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePricePerMember(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePricePerView(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePaybackPeriod(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LotMetrics) validateMembersCount(formats strfmt.Registry) error {

	if err := validate.Required("members_count", "body", m.MembersCount); err != nil {
		return err
	}

	return nil
}

func (m *LotMetrics) validateDailyCoverage(formats strfmt.Registry) error {

	if err := validate.Required("daily_coverage", "body", m.DailyCoverage); err != nil {
		return err
	}

	return nil
}

func (m *LotMetrics) validateMonthlyIncome(formats strfmt.Registry) error {

	if err := validate.Required("monthly_income", "body", m.MonthlyIncome); err != nil {
		return err
	}

	return nil
}

func (m *LotMetrics) validatePricePerMember(formats strfmt.Registry) error {

	if err := validate.Required("price_per_member", "body", m.PricePerMember); err != nil {
		return err
	}

	return nil
}

func (m *LotMetrics) validatePricePerView(formats strfmt.Registry) error {

	if err := validate.Required("price_per_view", "body", m.PricePerView); err != nil {
		return err
	}

	return nil
}

func (m *LotMetrics) validatePaybackPeriod(formats strfmt.Registry) error {

	if err := validate.Required("payback_period", "body", m.PaybackPeriod); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LotMetrics) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LotMetrics) UnmarshalBinary(b []byte) error {
	var res LotMetrics
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
