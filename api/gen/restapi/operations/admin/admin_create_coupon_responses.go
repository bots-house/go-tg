// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/bots-house/birzzha/api/gen/models"
)

// AdminCreateCouponCreatedCode is the HTTP code returned for type AdminCreateCouponCreated
const AdminCreateCouponCreatedCode int = 201

/*AdminCreateCouponCreated Created

swagger:response adminCreateCouponCreated
*/
type AdminCreateCouponCreated struct {

	/*
	  In: Body
	*/
	Payload *models.CouponItem `json:"body,omitempty"`
}

// NewAdminCreateCouponCreated creates AdminCreateCouponCreated with default headers values
func NewAdminCreateCouponCreated() *AdminCreateCouponCreated {

	return &AdminCreateCouponCreated{}
}

// WithPayload adds the payload to the admin create coupon created response
func (o *AdminCreateCouponCreated) WithPayload(payload *models.CouponItem) *AdminCreateCouponCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the admin create coupon created response
func (o *AdminCreateCouponCreated) SetPayload(payload *models.CouponItem) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AdminCreateCouponCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AdminCreateCouponBadRequestCode is the HTTP code returned for type AdminCreateCouponBadRequest
const AdminCreateCouponBadRequestCode int = 400

/*AdminCreateCouponBadRequest Bad Request

swagger:response adminCreateCouponBadRequest
*/
type AdminCreateCouponBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewAdminCreateCouponBadRequest creates AdminCreateCouponBadRequest with default headers values
func NewAdminCreateCouponBadRequest() *AdminCreateCouponBadRequest {

	return &AdminCreateCouponBadRequest{}
}

// WithPayload adds the payload to the admin create coupon bad request response
func (o *AdminCreateCouponBadRequest) WithPayload(payload *models.Error) *AdminCreateCouponBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the admin create coupon bad request response
func (o *AdminCreateCouponBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AdminCreateCouponBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AdminCreateCouponInternalServerErrorCode is the HTTP code returned for type AdminCreateCouponInternalServerError
const AdminCreateCouponInternalServerErrorCode int = 500

/*AdminCreateCouponInternalServerError Internal Server Error

swagger:response adminCreateCouponInternalServerError
*/
type AdminCreateCouponInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewAdminCreateCouponInternalServerError creates AdminCreateCouponInternalServerError with default headers values
func NewAdminCreateCouponInternalServerError() *AdminCreateCouponInternalServerError {

	return &AdminCreateCouponInternalServerError{}
}

// WithPayload adds the payload to the admin create coupon internal server error response
func (o *AdminCreateCouponInternalServerError) WithPayload(payload *models.Error) *AdminCreateCouponInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the admin create coupon internal server error response
func (o *AdminCreateCouponInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AdminCreateCouponInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
