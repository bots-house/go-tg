// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/bots-house/birzzha/api/authz"
)

// AdminCreatePostHandlerFunc turns a function with the right signature into a admin create post handler
type AdminCreatePostHandlerFunc func(AdminCreatePostParams, *authz.Identity) middleware.Responder

// Handle executing the request and returning a response
func (fn AdminCreatePostHandlerFunc) Handle(params AdminCreatePostParams, principal *authz.Identity) middleware.Responder {
	return fn(params, principal)
}

// AdminCreatePostHandler interface for that can handle valid admin create post params
type AdminCreatePostHandler interface {
	Handle(AdminCreatePostParams, *authz.Identity) middleware.Responder
}

// NewAdminCreatePost creates a new http.Handler for the admin create post operation
func NewAdminCreatePost(ctx *middleware.Context, handler AdminCreatePostHandler) *AdminCreatePost {
	return &AdminCreatePost{Context: ctx, Handler: handler}
}

/*AdminCreatePost swagger:route POST /admin/posts admin adminCreatePost

Create Post

Создать пост для публикации.

*/
type AdminCreatePost struct {
	Context *middleware.Context
	Handler AdminCreatePostHandler
}

func (o *AdminCreatePost) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewAdminCreatePostParams()

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