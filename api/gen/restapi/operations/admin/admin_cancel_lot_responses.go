// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/bots-house/birzzha/api/gen/models"
)

// AdminCancelLotOKCode is the HTTP code returned for type AdminCancelLotOK
const AdminCancelLotOKCode int = 200

/*AdminCancelLotOK OK

swagger:response adminCancelLotOK
*/
type AdminCancelLotOK struct {
}

// NewAdminCancelLotOK creates AdminCancelLotOK with default headers values
func NewAdminCancelLotOK() *AdminCancelLotOK {

	return &AdminCancelLotOK{}
}

// WriteResponse to the client
func (o *AdminCancelLotOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// AdminCancelLotBadRequestCode is the HTTP code returned for type AdminCancelLotBadRequest
const AdminCancelLotBadRequestCode int = 400

/*AdminCancelLotBadRequest Bad Request

swagger:response adminCancelLotBadRequest
*/
type AdminCancelLotBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewAdminCancelLotBadRequest creates AdminCancelLotBadRequest with default headers values
func NewAdminCancelLotBadRequest() *AdminCancelLotBadRequest {

	return &AdminCancelLotBadRequest{}
}

// WithPayload adds the payload to the admin cancel lot bad request response
func (o *AdminCancelLotBadRequest) WithPayload(payload *models.Error) *AdminCancelLotBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the admin cancel lot bad request response
func (o *AdminCancelLotBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AdminCancelLotBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AdminCancelLotInternalServerErrorCode is the HTTP code returned for type AdminCancelLotInternalServerError
const AdminCancelLotInternalServerErrorCode int = 500

/*AdminCancelLotInternalServerError Internal Server Error

swagger:response adminCancelLotInternalServerError
*/
type AdminCancelLotInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewAdminCancelLotInternalServerError creates AdminCancelLotInternalServerError with default headers values
func NewAdminCancelLotInternalServerError() *AdminCancelLotInternalServerError {

	return &AdminCancelLotInternalServerError{}
}

// WithPayload adds the payload to the admin cancel lot internal server error response
func (o *AdminCancelLotInternalServerError) WithPayload(payload *models.Error) *AdminCancelLotInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the admin cancel lot internal server error response
func (o *AdminCancelLotInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AdminCancelLotInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}