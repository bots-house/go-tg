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

// Landing landing
//
// swagger:model Landing
type Landing struct {

	// stats
	// Required: true
	Stats *LandingStats `json:"stats"`

	// channel
	// Required: true
	Channel *LandingChannel `json:"channel"`

	// reviews
	// Required: true
	Reviews *ReviewList `json:"reviews"`
}

// Validate validates this landing
func (m *Landing) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStats(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChannel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReviews(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Landing) validateStats(formats strfmt.Registry) error {

	if err := validate.Required("stats", "body", m.Stats); err != nil {
		return err
	}

	if m.Stats != nil {
		if err := m.Stats.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("stats")
			}
			return err
		}
	}

	return nil
}

func (m *Landing) validateChannel(formats strfmt.Registry) error {

	if err := validate.Required("channel", "body", m.Channel); err != nil {
		return err
	}

	if m.Channel != nil {
		if err := m.Channel.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("channel")
			}
			return err
		}
	}

	return nil
}

func (m *Landing) validateReviews(formats strfmt.Registry) error {

	if err := validate.Required("reviews", "body", m.Reviews); err != nil {
		return err
	}

	if m.Reviews != nil {
		if err := m.Reviews.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("reviews")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Landing) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Landing) UnmarshalBinary(b []byte) error {
	var res Landing
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}