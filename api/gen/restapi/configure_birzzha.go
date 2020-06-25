// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/bots-house/birzzha/api/authz"
	"github.com/bots-house/birzzha/api/gen/restapi/operations"
	"github.com/bots-house/birzzha/api/gen/restapi/operations/auth"
	"github.com/bots-house/birzzha/api/gen/restapi/operations/bot"
	"github.com/bots-house/birzzha/api/gen/restapi/operations/catalog"
)

//go:generate swagger generate server --target ../../gen --name Birzzha --spec ../../../../../../../var/folders/0k/708dty_x6c1411whczf7pxvh0000gn/T/swagger.yml433442978 --principal authz.Identity --exclude-main

func configureFlags(api *operations.BirzzhaAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.BirzzhaAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	// Applies when the "X-Token" header is set
	if api.TokenHeaderAuth == nil {
		api.TokenHeaderAuth = func(token string) (*authz.Identity, error) {
			return nil, errors.NotImplemented("api key auth (TokenHeader) X-Token from header param [X-Token] has not yet been implemented")
		}
	}
	// Applies when the "token" query is set
	if api.TokenQueryAuth == nil {
		api.TokenQueryAuth = func(token string) (*authz.Identity, error) {
			return nil, errors.NotImplemented("api key auth (TokenQuery) token from query param [token] has not yet been implemented")
		}
	}

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()
	if api.CatalogCreateLotHandler == nil {
		api.CatalogCreateLotHandler = catalog.CreateLotHandlerFunc(func(params catalog.CreateLotParams, principal *authz.Identity) middleware.Responder {
			return middleware.NotImplemented("operation catalog.CreateLot has not yet been implemented")
		})
	}
	if api.AuthCreateTokenHandler == nil {
		api.AuthCreateTokenHandler = auth.CreateTokenHandlerFunc(func(params auth.CreateTokenParams) middleware.Responder {
			return middleware.NotImplemented("operation auth.CreateToken has not yet been implemented")
		})
	}
	if api.BotGetBotInfoHandler == nil {
		api.BotGetBotInfoHandler = bot.GetBotInfoHandlerFunc(func(params bot.GetBotInfoParams) middleware.Responder {
			return middleware.NotImplemented("operation bot.GetBotInfo has not yet been implemented")
		})
	}
	if api.CatalogGetTopicsHandler == nil {
		api.CatalogGetTopicsHandler = catalog.GetTopicsHandlerFunc(func(params catalog.GetTopicsParams) middleware.Responder {
			return middleware.NotImplemented("operation catalog.GetTopics has not yet been implemented")
		})
	}
	if api.AuthGetUserHandler == nil {
		api.AuthGetUserHandler = auth.GetUserHandlerFunc(func(params auth.GetUserParams, principal *authz.Identity) middleware.Responder {
			return middleware.NotImplemented("operation auth.GetUser has not yet been implemented")
		})
	}
	if api.BotHandleUpdateHandler == nil {
		api.BotHandleUpdateHandler = bot.HandleUpdateHandlerFunc(func(params bot.HandleUpdateParams) middleware.Responder {
			return middleware.NotImplemented("operation bot.HandleUpdate has not yet been implemented")
		})
	}
	if api.AuthLoginViaBotHandler == nil {
		api.AuthLoginViaBotHandler = auth.LoginViaBotHandlerFunc(func(params auth.LoginViaBotParams) middleware.Responder {
			return middleware.NotImplemented("operation auth.LoginViaBot has not yet been implemented")
		})
	}
	if api.CatalogResolveTelegramHandler == nil {
		api.CatalogResolveTelegramHandler = catalog.ResolveTelegramHandlerFunc(func(params catalog.ResolveTelegramParams) middleware.Responder {
			return middleware.NotImplemented("operation catalog.ResolveTelegram has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
