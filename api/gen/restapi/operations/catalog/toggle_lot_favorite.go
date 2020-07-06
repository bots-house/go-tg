// Code generated by go-swagger; DO NOT EDIT.

package catalog

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/bots-house/birzzha/api/authz"
)

// ToggleLotFavoriteHandlerFunc turns a function with the right signature into a toggle lot favorite handler
type ToggleLotFavoriteHandlerFunc func(ToggleLotFavoriteParams, *authz.Identity) middleware.Responder

// Handle executing the request and returning a response
func (fn ToggleLotFavoriteHandlerFunc) Handle(params ToggleLotFavoriteParams, principal *authz.Identity) middleware.Responder {
	return fn(params, principal)
}

// ToggleLotFavoriteHandler interface for that can handle valid toggle lot favorite params
type ToggleLotFavoriteHandler interface {
	Handle(ToggleLotFavoriteParams, *authz.Identity) middleware.Responder
}

// NewToggleLotFavorite creates a new http.Handler for the toggle lot favorite operation
func NewToggleLotFavorite(ctx *middleware.Context, handler ToggleLotFavoriteHandler) *ToggleLotFavorite {
	return &ToggleLotFavorite{Context: ctx, Handler: handler}
}

/*ToggleLotFavorite swagger:route POST /lots/{id}/favorite catalog toggleLotFavorite

Change Favorite Status

Изменение статуса избранности лота.

*/
type ToggleLotFavorite struct {
	Context *middleware.Context
	Handler ToggleLotFavoriteHandler
}

func (o *ToggleLotFavorite) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewToggleLotFavoriteParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal *authz.Identity
	if uprinc != nil {
		principal = uprinc.(*authz.Identity) // this is really a authz.Identity, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}