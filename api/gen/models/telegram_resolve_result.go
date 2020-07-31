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

// TelegramResolveResult Сущность соотвествующая запросу.
//
// swagger:model TelegramResolveResult
type TelegramResolveResult struct {

	// channel
	Channel *TelegramResolveResultChannel `json:"channel,omitempty"`

	// group
	Group *TelegramResolveResultGroup `json:"group,omitempty"`
}

// Validate validates this telegram resolve result
func (m *TelegramResolveResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChannel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGroup(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TelegramResolveResult) validateChannel(formats strfmt.Registry) error {

	if swag.IsZero(m.Channel) { // not required
		return nil
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

func (m *TelegramResolveResult) validateGroup(formats strfmt.Registry) error {

	if swag.IsZero(m.Group) { // not required
		return nil
	}

	if m.Group != nil {
		if err := m.Group.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("group")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TelegramResolveResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TelegramResolveResult) UnmarshalBinary(b []byte) error {
	var res TelegramResolveResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TelegramResolveResultChannel Канал
//
// swagger:model TelegramResolveResultChannel
type TelegramResolveResultChannel struct {

	// Уникальный ID канала в Telegramа
	// Required: true
	ID *int64 `json:"id"`

	// Название канала
	// Required: true
	Name *string `json:"name"`

	// URL аватарки канала (может быть null)
	// Required: true
	Avatar *string `json:"avatar"`

	// Описание
	// Required: true
	Description *string `json:"description"`

	// Кол-во участников в канале
	// Required: true
	MembersCount *int64 `json:"members_count"`

	// username канала (может быть null)
	// Required: true
	Username *string `json:"username"`
}

// Validate validates this telegram resolve result channel
func (m *TelegramResolveResultChannel) Validate(formats strfmt.Registry) error {
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

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMembersCount(formats); err != nil {
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

func (m *TelegramResolveResultChannel) validateID(formats strfmt.Registry) error {

	if err := validate.Required("channel"+"."+"id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *TelegramResolveResultChannel) validateName(formats strfmt.Registry) error {

	if err := validate.Required("channel"+"."+"name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *TelegramResolveResultChannel) validateAvatar(formats strfmt.Registry) error {

	if err := validate.Required("channel"+"."+"avatar", "body", m.Avatar); err != nil {
		return err
	}

	return nil
}

func (m *TelegramResolveResultChannel) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("channel"+"."+"description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *TelegramResolveResultChannel) validateMembersCount(formats strfmt.Registry) error {

	if err := validate.Required("channel"+"."+"members_count", "body", m.MembersCount); err != nil {
		return err
	}

	return nil
}

func (m *TelegramResolveResultChannel) validateUsername(formats strfmt.Registry) error {

	if err := validate.Required("channel"+"."+"username", "body", m.Username); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TelegramResolveResultChannel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TelegramResolveResultChannel) UnmarshalBinary(b []byte) error {
	var res TelegramResolveResultChannel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TelegramResolveResultGroup Групповой чат
//
// swagger:model TelegramResolveResultGroup
type TelegramResolveResultGroup struct {

	// Уникальный ID чата в Telegrama
	// Required: true
	ID *int64 `json:"id"`

	// Название чата
	// Required: true
	Name *string `json:"name"`

	// Описание
	// Required: true
	Description *string `json:"description"`

	// URL аватарки чата (может быть null)
	// Required: true
	Avatar *string `json:"avatar"`

	// Кол-во участников в чате
	// Required: true
	MembersCount *int64 `json:"members_count"`

	// username чата (может быть null)
	// Required: true
	Username *string `json:"username"`
}

// Validate validates this telegram resolve result group
func (m *TelegramResolveResultGroup) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAvatar(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMembersCount(formats); err != nil {
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

func (m *TelegramResolveResultGroup) validateID(formats strfmt.Registry) error {

	if err := validate.Required("group"+"."+"id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *TelegramResolveResultGroup) validateName(formats strfmt.Registry) error {

	if err := validate.Required("group"+"."+"name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *TelegramResolveResultGroup) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("group"+"."+"description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *TelegramResolveResultGroup) validateAvatar(formats strfmt.Registry) error {

	if err := validate.Required("group"+"."+"avatar", "body", m.Avatar); err != nil {
		return err
	}

	return nil
}

func (m *TelegramResolveResultGroup) validateMembersCount(formats strfmt.Registry) error {

	if err := validate.Required("group"+"."+"members_count", "body", m.MembersCount); err != nil {
		return err
	}

	return nil
}

func (m *TelegramResolveResultGroup) validateUsername(formats strfmt.Registry) error {

	if err := validate.Required("group"+"."+"username", "body", m.Username); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TelegramResolveResultGroup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TelegramResolveResultGroup) UnmarshalBinary(b []byte) error {
	var res TelegramResolveResultGroup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
