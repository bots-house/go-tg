// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/bots-house/birzzha/api/gen/models"
)

// AdminDeleteCouponNoContentCode is the HTTP code returned for type AdminDeleteCouponNoContent
const AdminDeleteCouponNoContentCode int = 204

/*AdminDeleteCouponNoContent No Content

swagger:response adminDeleteCouponNoContent
*/
type AdminDeleteCouponNoContent struct {
}

// NewAdminDeleteCouponNoContent creates AdminDeleteCouponNoContent with default headers values
func NewAdminDeleteCouponNoContent() *AdminDeleteCouponNoContent {

	return &AdminDeleteCouponNoContent{}
}

// WriteResponse to the client
func (o *AdminDeleteCouponNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// AdminDeleteCouponBadRequestCode is the HTTP code returned for type AdminDeleteCouponBadRequest
const AdminDeleteCouponBadRequestCode int = 400

/*AdminDeleteCouponBadRequest Bad Request

swagger:response adminDeleteCouponBadRequest
*/
type AdminDeleteCouponBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewAdminDeleteCouponBadRequest creates AdminDeleteCouponBadRequest with default headers values
func NewAdminDeleteCouponBadRequest() *AdminDeleteCouponBadRequest {

	return &AdminDeleteCouponBadRequest{}
}

// WithPayload adds the payload to the admin delete coupon bad request response
func (o *AdminDeleteCouponBadRequest) WithPayload(payload *models.Error) *AdminDeleteCouponBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the admin delete coupon bad request response
func (o *AdminDeleteCouponBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AdminDeleteCouponBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AdminDeleteCouponInternalServerErrorCode is the HTTP code returned for type AdminDeleteCouponInternalServerError
const AdminDeleteCouponInternalServerErrorCode int = 500

/*AdminDeleteCouponInternalServerError Internal Server Error

swagger:response adminDeleteCouponInternalServerError
*/
type AdminDeleteCouponInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewAdminDeleteCouponInternalServerError creates AdminDeleteCouponInternalServerError with default headers values
func NewAdminDeleteCouponInternalServerError() *AdminDeleteCouponInternalServerError {

	return &AdminDeleteCouponInternalServerError{}
}

// WithPayload adds the payload to the admin delete coupon internal server error response
func (o *AdminDeleteCouponInternalServerError) WithPayload(payload *models.Error) *AdminDeleteCouponInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the admin delete coupon internal server error response
func (o *AdminDeleteCouponInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AdminDeleteCouponInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}