// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/bots-house/birzzha/api/gen/models"
)

// ToggleUserAdminOKCode is the HTTP code returned for type ToggleUserAdminOK
const ToggleUserAdminOKCode int = 200

/*ToggleUserAdminOK OK

swagger:response toggleUserAdminOK
*/
type ToggleUserAdminOK struct {

	/*
	  In: Body
	*/
	Payload *models.AdminFullUser `json:"body,omitempty"`
}

// NewToggleUserAdminOK creates ToggleUserAdminOK with default headers values
func NewToggleUserAdminOK() *ToggleUserAdminOK {

	return &ToggleUserAdminOK{}
}

// WithPayload adds the payload to the toggle user admin o k response
func (o *ToggleUserAdminOK) WithPayload(payload *models.AdminFullUser) *ToggleUserAdminOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the toggle user admin o k response
func (o *ToggleUserAdminOK) SetPayload(payload *models.AdminFullUser) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ToggleUserAdminOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ToggleUserAdminBadRequestCode is the HTTP code returned for type ToggleUserAdminBadRequest
const ToggleUserAdminBadRequestCode int = 400

/*ToggleUserAdminBadRequest Bad Request

swagger:response toggleUserAdminBadRequest
*/
type ToggleUserAdminBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewToggleUserAdminBadRequest creates ToggleUserAdminBadRequest with default headers values
func NewToggleUserAdminBadRequest() *ToggleUserAdminBadRequest {

	return &ToggleUserAdminBadRequest{}
}

// WithPayload adds the payload to the toggle user admin bad request response
func (o *ToggleUserAdminBadRequest) WithPayload(payload *models.Error) *ToggleUserAdminBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the toggle user admin bad request response
func (o *ToggleUserAdminBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ToggleUserAdminBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ToggleUserAdminInternalServerErrorCode is the HTTP code returned for type ToggleUserAdminInternalServerError
const ToggleUserAdminInternalServerErrorCode int = 500

/*ToggleUserAdminInternalServerError Internal Server Error

swagger:response toggleUserAdminInternalServerError
*/
type ToggleUserAdminInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewToggleUserAdminInternalServerError creates ToggleUserAdminInternalServerError with default headers values
func NewToggleUserAdminInternalServerError() *ToggleUserAdminInternalServerError {

	return &ToggleUserAdminInternalServerError{}
}

// WithPayload adds the payload to the toggle user admin internal server error response
func (o *ToggleUserAdminInternalServerError) WithPayload(payload *models.Error) *ToggleUserAdminInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the toggle user admin internal server error response
func (o *ToggleUserAdminInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ToggleUserAdminInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
