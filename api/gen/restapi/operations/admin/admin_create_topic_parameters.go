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

// NewAdminCreateTopicParams creates a new AdminCreateTopicParams object
// no default values defined in spec.
func NewAdminCreateTopicParams() AdminCreateTopicParams {

	return AdminCreateTopicParams{}
}

// AdminCreateTopicParams contains all the bound params for the admin create topic operation
// typically these are obtained from a http.Request
//
// swagger:parameters adminCreateTopic
type AdminCreateTopicParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Категория.
	  Required: true
	  In: body
	*/
	Topic *models.InputTopic
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewAdminCreateTopicParams() beforehand.
func (o *AdminCreateTopicParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.InputTopic
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("topic", "body", ""))
			} else {
				res = append(res, errors.NewParseError("topic", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Topic = &body
			}
		}
	} else {
		res = append(res, errors.Required("topic", "body", ""))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}