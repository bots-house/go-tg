// Code generated by go-swagger; DO NOT EDIT.

package personal_area

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/bots-house/birzzha/api/gen/models"
)

// GetPaymentStatusOKCode is the HTTP code returned for type GetPaymentStatusOK
const GetPaymentStatusOKCode int = 200

/*GetPaymentStatusOK OK

swagger:response getPaymentStatusOK
*/
type GetPaymentStatusOK struct {

	/*
	  In: Body
	*/
	Payload *models.PaymentStatus `json:"body,omitempty"`
}

// NewGetPaymentStatusOK creates GetPaymentStatusOK with default headers values
func NewGetPaymentStatusOK() *GetPaymentStatusOK {

	return &GetPaymentStatusOK{}
}

// WithPayload adds the payload to the get payment status o k response
func (o *GetPaymentStatusOK) WithPayload(payload *models.PaymentStatus) *GetPaymentStatusOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get payment status o k response
func (o *GetPaymentStatusOK) SetPayload(payload *models.PaymentStatus) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPaymentStatusOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetPaymentStatusBadRequestCode is the HTTP code returned for type GetPaymentStatusBadRequest
const GetPaymentStatusBadRequestCode int = 400

/*GetPaymentStatusBadRequest Bad Request

swagger:response getPaymentStatusBadRequest
*/
type GetPaymentStatusBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetPaymentStatusBadRequest creates GetPaymentStatusBadRequest with default headers values
func NewGetPaymentStatusBadRequest() *GetPaymentStatusBadRequest {

	return &GetPaymentStatusBadRequest{}
}

// WithPayload adds the payload to the get payment status bad request response
func (o *GetPaymentStatusBadRequest) WithPayload(payload *models.Error) *GetPaymentStatusBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get payment status bad request response
func (o *GetPaymentStatusBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPaymentStatusBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetPaymentStatusInternalServerErrorCode is the HTTP code returned for type GetPaymentStatusInternalServerError
const GetPaymentStatusInternalServerErrorCode int = 500

/*GetPaymentStatusInternalServerError Internal Server Error

swagger:response getPaymentStatusInternalServerError
*/
type GetPaymentStatusInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetPaymentStatusInternalServerError creates GetPaymentStatusInternalServerError with default headers values
func NewGetPaymentStatusInternalServerError() *GetPaymentStatusInternalServerError {

	return &GetPaymentStatusInternalServerError{}
}

// WithPayload adds the payload to the get payment status internal server error response
func (o *GetPaymentStatusInternalServerError) WithPayload(payload *models.Error) *GetPaymentStatusInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get payment status internal server error response
func (o *GetPaymentStatusInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPaymentStatusInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
