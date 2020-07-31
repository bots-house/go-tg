// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/bots-house/birzzha/api/authz"
)

// AdminDeclineLotHandlerFunc turns a function with the right signature into a admin decline lot handler
type AdminDeclineLotHandlerFunc func(AdminDeclineLotParams, *authz.Identity) middleware.Responder

// Handle executing the request and returning a response
func (fn AdminDeclineLotHandlerFunc) Handle(params AdminDeclineLotParams, principal *authz.Identity) middleware.Responder {
	return fn(params, principal)
}

// AdminDeclineLotHandler interface for that can handle valid admin decline lot params
type AdminDeclineLotHandler interface {
	Handle(AdminDeclineLotParams, *authz.Identity) middleware.Responder
}

// NewAdminDeclineLot creates a new http.Handler for the admin decline lot operation
func NewAdminDeclineLot(ctx *middleware.Context, handler AdminDeclineLotHandler) *AdminDeclineLot {
	return &AdminDeclineLot{Context: ctx, Handler: handler}
}

/*AdminDeclineLot swagger:route PUT /admin/lots/{id}/decline admin adminDeclineLot

Decline Lot

Отклонить лот.

*/
type AdminDeclineLot struct {
	Context *middleware.Context
	Handler AdminDeclineLotHandler
}

func (o *AdminDeclineLot) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewAdminDeclineLotParams()

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
