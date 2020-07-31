// Code generated by go-swagger; DO NOT EDIT.

package catalog

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

// NewGetDailyCoverageParams creates a new GetDailyCoverageParams object
// no default values defined in spec.
func NewGetDailyCoverageParams() GetDailyCoverageParams {

	return GetDailyCoverageParams{}
}

// GetDailyCoverageParams contains all the bound params for the get daily coverage operation
// typically these are obtained from a http.Request
//
// swagger:parameters getDailyCoverage
type GetDailyCoverageParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*id канала
	  Required: true
	  In: query
	*/
	ChannelID int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetDailyCoverageParams() beforehand.
func (o *GetDailyCoverageParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qChannelID, qhkChannelID, _ := qs.GetOK("channel_id")
	if err := o.bindChannelID(qChannelID, qhkChannelID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindChannelID binds and validates parameter ChannelID from query.
func (o *GetDailyCoverageParams) bindChannelID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("channel_id", "query", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false
	if err := validate.RequiredString("channel_id", "query", raw); err != nil {
		return err
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("channel_id", "query", "int64", raw)
	}
	o.ChannelID = value

	return nil
}
