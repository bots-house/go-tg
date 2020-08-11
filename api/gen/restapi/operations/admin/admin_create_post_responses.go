// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/bots-house/birzzha/api/gen/models"
)

// AdminCreatePostCreatedCode is the HTTP code returned for type AdminCreatePostCreated
const AdminCreatePostCreatedCode int = 201

/*AdminCreatePostCreated Created

swagger:response adminCreatePostCreated
*/
type AdminCreatePostCreated struct {

	/*
	  In: Body
	*/
	Payload *models.Post `json:"body,omitempty"`
}

// NewAdminCreatePostCreated creates AdminCreatePostCreated with default headers values
func NewAdminCreatePostCreated() *AdminCreatePostCreated {

	return &AdminCreatePostCreated{}
}

// WithPayload adds the payload to the admin create post created response
func (o *AdminCreatePostCreated) WithPayload(payload *models.Post) *AdminCreatePostCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the admin create post created response
func (o *AdminCreatePostCreated) SetPayload(payload *models.Post) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AdminCreatePostCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AdminCreatePostBadRequestCode is the HTTP code returned for type AdminCreatePostBadRequest
const AdminCreatePostBadRequestCode int = 400

/*AdminCreatePostBadRequest Bad Request

swagger:response adminCreatePostBadRequest
*/
type AdminCreatePostBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewAdminCreatePostBadRequest creates AdminCreatePostBadRequest with default headers values
func NewAdminCreatePostBadRequest() *AdminCreatePostBadRequest {

	return &AdminCreatePostBadRequest{}
}

// WithPayload adds the payload to the admin create post bad request response
func (o *AdminCreatePostBadRequest) WithPayload(payload *models.Error) *AdminCreatePostBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the admin create post bad request response
func (o *AdminCreatePostBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AdminCreatePostBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AdminCreatePostInternalServerErrorCode is the HTTP code returned for type AdminCreatePostInternalServerError
const AdminCreatePostInternalServerErrorCode int = 500

/*AdminCreatePostInternalServerError Internal Server Error

swagger:response adminCreatePostInternalServerError
*/
type AdminCreatePostInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewAdminCreatePostInternalServerError creates AdminCreatePostInternalServerError with default headers values
func NewAdminCreatePostInternalServerError() *AdminCreatePostInternalServerError {

	return &AdminCreatePostInternalServerError{}
}

// WithPayload adds the payload to the admin create post internal server error response
func (o *AdminCreatePostInternalServerError) WithPayload(payload *models.Error) *AdminCreatePostInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the admin create post internal server error response
func (o *AdminCreatePostInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AdminCreatePostInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}