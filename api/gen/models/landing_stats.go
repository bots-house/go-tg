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

// LandingStats landing stats
//
// swagger:model LandingStats
type LandingStats struct {

	// Кол-во уникальных посетителей сайта.
	// Required: true
	UniqueVisitorsPerMonth *int64 `json:"unique_visitors_per_month"`

	// Cредний охват одного обьявления на сайте
	// Required: true
	AvgLotSiteReach *int64 `json:"avg_lot_site_reach"`

	// Cредний охват одного обьявления в телеграм канале
	// Required: true
	AvgLotChannelReach *int64 `json:"avg_lot_channel_reach"`
}

// Validate validates this landing stats
func (m *LandingStats) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateUniqueVisitorsPerMonth(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAvgLotSiteReach(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAvgLotChannelReach(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LandingStats) validateUniqueVisitorsPerMonth(formats strfmt.Registry) error {

	if err := validate.Required("unique_visitors_per_month", "body", m.UniqueVisitorsPerMonth); err != nil {
		return err
	}

	return nil
}

func (m *LandingStats) validateAvgLotSiteReach(formats strfmt.Registry) error {

	if err := validate.Required("avg_lot_site_reach", "body", m.AvgLotSiteReach); err != nil {
		return err
	}

	return nil
}

func (m *LandingStats) validateAvgLotChannelReach(formats strfmt.Registry) error {

	if err := validate.Required("avg_lot_channel_reach", "body", m.AvgLotChannelReach); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LandingStats) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LandingStats) UnmarshalBinary(b []byte) error {
	var res LandingStats
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
