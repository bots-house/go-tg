// Code generated by go-swagger; DO NOT EDIT.

package catalog

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/bots-house/birzzha/api/gen/models"
)

// GetFilterBoundariesOKCode is the HTTP code returned for type GetFilterBoundariesOK
const GetFilterBoundariesOKCode int = 200

/*GetFilterBoundariesOK OK

swagger:response getFilterBoundariesOK
*/
type GetFilterBoundariesOK struct {

	/*
	  In: Body
	*/
	Payload *models.LotFilterBoundaries `json:"body,omitempty"`
}

// NewGetFilterBoundariesOK creates GetFilterBoundariesOK with default headers values
func NewGetFilterBoundariesOK() *GetFilterBoundariesOK {

	return &GetFilterBoundariesOK{}
}

// WithPayload adds the payload to the get filter boundaries o k response
func (o *GetFilterBoundariesOK) WithPayload(payload *models.LotFilterBoundaries) *GetFilterBoundariesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get filter boundaries o k response
func (o *GetFilterBoundariesOK) SetPayload(payload *models.LotFilterBoundaries) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetFilterBoundariesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetFilterBoundariesInternalServerErrorCode is the HTTP code returned for type GetFilterBoundariesInternalServerError
const GetFilterBoundariesInternalServerErrorCode int = 500

/*GetFilterBoundariesInternalServerError Internal Server Error

swagger:response getFilterBoundariesInternalServerError
*/
type GetFilterBoundariesInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetFilterBoundariesInternalServerError creates GetFilterBoundariesInternalServerError with default headers values
func NewGetFilterBoundariesInternalServerError() *GetFilterBoundariesInternalServerError {

	return &GetFilterBoundariesInternalServerError{}
}

// WithPayload adds the payload to the get filter boundaries internal server error response
func (o *GetFilterBoundariesInternalServerError) WithPayload(payload *models.Error) *GetFilterBoundariesInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get filter boundaries internal server error response
func (o *GetFilterBoundariesInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetFilterBoundariesInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
