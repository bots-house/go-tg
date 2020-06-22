// Code generated by go-swagger; DO NOT EDIT.

package bot

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/bots-house/birzzha/api/gen/models"
)

// GetBotInfoOKCode is the HTTP code returned for type GetBotInfoOK
const GetBotInfoOKCode int = 200

/*GetBotInfoOK OK

swagger:response getBotInfoOK
*/
type GetBotInfoOK struct {

	/*
	  In: Body
	*/
	Payload *models.BotInfo `json:"body,omitempty"`
}

// NewGetBotInfoOK creates GetBotInfoOK with default headers values
func NewGetBotInfoOK() *GetBotInfoOK {

	return &GetBotInfoOK{}
}

// WithPayload adds the payload to the get bot info o k response
func (o *GetBotInfoOK) WithPayload(payload *models.BotInfo) *GetBotInfoOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get bot info o k response
func (o *GetBotInfoOK) SetPayload(payload *models.BotInfo) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetBotInfoOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
