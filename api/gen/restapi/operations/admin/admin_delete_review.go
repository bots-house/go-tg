// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/bots-house/birzzha/api/authz"
)

// AdminDeleteReviewHandlerFunc turns a function with the right signature into a admin delete review handler
type AdminDeleteReviewHandlerFunc func(AdminDeleteReviewParams, *authz.Identity) middleware.Responder

// Handle executing the request and returning a response
func (fn AdminDeleteReviewHandlerFunc) Handle(params AdminDeleteReviewParams, principal *authz.Identity) middleware.Responder {
	return fn(params, principal)
}

// AdminDeleteReviewHandler interface for that can handle valid admin delete review params
type AdminDeleteReviewHandler interface {
	Handle(AdminDeleteReviewParams, *authz.Identity) middleware.Responder
}

// NewAdminDeleteReview creates a new http.Handler for the admin delete review operation
func NewAdminDeleteReview(ctx *middleware.Context, handler AdminDeleteReviewHandler) *AdminDeleteReview {
	return &AdminDeleteReview{Context: ctx, Handler: handler}
}

/*AdminDeleteReview swagger:route DELETE /admin/reviews/{id} admin adminDeleteReview

Delete Review

Удалить отзыв.

*/
type AdminDeleteReview struct {
	Context *middleware.Context
	Handler AdminDeleteReviewHandler
}

func (o *AdminDeleteReview) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewAdminDeleteReviewParams()

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