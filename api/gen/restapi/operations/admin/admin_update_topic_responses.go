// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/bots-house/birzzha/api/gen/models"
)

// AdminUpdateTopicOKCode is the HTTP code returned for type AdminUpdateTopicOK
const AdminUpdateTopicOKCode int = 200

/*AdminUpdateTopicOK OK

swagger:response adminUpdateTopicOK
*/
type AdminUpdateTopicOK struct {

	/*
	  In: Body
	*/
	Payload *models.AdminTopic `json:"body,omitempty"`
}

// NewAdminUpdateTopicOK creates AdminUpdateTopicOK with default headers values
func NewAdminUpdateTopicOK() *AdminUpdateTopicOK {

	return &AdminUpdateTopicOK{}
}

// WithPayload adds the payload to the admin update topic o k response
func (o *AdminUpdateTopicOK) WithPayload(payload *models.AdminTopic) *AdminUpdateTopicOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the admin update topic o k response
func (o *AdminUpdateTopicOK) SetPayload(payload *models.AdminTopic) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AdminUpdateTopicOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AdminUpdateTopicBadRequestCode is the HTTP code returned for type AdminUpdateTopicBadRequest
const AdminUpdateTopicBadRequestCode int = 400

/*AdminUpdateTopicBadRequest Bad Request

swagger:response adminUpdateTopicBadRequest
*/
type AdminUpdateTopicBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewAdminUpdateTopicBadRequest creates AdminUpdateTopicBadRequest with default headers values
func NewAdminUpdateTopicBadRequest() *AdminUpdateTopicBadRequest {

	return &AdminUpdateTopicBadRequest{}
}

// WithPayload adds the payload to the admin update topic bad request response
func (o *AdminUpdateTopicBadRequest) WithPayload(payload *models.Error) *AdminUpdateTopicBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the admin update topic bad request response
func (o *AdminUpdateTopicBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AdminUpdateTopicBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AdminUpdateTopicInternalServerErrorCode is the HTTP code returned for type AdminUpdateTopicInternalServerError
const AdminUpdateTopicInternalServerErrorCode int = 500

/*AdminUpdateTopicInternalServerError Internal Server Error

swagger:response adminUpdateTopicInternalServerError
*/
type AdminUpdateTopicInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewAdminUpdateTopicInternalServerError creates AdminUpdateTopicInternalServerError with default headers values
func NewAdminUpdateTopicInternalServerError() *AdminUpdateTopicInternalServerError {

	return &AdminUpdateTopicInternalServerError{}
}

// WithPayload adds the payload to the admin update topic internal server error response
func (o *AdminUpdateTopicInternalServerError) WithPayload(payload *models.Error) *AdminUpdateTopicInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the admin update topic internal server error response
func (o *AdminUpdateTopicInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AdminUpdateTopicInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}