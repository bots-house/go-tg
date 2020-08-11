// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/bots-house/birzzha/api/gen/models"
)

// AdminUpdateSettingsPricesOKCode is the HTTP code returned for type AdminUpdateSettingsPricesOK
const AdminUpdateSettingsPricesOKCode int = 200

/*AdminUpdateSettingsPricesOK OK

swagger:response adminUpdateSettingsPricesOK
*/
type AdminUpdateSettingsPricesOK struct {

	/*
	  In: Body
	*/
	Payload *models.AdminSettingsPrices `json:"body,omitempty"`
}

// NewAdminUpdateSettingsPricesOK creates AdminUpdateSettingsPricesOK with default headers values
func NewAdminUpdateSettingsPricesOK() *AdminUpdateSettingsPricesOK {

	return &AdminUpdateSettingsPricesOK{}
}

// WithPayload adds the payload to the admin update settings prices o k response
func (o *AdminUpdateSettingsPricesOK) WithPayload(payload *models.AdminSettingsPrices) *AdminUpdateSettingsPricesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the admin update settings prices o k response
func (o *AdminUpdateSettingsPricesOK) SetPayload(payload *models.AdminSettingsPrices) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AdminUpdateSettingsPricesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AdminUpdateSettingsPricesBadRequestCode is the HTTP code returned for type AdminUpdateSettingsPricesBadRequest
const AdminUpdateSettingsPricesBadRequestCode int = 400

/*AdminUpdateSettingsPricesBadRequest Bad Request

swagger:response adminUpdateSettingsPricesBadRequest
*/
type AdminUpdateSettingsPricesBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewAdminUpdateSettingsPricesBadRequest creates AdminUpdateSettingsPricesBadRequest with default headers values
func NewAdminUpdateSettingsPricesBadRequest() *AdminUpdateSettingsPricesBadRequest {

	return &AdminUpdateSettingsPricesBadRequest{}
}

// WithPayload adds the payload to the admin update settings prices bad request response
func (o *AdminUpdateSettingsPricesBadRequest) WithPayload(payload *models.Error) *AdminUpdateSettingsPricesBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the admin update settings prices bad request response
func (o *AdminUpdateSettingsPricesBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AdminUpdateSettingsPricesBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AdminUpdateSettingsPricesInternalServerErrorCode is the HTTP code returned for type AdminUpdateSettingsPricesInternalServerError
const AdminUpdateSettingsPricesInternalServerErrorCode int = 500

/*AdminUpdateSettingsPricesInternalServerError Internal Server Error

swagger:response adminUpdateSettingsPricesInternalServerError
*/
type AdminUpdateSettingsPricesInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewAdminUpdateSettingsPricesInternalServerError creates AdminUpdateSettingsPricesInternalServerError with default headers values
func NewAdminUpdateSettingsPricesInternalServerError() *AdminUpdateSettingsPricesInternalServerError {

	return &AdminUpdateSettingsPricesInternalServerError{}
}

// WithPayload adds the payload to the admin update settings prices internal server error response
func (o *AdminUpdateSettingsPricesInternalServerError) WithPayload(payload *models.Error) *AdminUpdateSettingsPricesInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the admin update settings prices internal server error response
func (o *AdminUpdateSettingsPricesInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AdminUpdateSettingsPricesInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}