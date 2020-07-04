// Code generated by go-swagger; DO NOT EDIT.

package catalog

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/bots-house/birzzha/api/gen/models"
)

// GetSimilarLotsOKCode is the HTTP code returned for type GetSimilarLotsOK
const GetSimilarLotsOKCode int = 200

/*GetSimilarLotsOK OK

swagger:response getSimilarLotsOK
*/
type GetSimilarLotsOK struct {

	/*
	  In: Body
	*/
	Payload *models.LotList `json:"body,omitempty"`
}

// NewGetSimilarLotsOK creates GetSimilarLotsOK with default headers values
func NewGetSimilarLotsOK() *GetSimilarLotsOK {

	return &GetSimilarLotsOK{}
}

// WithPayload adds the payload to the get similar lots o k response
func (o *GetSimilarLotsOK) WithPayload(payload *models.LotList) *GetSimilarLotsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get similar lots o k response
func (o *GetSimilarLotsOK) SetPayload(payload *models.LotList) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSimilarLotsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetSimilarLotsBadRequestCode is the HTTP code returned for type GetSimilarLotsBadRequest
const GetSimilarLotsBadRequestCode int = 400

/*GetSimilarLotsBadRequest Bad Request

swagger:response getSimilarLotsBadRequest
*/
type GetSimilarLotsBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetSimilarLotsBadRequest creates GetSimilarLotsBadRequest with default headers values
func NewGetSimilarLotsBadRequest() *GetSimilarLotsBadRequest {

	return &GetSimilarLotsBadRequest{}
}

// WithPayload adds the payload to the get similar lots bad request response
func (o *GetSimilarLotsBadRequest) WithPayload(payload *models.Error) *GetSimilarLotsBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get similar lots bad request response
func (o *GetSimilarLotsBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSimilarLotsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetSimilarLotsInternalServerErrorCode is the HTTP code returned for type GetSimilarLotsInternalServerError
const GetSimilarLotsInternalServerErrorCode int = 500

/*GetSimilarLotsInternalServerError Internal Server Error

swagger:response getSimilarLotsInternalServerError
*/
type GetSimilarLotsInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetSimilarLotsInternalServerError creates GetSimilarLotsInternalServerError with default headers values
func NewGetSimilarLotsInternalServerError() *GetSimilarLotsInternalServerError {

	return &GetSimilarLotsInternalServerError{}
}

// WithPayload adds the payload to the get similar lots internal server error response
func (o *GetSimilarLotsInternalServerError) WithPayload(payload *models.Error) *GetSimilarLotsInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get similar lots internal server error response
func (o *GetSimilarLotsInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSimilarLotsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
