// Code generated by go-swagger; DO NOT EDIT.

package bot

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// HandleUpdateOKCode is the HTTP code returned for type HandleUpdateOK
const HandleUpdateOKCode int = 200

/*HandleUpdateOK OK

swagger:response handleUpdateOK
*/
type HandleUpdateOK struct {
}

// NewHandleUpdateOK creates HandleUpdateOK with default headers values
func NewHandleUpdateOK() *HandleUpdateOK {

	return &HandleUpdateOK{}
}

// WriteResponse to the client
func (o *HandleUpdateOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// HandleUpdateUnauthorizedCode is the HTTP code returned for type HandleUpdateUnauthorized
const HandleUpdateUnauthorizedCode int = 401

/*HandleUpdateUnauthorized Unauthorized

swagger:response handleUpdateUnauthorized
*/
type HandleUpdateUnauthorized struct {
}

// NewHandleUpdateUnauthorized creates HandleUpdateUnauthorized with default headers values
func NewHandleUpdateUnauthorized() *HandleUpdateUnauthorized {

	return &HandleUpdateUnauthorized{}
}

// WriteResponse to the client
func (o *HandleUpdateUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// HandleUpdateInternalServerErrorCode is the HTTP code returned for type HandleUpdateInternalServerError
const HandleUpdateInternalServerErrorCode int = 500

/*HandleUpdateInternalServerError Internal Server Error

swagger:response handleUpdateInternalServerError
*/
type HandleUpdateInternalServerError struct {
}

// NewHandleUpdateInternalServerError creates HandleUpdateInternalServerError with default headers values
func NewHandleUpdateInternalServerError() *HandleUpdateInternalServerError {

	return &HandleUpdateInternalServerError{}
}

// WriteResponse to the client
func (o *HandleUpdateInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
