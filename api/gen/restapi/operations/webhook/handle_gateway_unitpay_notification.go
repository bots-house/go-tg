// Code generated by go-swagger; DO NOT EDIT.

package webhook

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// HandleGatewayUnitpayNotificationHandlerFunc turns a function with the right signature into a handle gateway unitpay notification handler
type HandleGatewayUnitpayNotificationHandlerFunc func(HandleGatewayUnitpayNotificationParams) middleware.Responder

// Handle executing the request and returning a response
func (fn HandleGatewayUnitpayNotificationHandlerFunc) Handle(params HandleGatewayUnitpayNotificationParams) middleware.Responder {
	return fn(params)
}

// HandleGatewayUnitpayNotificationHandler interface for that can handle valid handle gateway unitpay notification params
type HandleGatewayUnitpayNotificationHandler interface {
	Handle(HandleGatewayUnitpayNotificationParams) middleware.Responder
}

// NewHandleGatewayUnitpayNotification creates a new http.Handler for the handle gateway unitpay notification operation
func NewHandleGatewayUnitpayNotification(ctx *middleware.Context, handler HandleGatewayUnitpayNotificationHandler) *HandleGatewayUnitpayNotification {
	return &HandleGatewayUnitpayNotification{Context: ctx, Handler: handler}
}

/*HandleGatewayUnitpayNotification swagger:route GET /webhooks/gateways/{name} webhook handleGatewayUnitpayNotification

Handle Gateway Notification

Обработка уведомления о изменении состояния платeжа (unitpay)


*/
type HandleGatewayUnitpayNotification struct {
	Context *middleware.Context
	Handler HandleGatewayUnitpayNotificationHandler
}

func (o *HandleGatewayUnitpayNotification) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewHandleGatewayUnitpayNotificationParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
