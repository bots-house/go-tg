// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewAdminGetLotStatusesParams creates a new AdminGetLotStatusesParams object
// no default values defined in spec.
func NewAdminGetLotStatusesParams() AdminGetLotStatusesParams {

	return AdminGetLotStatusesParams{}
}

// AdminGetLotStatusesParams contains all the bound params for the admin get lot statuses operation
// typically these are obtained from a http.Request
//
// swagger:parameters adminGetLotStatuses
type AdminGetLotStatusesParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*ID пользователя.
	  In: query
	*/
	UserID *int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewAdminGetLotStatusesParams() beforehand.
func (o *AdminGetLotStatusesParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qUserID, qhkUserID, _ := qs.GetOK("user_id")
	if err := o.bindUserID(qUserID, qhkUserID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindUserID binds and validates parameter UserID from query.
func (o *AdminGetLotStatusesParams) bindUserID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("user_id", "query", "int64", raw)
	}
	o.UserID = &value

	return nil
}
