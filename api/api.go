package api

import (
	"net/http"

	"github.com/bots-house/birzzha/api/authz"
	"github.com/bots-house/birzzha/api/gen/restapi"
	"github.com/bots-house/birzzha/api/gen/restapi/operations"
	"github.com/bots-house/birzzha/api/gen/restapi/operations/auth"
	"github.com/bots-house/birzzha/bot"
	authsrv "github.com/bots-house/birzzha/service/auth"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
)

type Handler struct {
	Auth *authsrv.Service
	Bot  *bot.Bot
}

func (h Handler) newAPI() *operations.BirzzhaAPI {
	spec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		panic("load spec failed: " + err.Error())
	}

	return operations.NewBirzzhaAPI(spec)
}

func (h Handler) setupProducersAndConsumers(api *operations.BirzzhaAPI) {
	// set JSON producer/consumer
	api.JSONProducer = runtime.JSONProducer()
	api.JSONConsumer = runtime.JSONConsumer()
}

func (h Handler) setupAuth(api *operations.BirzzhaAPI) {
	api.TokenHeaderAuth = authz.Parse
	api.TokenQueryAuth = authz.Parse

	api.APIAuthorizer = &authz.Authorizer{
		Srv: h.Auth,
	}
}

func (h Handler) setupHandlers(api *operations.BirzzhaAPI) {
	api.AuthCreateTokenHandler = auth.CreateTokenHandlerFunc(h.createToken)
	api.AuthGetUserHandler = auth.GetUserHandlerFunc(h.getUser)

	api.HandleBotUpdateHandler = operations.HandleBotUpdateHandlerFunc(h.handleBotUpdate)
}

func (h Handler) wrapMiddleware(handler http.Handler) http.Handler {
	// handler = common.WrapMiddlewareRecovery(handler)
	// handler = common.WrapMiddlewareFS(handler, h.Service.Config.MediaStoragePath)

	return handler
}

func (h Handler) Make() http.Handler {
	api := h.newAPI()

	h.setupProducersAndConsumers(api)
	h.setupHandlers(api)
	h.setupAuth(api)

	return h.wrapMiddleware(api.Serve(nil))
}
