// Code generated by go-swagger; DO NOT EDIT.

package catalog

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/bots-house/birzzha/api/gen/models"
)

// GetLotsOKCode is the HTTP code returned for type GetLotsOK
const GetLotsOKCode int = 200

/*GetLotsOK OK

swagger:response getLotsOK
*/
type GetLotsOK struct {

	/*
	  In: Body
	*/
	Payload []*models.LotListItem `json:"body,omitempty"`
}

// NewGetLotsOK creates GetLotsOK with default headers values
func NewGetLotsOK() *GetLotsOK {

	return &GetLotsOK{}
}

// WithPayload adds the payload to the get lots o k response
func (o *GetLotsOK) WithPayload(payload []*models.LotListItem) *GetLotsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get lots o k response
func (o *GetLotsOK) SetPayload(payload []*models.LotListItem) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetLotsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.LotListItem, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetLotsBadRequestCode is the HTTP code returned for type GetLotsBadRequest
const GetLotsBadRequestCode int = 400

/*GetLotsBadRequest Bad Request

swagger:response getLotsBadRequest
*/
type GetLotsBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetLotsBadRequest creates GetLotsBadRequest with default headers values
func NewGetLotsBadRequest() *GetLotsBadRequest {

	return &GetLotsBadRequest{}
}

// WithPayload adds the payload to the get lots bad request response
func (o *GetLotsBadRequest) WithPayload(payload *models.Error) *GetLotsBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get lots bad request response
func (o *GetLotsBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetLotsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetLotsInternalServerErrorCode is the HTTP code returned for type GetLotsInternalServerError
const GetLotsInternalServerErrorCode int = 500

/*GetLotsInternalServerError Internal Server Error

swagger:response getLotsInternalServerError
*/
type GetLotsInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetLotsInternalServerError creates GetLotsInternalServerError with default headers values
func NewGetLotsInternalServerError() *GetLotsInternalServerError {

	return &GetLotsInternalServerError{}
}

// WithPayload adds the payload to the get lots internal server error response
func (o *GetLotsInternalServerError) WithPayload(payload *models.Error) *GetLotsInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get lots internal server error response
func (o *GetLotsInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetLotsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}