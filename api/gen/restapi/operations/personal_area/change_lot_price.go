// Code generated by go-swagger; DO NOT EDIT.

package personal_area

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/bots-house/birzzha/api/authz"
)

// ChangeLotPriceHandlerFunc turns a function with the right signature into a change lot price handler
type ChangeLotPriceHandlerFunc func(ChangeLotPriceParams, *authz.Identity) middleware.Responder

// Handle executing the request and returning a response
func (fn ChangeLotPriceHandlerFunc) Handle(params ChangeLotPriceParams, principal *authz.Identity) middleware.Responder {
	return fn(params, principal)
}

// ChangeLotPriceHandler interface for that can handle valid change lot price params
type ChangeLotPriceHandler interface {
	Handle(ChangeLotPriceParams, *authz.Identity) middleware.Responder
}

// NewChangeLotPrice creates a new http.Handler for the change lot price operation
func NewChangeLotPrice(ctx *middleware.Context, handler ChangeLotPriceHandler) *ChangeLotPrice {
	return &ChangeLotPrice{Context: ctx, Handler: handler}
}

/*ChangeLotPrice swagger:route PUT /lots/{id}/change-price personal-area changeLotPrice

Change Lot Price

Изменяет цену лота если он не опубликованный, в ином случае возвращает ошибку.
Возможные ошибки:
  - `cant_change_lot_price_free` - Невозможно обновить цену лота бесплатно так как лот уже опубликован;


*/
type ChangeLotPrice struct {
	Context *middleware.Context
	Handler ChangeLotPriceHandler
}

func (o *ChangeLotPrice) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewChangeLotPriceParams()

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