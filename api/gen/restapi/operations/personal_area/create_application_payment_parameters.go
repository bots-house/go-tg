// Code generated by go-swagger; DO NOT EDIT.

package personal_area

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NewCreateApplicationPaymentParams creates a new CreateApplicationPaymentParams object
// no default values defined in spec.
func NewCreateApplicationPaymentParams() CreateApplicationPaymentParams {

	return CreateApplicationPaymentParams{}
}

// CreateApplicationPaymentParams contains all the bound params for the create application payment operation
// typically these are obtained from a http.Request
//
// swagger:parameters createApplicationPayment
type CreateApplicationPaymentParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: query
	*/
	Gateway string
	/*ID неоплаченного лота
	  Required: true
	  In: path
	*/
	ID int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewCreateApplicationPaymentParams() beforehand.
func (o *CreateApplicationPaymentParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qGateway, qhkGateway, _ := qs.GetOK("gateway")
	if err := o.bindGateway(qGateway, qhkGateway, route.Formats); err != nil {
		res = append(res, err)
	}

	rID, rhkID, _ := route.Params.GetOK("id")
	if err := o.bindID(rID, rhkID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindGateway binds and validates parameter Gateway from query.
func (o *CreateApplicationPaymentParams) bindGateway(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("gateway", "query", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false
	if err := validate.RequiredString("gateway", "query", raw); err != nil {
		return err
	}

	o.Gateway = raw

	if err := o.validateGateway(formats); err != nil {
		return err
	}

	return nil
}

// validateGateway carries on validations for parameter Gateway
func (o *CreateApplicationPaymentParams) validateGateway(formats strfmt.Registry) error {

	if err := validate.EnumCase("gateway", "query", o.Gateway, []interface{}{"interkassa", "direct"}, true); err != nil {
		return err
	}

	return nil
}

// bindID binds and validates parameter ID from path.
func (o *CreateApplicationPaymentParams) bindID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("id", "path", "int64", raw)
	}
	o.ID = value

	return nil
}