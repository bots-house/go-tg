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

// Post post
//
// swagger:model Post
type Post struct {

	// ID поста.
	// Required: true
	ID *int64 `json:"id"`

	// ID лота.
	// Required: true
	LotID *int64 `json:"lot_id"`

	// Название поста.
	// Required: true
	Title *string `json:"title"`

	// Текст поста.
	// Required: true
	Text *string `json:"text"`

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

// Validate validates this post
func (m *Post) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLotID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTitle(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateText(formats); err != nil {
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

func (m *Post) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *Post) validateLotID(formats strfmt.Registry) error {

	if err := validate.Required("lot_id", "body", m.LotID); err != nil {
		return err
	}

	return nil
}

func (m *Post) validateTitle(formats strfmt.Registry) error {

	if err := validate.Required("title", "body", m.Title); err != nil {
		return err
	}

	return nil
}

func (m *Post) validateText(formats strfmt.Registry) error {

	if err := validate.Required("text", "body", m.Text); err != nil {
		return err
	}

	return nil
}

func (m *Post) validateDisableWebPagePreview(formats strfmt.Registry) error {

	if err := validate.Required("disable_web_page_preview", "body", m.DisableWebPagePreview); err != nil {
		return err
	}

	return nil
}

func (m *Post) validateScheduledAt(formats strfmt.Registry) error {

	if err := validate.Required("scheduled_at", "body", m.ScheduledAt); err != nil {
		return err
	}

	return nil
}

func (m *Post) validatePublishedAt(formats strfmt.Registry) error {

	if err := validate.Required("published_at", "body", m.PublishedAt); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Post) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Post) UnmarshalBinary(b []byte) error {
	var res Post
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
