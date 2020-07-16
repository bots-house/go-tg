// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/bots-house/birzzha/api/authz"
)

// AdminGetUsersHandlerFunc turns a function with the right signature into a admin get users handler
type AdminGetUsersHandlerFunc func(AdminGetUsersParams, *authz.Identity) middleware.Responder

// Handle executing the request and returning a response
func (fn AdminGetUsersHandlerFunc) Handle(params AdminGetUsersParams, principal *authz.Identity) middleware.Responder {
	return fn(params, principal)
}

// AdminGetUsersHandler interface for that can handle valid admin get users params
type AdminGetUsersHandler interface {
	Handle(AdminGetUsersParams, *authz.Identity) middleware.Responder
}

// NewAdminGetUsers creates a new http.Handler for the admin get users operation
func NewAdminGetUsers(ctx *middleware.Context, handler AdminGetUsersHandler) *AdminGetUsers {
	return &AdminGetUsers{Context: ctx, Handler: handler}
}

/*AdminGetUsers swagger:route GET /admin/users admin adminGetUsers

Get Users

Получить список пользователей.

*/
type AdminGetUsers struct {
	Context *middleware.Context
	Handler AdminGetUsersHandler
}

func (o *AdminGetUsers) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewAdminGetUsersParams()

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
