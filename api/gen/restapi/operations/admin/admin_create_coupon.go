// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/bots-house/birzzha/api/authz"
)

// AdminCreateCouponHandlerFunc turns a function with the right signature into a admin create coupon handler
type AdminCreateCouponHandlerFunc func(AdminCreateCouponParams, *authz.Identity) middleware.Responder

// Handle executing the request and returning a response
func (fn AdminCreateCouponHandlerFunc) Handle(params AdminCreateCouponParams, principal *authz.Identity) middleware.Responder {
	return fn(params, principal)
}

// AdminCreateCouponHandler interface for that can handle valid admin create coupon params
type AdminCreateCouponHandler interface {
	Handle(AdminCreateCouponParams, *authz.Identity) middleware.Responder
}

// NewAdminCreateCoupon creates a new http.Handler for the admin create coupon operation
func NewAdminCreateCoupon(ctx *middleware.Context, handler AdminCreateCouponHandler) *AdminCreateCoupon {
	return &AdminCreateCoupon{Context: ctx, Handler: handler}
}

/*AdminCreateCoupon swagger:route POST /admin/coupons admin adminCreateCoupon

Create Coupon

Создание купона.

Возможные ошибки:
  - `coupon_with_this_code_already_exist` - Купон с таким кодом уже существует;
  - `coupon_discount_must_be_greater_than_zero` - Скидка купона должна быть выше 0;
  - `coupon_discount_must_be_smaller` - Скидка купона должна быть ниже;



*/
type AdminCreateCoupon struct {
	Context *middleware.Context
	Handler AdminCreateCouponHandler
}

func (o *AdminCreateCoupon) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewAdminCreateCouponParams()

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
