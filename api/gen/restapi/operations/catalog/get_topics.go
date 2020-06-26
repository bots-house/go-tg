// Code generated by go-swagger; DO NOT EDIT.

package catalog

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetTopicsHandlerFunc turns a function with the right signature into a get topics handler
type GetTopicsHandlerFunc func(GetTopicsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetTopicsHandlerFunc) Handle(params GetTopicsParams) middleware.Responder {
	return fn(params)
}

// GetTopicsHandler interface for that can handle valid get topics params
type GetTopicsHandler interface {
	Handle(GetTopicsParams) middleware.Responder
}

// NewGetTopics creates a new http.Handler for the get topics operation
func NewGetTopics(ctx *middleware.Context, handler GetTopicsHandler) *GetTopics {
	return &GetTopics{Context: ctx, Handler: handler}
}

/*GetTopics swagger:route GET /topics catalog getTopics

Get Topics

Получить список категорий.

*/
type GetTopics struct {
	Context *middleware.Context
	Handler GetTopicsHandler
}

func (o *GetTopics) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetTopicsParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}