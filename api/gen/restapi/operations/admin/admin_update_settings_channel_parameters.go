// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/bots-house/birzzha/api/gen/models"
)

// NewAdminUpdateSettingsChannelParams creates a new AdminUpdateSettingsChannelParams object
// no default values defined in spec.
func NewAdminUpdateSettingsChannelParams() AdminUpdateSettingsChannelParams {

	return AdminUpdateSettingsChannelParams{}
}

// AdminUpdateSettingsChannelParams contains all the bound params for the admin update settings channel operation
// typically these are obtained from a http.Request
//
// swagger:parameters adminUpdateSettingsChannel
type AdminUpdateSettingsChannelParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Настройки канала.
	  Required: true
	  In: body
	*/
	Channel *models.AdminSettingsChannel
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewAdminUpdateSettingsChannelParams() beforehand.
func (o *AdminUpdateSettingsChannelParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.AdminSettingsChannel
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("channel", "body", ""))
			} else {
				res = append(res, errors.NewParseError("channel", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Channel = &body
			}
		}
	} else {
		res = append(res, errors.Required("channel", "body", ""))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}