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

// PostItem post item
//
// swagger:model PostItem
type PostItem struct {

	// ID поста.
	// Required: true
	ID *int64 `json:"id"`

	// lot
	// Required: true
	Lot *PostLot `json:"lot"`

	// Текст поста.
	// Required: true
	Text *string `json:"text"`

	// Название поста.
	// Required: true
	Title *string `json:"title"`

	// Отключить или выключить web page preview.
	// Required: true
	DisableWebPagePreview *bool `json:"disable_web_page_preview"`

	// Время планирования поста.
	// Required: true
	ScheduledAt *int64 `json:"scheduled_at"`

	// Время публикации поста.
	// Required: true
	PublishedAt *int64 `json:"published_at"`
}

// Validate validates this post item
func (m *PostItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLot(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateText(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTitle(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDisableWebPagePreview(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScheduledAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePublishedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostItem) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *PostItem) validateLot(formats strfmt.Registry) error {

	if err := validate.Required("lot", "body", m.Lot); err != nil {
		return err
	}

	if m.Lot != nil {
		if err := m.Lot.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lot")
			}
			return err
		}
	}

	return nil
}

func (m *PostItem) validateText(formats strfmt.Registry) error {

	if err := validate.Required("text", "body", m.Text); err != nil {
		return err
	}

	return nil
}

func (m *PostItem) validateTitle(formats strfmt.Registry) error {

	if err := validate.Required("title", "body", m.Title); err != nil {
		return err
	}

	return nil
}

func (m *PostItem) validateDisableWebPagePreview(formats strfmt.Registry) error {

	if err := validate.Required("disable_web_page_preview", "body", m.DisableWebPagePreview); err != nil {
		return err
	}

	return nil
}

func (m *PostItem) validateScheduledAt(formats strfmt.Registry) error {

	if err := validate.Required("scheduled_at", "body", m.ScheduledAt); err != nil {
		return err
	}

	return nil
}

func (m *PostItem) validatePublishedAt(formats strfmt.Registry) error {

	if err := validate.Required("published_at", "body", m.PublishedAt); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PostItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostItem) UnmarshalBinary(b []byte) error {
	var res PostItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}