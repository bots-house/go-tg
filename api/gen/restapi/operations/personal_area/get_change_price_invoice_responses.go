// Code generated by go-swagger; DO NOT EDIT.

package personal_area

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/bots-house/birzzha/api/gen/models"
)

// GetChangePriceInvoiceOKCode is the HTTP code returned for type GetChangePriceInvoiceOK
const GetChangePriceInvoiceOKCode int = 200

/*GetChangePriceInvoiceOK OK

swagger:response getChangePriceInvoiceOK
*/
type GetChangePriceInvoiceOK struct {

	/*
	  In: Body
	*/
	Payload *models.ChangePriceInvoice `json:"body,omitempty"`
}

// NewGetChangePriceInvoiceOK creates GetChangePriceInvoiceOK with default headers values
func NewGetChangePriceInvoiceOK() *GetChangePriceInvoiceOK {

	return &GetChangePriceInvoiceOK{}
}

// WithPayload adds the payload to the get change price invoice o k response
func (o *GetChangePriceInvoiceOK) WithPayload(payload *models.ChangePriceInvoice) *GetChangePriceInvoiceOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get change price invoice o k response
func (o *GetChangePriceInvoiceOK) SetPayload(payload *models.ChangePriceInvoice) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetChangePriceInvoiceOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetChangePriceInvoiceBadRequestCode is the HTTP code returned for type GetChangePriceInvoiceBadRequest
const GetChangePriceInvoiceBadRequestCode int = 400

/*GetChangePriceInvoiceBadRequest Bad Request

swagger:response getChangePriceInvoiceBadRequest
*/
type GetChangePriceInvoiceBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetChangePriceInvoiceBadRequest creates GetChangePriceInvoiceBadRequest with default headers values
func NewGetChangePriceInvoiceBadRequest() *GetChangePriceInvoiceBadRequest {

	return &GetChangePriceInvoiceBadRequest{}
}

// WithPayload adds the payload to the get change price invoice bad request response
func (o *GetChangePriceInvoiceBadRequest) WithPayload(payload *models.Error) *GetChangePriceInvoiceBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get change price invoice bad request response
func (o *GetChangePriceInvoiceBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetChangePriceInvoiceBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetChangePriceInvoiceInternalServerErrorCode is the HTTP code returned for type GetChangePriceInvoiceInternalServerError
const GetChangePriceInvoiceInternalServerErrorCode int = 500

/*GetChangePriceInvoiceInternalServerError Internal Server Error

swagger:response getChangePriceInvoiceInternalServerError
*/
type GetChangePriceInvoiceInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetChangePriceInvoiceInternalServerError creates GetChangePriceInvoiceInternalServerError with default headers values
func NewGetChangePriceInvoiceInternalServerError() *GetChangePriceInvoiceInternalServerError {

	return &GetChangePriceInvoiceInternalServerError{}
}

// WithPayload adds the payload to the get change price invoice internal server error response
func (o *GetChangePriceInvoiceInternalServerError) WithPayload(payload *models.Error) *GetChangePriceInvoiceInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get change price invoice internal server error response
func (o *GetChangePriceInvoiceInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetChangePriceInvoiceInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
