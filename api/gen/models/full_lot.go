// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// FullLot Детали лота.
//
// swagger:model FullLot
type FullLot struct {

	// ID лота
	// Required: true
	ID *int64 `json:"id"`

	// Название лота (канала) в Telegram
	// Required: true
	Name *string `json:"name"`

	// Аватарка лота
	// Required: true
	Avatar *string `json:"avatar"`

	// @username канала (может быть null)
	// Required: true
	Username *string `json:"username"`

	// Ссылка для вступления (как приватная так и публичная)
	// Required: true
	Link *string `json:"link"`

	// Описание лота.
	// Required: true
	Bio *string `json:"bio"`

	// price
	// Required: true
	Price *LotPrice `json:"price"`

	// Комментарий к лоту
	// Required: true
	Comment *string `json:"comment"`

	// metrics
	// Required: true
	Metrics *LotMetrics `json:"metrics"`

	// topics
	// Required: true
	Topics []int64 `json:"topics"`

	// True, если лот в избранном
	// Required: true
	InFavorites *bool `json:"in_favorites"`

	// Дата создания
	// Required: true
	CreatedAt *int64 `json:"created_at"`

	// user
	// Required: true
	User *LotOwner `json:"user"`

	// extra
	// Required: true
	Extra []*LotExtraResource `json:"extra"`

	// Количество просмотров.
	// Required: true
	Views *int64 `json:"views"`

	// Ссылка на tgstat.ru.
	// Required: true
	TgstatLink *string `json:"tgstat_link"`

	// Ссылка на telemetr.me.
	// Required: true
	TelemetrLink *string `json:"telemetr_link"`

	// files
	// Required: true
	Files []*OwnedLotUploadedFile `json:"files"`
}

// Validate validates this full lot
func (m *FullLot) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAvatar(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsername(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLink(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBio(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateComment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetrics(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTopics(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInFavorites(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUser(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExtra(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateViews(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTgstatLink(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTelemetrLink(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFiles(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FullLot) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *FullLot) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *FullLot) validateAvatar(formats strfmt.Registry) error {

	if err := validate.Required("avatar", "body", m.Avatar); err != nil {
		return err
	}

	return nil
}

func (m *FullLot) validateUsername(formats strfmt.Registry) error {

	if err := validate.Required("username", "body", m.Username); err != nil {
		return err
	}

	return nil
}

func (m *FullLot) validateLink(formats strfmt.Registry) error {

	if err := validate.Required("link", "body", m.Link); err != nil {
		return err
	}

	return nil
}

func (m *FullLot) validateBio(formats strfmt.Registry) error {

	if err := validate.Required("bio", "body", m.Bio); err != nil {
		return err
	}

	return nil
}

func (m *FullLot) validatePrice(formats strfmt.Registry) error {

	if err := validate.Required("price", "body", m.Price); err != nil {
		return err
	}

	if m.Price != nil {
		if err := m.Price.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("price")
			}
			return err
		}
	}

	return nil
}

func (m *FullLot) validateComment(formats strfmt.Registry) error {

	if err := validate.Required("comment", "body", m.Comment); err != nil {
		return err
	}

	return nil
}

func (m *FullLot) validateMetrics(formats strfmt.Registry) error {

	if err := validate.Required("metrics", "body", m.Metrics); err != nil {
		return err
	}

	if m.Metrics != nil {
		if err := m.Metrics.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metrics")
			}
			return err
		}
	}

	return nil
}

func (m *FullLot) validateTopics(formats strfmt.Registry) error {

	if err := validate.Required("topics", "body", m.Topics); err != nil {
		return err
	}

	for i := 0; i < len(m.Topics); i++ {

		if err := validate.MinLength("topics"+"."+strconv.Itoa(i), "body", string(m.Topics[i]), 1); err != nil {
			return err
		}

		if err := validate.MaxLength("topics"+"."+strconv.Itoa(i), "body", string(m.Topics[i]), 3); err != nil {
			return err
		}

	}

	return nil
}

func (m *FullLot) validateInFavorites(formats strfmt.Registry) error {

	if err := validate.Required("in_favorites", "body", m.InFavorites); err != nil {
		return err
	}

	return nil
}

func (m *FullLot) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("created_at", "body", m.CreatedAt); err != nil {
		return err
	}

	return nil
}

func (m *FullLot) validateUser(formats strfmt.Registry) error {

	if err := validate.Required("user", "body", m.User); err != nil {
		return err
	}

	if m.User != nil {
		if err := m.User.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("user")
			}
			return err
		}
	}

	return nil
}

func (m *FullLot) validateExtra(formats strfmt.Registry) error {

	if err := validate.Required("extra", "body", m.Extra); err != nil {
		return err
	}

	for i := 0; i < len(m.Extra); i++ {
		if swag.IsZero(m.Extra[i]) { // not required
			continue
		}

		if m.Extra[i] != nil {
			if err := m.Extra[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("extra" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FullLot) validateViews(formats strfmt.Registry) error {

	if err := validate.Required("views", "body", m.Views); err != nil {
		return err
	}

	return nil
}

func (m *FullLot) validateTgstatLink(formats strfmt.Registry) error {

	if err := validate.Required("tgstat_link", "body", m.TgstatLink); err != nil {
		return err
	}

	return nil
}

func (m *FullLot) validateTelemetrLink(formats strfmt.Registry) error {

	if err := validate.Required("telemetr_link", "body", m.TelemetrLink); err != nil {
		return err
	}

	return nil
}

func (m *FullLot) validateFiles(formats strfmt.Registry) error {

	if err := validate.Required("files", "body", m.Files); err != nil {
		return err
	}

	for i := 0; i < len(m.Files); i++ {
		if swag.IsZero(m.Files[i]) { // not required
			continue
		}

		if m.Files[i] != nil {
			if err := m.Files[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("files" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *FullLot) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FullLot) UnmarshalBinary(b []byte) error {
	var res FullLot
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
