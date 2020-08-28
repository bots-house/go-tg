// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/bots-house/birzzha/api/authz"
)

// AdminCancelLotHandlerFunc turns a function with the right signature into a admin cancel lot handler
type AdminCancelLotHandlerFunc func(AdminCancelLotParams, *authz.Identity) middleware.Responder

// Handle executing the request and returning a response
func (fn AdminCancelLotHandlerFunc) Handle(params AdminCancelLotParams, principal *authz.Identity) middleware.Responder {
	return fn(params, principal)
}

// AdminCancelLotHandler interface for that can handle valid admin cancel lot params
type AdminCancelLotHandler interface {
	Handle(AdminCancelLotParams, *authz.Identity) middleware.Responder
}

// NewAdminCancelLot creates a new http.Handler for the admin cancel lot operation
func NewAdminCancelLot(ctx *middleware.Context, handler AdminCancelLotHandler) *AdminCancelLot {
	return &AdminCancelLot{Context: ctx, Handler: handler}
}

/*AdminCancelLot swagger:route POST /admin/lots/{id}/cancel admin adminCancelLot

Cancel Lot

Снять лот с продажи


*/
type AdminCancelLot struct {
	Context *middleware.Context
	Handler AdminCancelLotHandler
}

func (o *AdminCancelLot) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewAdminCancelLotParams()

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