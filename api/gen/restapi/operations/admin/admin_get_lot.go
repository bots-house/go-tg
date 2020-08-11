// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/bots-house/birzzha/api/authz"
)

// AdminGetLotHandlerFunc turns a function with the right signature into a admin get lot handler
type AdminGetLotHandlerFunc func(AdminGetLotParams, *authz.Identity) middleware.Responder

// Handle executing the request and returning a response
func (fn AdminGetLotHandlerFunc) Handle(params AdminGetLotParams, principal *authz.Identity) middleware.Responder {
	return fn(params, principal)
}

// AdminGetLotHandler interface for that can handle valid admin get lot params
type AdminGetLotHandler interface {
	Handle(AdminGetLotParams, *authz.Identity) middleware.Responder
}

// NewAdminGetLot creates a new http.Handler for the admin get lot operation
func NewAdminGetLot(ctx *middleware.Context, handler AdminGetLotHandler) *AdminGetLot {
	return &AdminGetLot{Context: ctx, Handler: handler}
}

/*AdminGetLot swagger:route GET /admin/lots/{id} admin adminGetLot

Get Lot

Получить детали лота.

*/
type AdminGetLot struct {
	Context *middleware.Context
	Handler AdminGetLotHandler
}

func (o *AdminGetLot) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewAdminGetLotParams()

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