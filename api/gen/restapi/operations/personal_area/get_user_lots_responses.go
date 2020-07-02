// Code generated by go-swagger; DO NOT EDIT.

package personal_area

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/bots-house/birzzha/api/gen/models"
)

// GetUserLotsOKCode is the HTTP code returned for type GetUserLotsOK
const GetUserLotsOKCode int = 200

/*GetUserLotsOK OK

swagger:response getUserLotsOK
*/
type GetUserLotsOK struct {

	/*
	  In: Body
	*/
	Payload []*models.OwnedLot `json:"body,omitempty"`
}

// NewGetUserLotsOK creates GetUserLotsOK with default headers values
func NewGetUserLotsOK() *GetUserLotsOK {

	return &GetUserLotsOK{}
}

// WithPayload adds the payload to the get user lots o k response
func (o *GetUserLotsOK) WithPayload(payload []*models.OwnedLot) *GetUserLotsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get user lots o k response
func (o *GetUserLotsOK) SetPayload(payload []*models.OwnedLot) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUserLotsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.OwnedLot, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetUserLotsBadRequestCode is the HTTP code returned for type GetUserLotsBadRequest
const GetUserLotsBadRequestCode int = 400

/*GetUserLotsBadRequest Bad Request

swagger:response getUserLotsBadRequest
*/
type GetUserLotsBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetUserLotsBadRequest creates GetUserLotsBadRequest with default headers values
func NewGetUserLotsBadRequest() *GetUserLotsBadRequest {

	return &GetUserLotsBadRequest{}
}

// WithPayload adds the payload to the get user lots bad request response
func (o *GetUserLotsBadRequest) WithPayload(payload *models.Error) *GetUserLotsBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get user lots bad request response
func (o *GetUserLotsBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUserLotsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetUserLotsInternalServerErrorCode is the HTTP code returned for type GetUserLotsInternalServerError
const GetUserLotsInternalServerErrorCode int = 500

/*GetUserLotsInternalServerError Internal Server Error

swagger:response getUserLotsInternalServerError
*/
type GetUserLotsInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetUserLotsInternalServerError creates GetUserLotsInternalServerError with default headers values
func NewGetUserLotsInternalServerError() *GetUserLotsInternalServerError {

	return &GetUserLotsInternalServerError{}
}

// WithPayload adds the payload to the get user lots internal server error response
func (o *GetUserLotsInternalServerError) WithPayload(payload *models.Error) *GetUserLotsInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get user lots internal server error response
func (o *GetUserLotsInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUserLotsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}