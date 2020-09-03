// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CouponItem coupon item
//
// swagger:model CouponItem
type CouponItem struct {

	// ID купона.
	// Required: true
	ID *int64 `json:"id"`

	// Код купона.
	// Required: true
	Code *string `json:"code"`

	// Скидка (пример: 0.1 = 10%)
	//
	// Required: true
	Discount *float64 `json:"discount"`

	// Для каких типов платежей применим.
	// Required: true
	// Max Items: 2
	// Min Items: 1
	Purposes []string `json:"purposes"`

	// Дата истечения купона.
	// Required: true
	ExpireAt *int64 `json:"expire_at"`

	// Дата создания купона.
	// Required: true
	CreatedAt *int64 `json:"created_at"`

	// Количество применений.
	// Required: true
	AppliesCount *int64 `json:"applies_count"`

	// Сколько может применятся одним пользователем.
	// Required: true
	MaxAppliesByUserLimit *int64 `json:"max_applies_by_user_limit"`

	// Количество применений всеми пользователями.
	// Required: true
	MaxAppliesLimit *int64 `json:"max_applies_limit"`
}

// Validate validates this coupon item
func (m *CouponItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDiscount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePurposes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExpireAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAppliesCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaxAppliesByUserLimit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaxAppliesLimit(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CouponItem) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *CouponItem) validateCode(formats strfmt.Registry) error {

	if err := validate.Required("code", "body", m.Code); err != nil {
		return err
	}

	return nil
}

func (m *CouponItem) validateDiscount(formats strfmt.Registry) error {

	if err := validate.Required("discount", "body", m.Discount); err != nil {
		return err
	}

	return nil
}

var couponItemPurposesItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["application","change_price"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		couponItemPurposesItemsEnum = append(couponItemPurposesItemsEnum, v)
	}
}

func (m *CouponItem) validatePurposesItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, couponItemPurposesItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CouponItem) validatePurposes(formats strfmt.Registry) error {

	if err := validate.Required("purposes", "body", m.Purposes); err != nil {
		return err
	}

	iPurposesSize := int64(len(m.Purposes))

	if err := validate.MinItems("purposes", "body", iPurposesSize, 1); err != nil {
		return err
	}

	if err := validate.MaxItems("purposes", "body", iPurposesSize, 2); err != nil {
		return err
	}

	for i := 0; i < len(m.Purposes); i++ {

		// value enum
		if err := m.validatePurposesItemsEnum("purposes"+"."+strconv.Itoa(i), "body", m.Purposes[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *CouponItem) validateExpireAt(formats strfmt.Registry) error {

	if err := validate.Required("expire_at", "body", m.ExpireAt); err != nil {
		return err
	}

	return nil
}

func (m *CouponItem) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("created_at", "body", m.CreatedAt); err != nil {
		return err
	}

	return nil
}

func (m *CouponItem) validateAppliesCount(formats strfmt.Registry) error {

	if err := validate.Required("applies_count", "body", m.AppliesCount); err != nil {
		return err
	}

	return nil
}

func (m *CouponItem) validateMaxAppliesByUserLimit(formats strfmt.Registry) error {

	if err := validate.Required("max_applies_by_user_limit", "body", m.MaxAppliesByUserLimit); err != nil {
		return err
	}

	return nil
}

func (m *CouponItem) validateMaxAppliesLimit(formats strfmt.Registry) error {

	if err := validate.Required("max_applies_limit", "body", m.MaxAppliesLimit); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CouponItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CouponItem) UnmarshalBinary(b []byte) error {
	var res CouponItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
