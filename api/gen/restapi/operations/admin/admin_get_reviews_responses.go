// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/bots-house/birzzha/api/gen/models"
)

// AdminGetReviewsOKCode is the HTTP code returned for type AdminGetReviewsOK
const AdminGetReviewsOKCode int = 200

/*AdminGetReviewsOK OK

swagger:response adminGetReviewsOK
*/
type AdminGetReviewsOK struct {

	/*
	  In: Body
	*/
	Payload *models.ReviewList `json:"body,omitempty"`
}

// NewAdminGetReviewsOK creates AdminGetReviewsOK with default headers values
func NewAdminGetReviewsOK() *AdminGetReviewsOK {

	return &AdminGetReviewsOK{}
}

// WithPayload adds the payload to the admin get reviews o k response
func (o *AdminGetReviewsOK) WithPayload(payload *models.ReviewList) *AdminGetReviewsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the admin get reviews o k response
func (o *AdminGetReviewsOK) SetPayload(payload *models.ReviewList) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AdminGetReviewsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AdminGetReviewsBadRequestCode is the HTTP code returned for type AdminGetReviewsBadRequest
const AdminGetReviewsBadRequestCode int = 400

/*AdminGetReviewsBadRequest Bad Request

swagger:response adminGetReviewsBadRequest
*/
type AdminGetReviewsBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewAdminGetReviewsBadRequest creates AdminGetReviewsBadRequest with default headers values
func NewAdminGetReviewsBadRequest() *AdminGetReviewsBadRequest {

	return &AdminGetReviewsBadRequest{}
}

// WithPayload adds the payload to the admin get reviews bad request response
func (o *AdminGetReviewsBadRequest) WithPayload(payload *models.Error) *AdminGetReviewsBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the admin get reviews bad request response
func (o *AdminGetReviewsBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AdminGetReviewsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AdminGetReviewsInternalServerErrorCode is the HTTP code returned for type AdminGetReviewsInternalServerError
const AdminGetReviewsInternalServerErrorCode int = 500

/*AdminGetReviewsInternalServerError Internal Server Error

swagger:response adminGetReviewsInternalServerError
*/
type AdminGetReviewsInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewAdminGetReviewsInternalServerError creates AdminGetReviewsInternalServerError with default headers values
func NewAdminGetReviewsInternalServerError() *AdminGetReviewsInternalServerError {

	return &AdminGetReviewsInternalServerError{}
}

// WithPayload adds the payload to the admin get reviews internal server error response
func (o *AdminGetReviewsInternalServerError) WithPayload(payload *models.Error) *AdminGetReviewsInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the admin get reviews internal server error response
func (o *AdminGetReviewsInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AdminGetReviewsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
