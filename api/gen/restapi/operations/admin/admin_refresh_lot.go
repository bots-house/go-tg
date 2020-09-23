// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/bots-house/birzzha/api/authz"
)

// AdminRefreshLotHandlerFunc turns a function with the right signature into a admin refresh lot handler
type AdminRefreshLotHandlerFunc func(AdminRefreshLotParams, *authz.Identity) middleware.Responder

// Handle executing the request and returning a response
func (fn AdminRefreshLotHandlerFunc) Handle(params AdminRefreshLotParams, principal *authz.Identity) middleware.Responder {
	return fn(params, principal)
}

// AdminRefreshLotHandler interface for that can handle valid admin refresh lot params
type AdminRefreshLotHandler interface {
	Handle(AdminRefreshLotParams, *authz.Identity) middleware.Responder
}

// NewAdminRefreshLot creates a new http.Handler for the admin refresh lot operation
func NewAdminRefreshLot(ctx *middleware.Context, handler AdminRefreshLotHandler) *AdminRefreshLot {
	return &AdminRefreshLot{Context: ctx, Handler: handler}
}

/*AdminRefreshLot swagger:route POST /admin/lots/{id}/refresh admin adminRefreshLot

Refresh Lot Details

Обновить данные лота.

*/
type AdminRefreshLot struct {
	Context *middleware.Context
	Handler AdminRefreshLotHandler
}

func (o *AdminRefreshLot) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewAdminRefreshLotParams()

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