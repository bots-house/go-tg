// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/runtime/security"
	"github.com/go-openapi/spec"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/bots-house/birzzha/api/authz"
	"github.com/bots-house/birzzha/api/gen/restapi/operations/auth"
	"github.com/bots-house/birzzha/api/gen/restapi/operations/bot"
	"github.com/bots-house/birzzha/api/gen/restapi/operations/catalog"
)

// NewBirzzhaAPI creates a new Birzzha instance
func NewBirzzhaAPI(spec *loads.Document) *BirzzhaAPI {
	return &BirzzhaAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		PreServerShutdown:   func() {},
		ServerShutdown:      func() {},
		spec:                spec,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,

		JSONConsumer: runtime.JSONConsumer(),

		JSONProducer: runtime.JSONProducer(),

		CatalogCreateLotHandler: catalog.CreateLotHandlerFunc(func(params catalog.CreateLotParams, principal *authz.Identity) middleware.Responder {
			return middleware.NotImplemented("operation catalog.CreateLot has not yet been implemented")
		}),
		AuthCreateTokenHandler: auth.CreateTokenHandlerFunc(func(params auth.CreateTokenParams) middleware.Responder {
			return middleware.NotImplemented("operation auth.CreateToken has not yet been implemented")
		}),
		BotGetBotInfoHandler: bot.GetBotInfoHandlerFunc(func(params bot.GetBotInfoParams) middleware.Responder {
			return middleware.NotImplemented("operation bot.GetBotInfo has not yet been implemented")
		}),
		CatalogGetTopicsHandler: catalog.GetTopicsHandlerFunc(func(params catalog.GetTopicsParams) middleware.Responder {
			return middleware.NotImplemented("operation catalog.GetTopics has not yet been implemented")
		}),
		AuthGetUserHandler: auth.GetUserHandlerFunc(func(params auth.GetUserParams, principal *authz.Identity) middleware.Responder {
			return middleware.NotImplemented("operation auth.GetUser has not yet been implemented")
		}),
		BotHandleUpdateHandler: bot.HandleUpdateHandlerFunc(func(params bot.HandleUpdateParams) middleware.Responder {
			return middleware.NotImplemented("operation bot.HandleUpdate has not yet been implemented")
		}),
		AuthLoginViaBotHandler: auth.LoginViaBotHandlerFunc(func(params auth.LoginViaBotParams) middleware.Responder {
			return middleware.NotImplemented("operation auth.LoginViaBot has not yet been implemented")
		}),
		CatalogResolveTelegramHandler: catalog.ResolveTelegramHandlerFunc(func(params catalog.ResolveTelegramParams) middleware.Responder {
			return middleware.NotImplemented("operation catalog.ResolveTelegram has not yet been implemented")
		}),

		// Applies when the "X-Token" header is set
		TokenHeaderAuth: func(token string) (*authz.Identity, error) {
			return nil, errors.NotImplemented("api key auth (TokenHeader) X-Token from header param [X-Token] has not yet been implemented")
		},
		// Applies when the "token" query is set
		TokenQueryAuth: func(token string) (*authz.Identity, error) {
			return nil, errors.NotImplemented("api key auth (TokenQuery) token from query param [token] has not yet been implemented")
		},
		// default authorizer is authorized meaning no requests are blocked
		APIAuthorizer: security.Authorized(),
	}
}

/*BirzzhaAPI the birzzha API */
type BirzzhaAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator
	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator
	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for the following mime types:
	//   - application/json
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for the following mime types:
	//   - application/json
	JSONProducer runtime.Producer

	// TokenHeaderAuth registers a function that takes a token and returns a principal
	// it performs authentication based on an api key X-Token provided in the header
	TokenHeaderAuth func(string) (*authz.Identity, error)

	// TokenQueryAuth registers a function that takes a token and returns a principal
	// it performs authentication based on an api key token provided in the query
	TokenQueryAuth func(string) (*authz.Identity, error)

	// APIAuthorizer provides access control (ACL/RBAC/ABAC) by providing access to the request and authenticated principal
	APIAuthorizer runtime.Authorizer

	// CatalogCreateLotHandler sets the operation handler for the create lot operation
	CatalogCreateLotHandler catalog.CreateLotHandler
	// AuthCreateTokenHandler sets the operation handler for the create token operation
	AuthCreateTokenHandler auth.CreateTokenHandler
	// BotGetBotInfoHandler sets the operation handler for the get bot info operation
	BotGetBotInfoHandler bot.GetBotInfoHandler
	// CatalogGetTopicsHandler sets the operation handler for the get topics operation
	CatalogGetTopicsHandler catalog.GetTopicsHandler
	// AuthGetUserHandler sets the operation handler for the get user operation
	AuthGetUserHandler auth.GetUserHandler
	// BotHandleUpdateHandler sets the operation handler for the handle update operation
	BotHandleUpdateHandler bot.HandleUpdateHandler
	// AuthLoginViaBotHandler sets the operation handler for the login via bot operation
	AuthLoginViaBotHandler auth.LoginViaBotHandler
	// CatalogResolveTelegramHandler sets the operation handler for the resolve telegram operation
	CatalogResolveTelegramHandler catalog.ResolveTelegramHandler
	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// PreServerShutdown is called before the HTTP(S) server is shutdown
	// This allows for custom functions to get executed before the HTTP(S) server stops accepting traffic
	PreServerShutdown func()

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// SetDefaultProduces sets the default produces media type
func (o *BirzzhaAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *BirzzhaAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *BirzzhaAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *BirzzhaAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *BirzzhaAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *BirzzhaAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *BirzzhaAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the BirzzhaAPI
func (o *BirzzhaAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.TokenHeaderAuth == nil {
		unregistered = append(unregistered, "XTokenAuth")
	}
	if o.TokenQueryAuth == nil {
		unregistered = append(unregistered, "TokenAuth")
	}

	if o.CatalogCreateLotHandler == nil {
		unregistered = append(unregistered, "catalog.CreateLotHandler")
	}
	if o.AuthCreateTokenHandler == nil {
		unregistered = append(unregistered, "auth.CreateTokenHandler")
	}
	if o.BotGetBotInfoHandler == nil {
		unregistered = append(unregistered, "bot.GetBotInfoHandler")
	}
	if o.CatalogGetTopicsHandler == nil {
		unregistered = append(unregistered, "catalog.GetTopicsHandler")
	}
	if o.AuthGetUserHandler == nil {
		unregistered = append(unregistered, "auth.GetUserHandler")
	}
	if o.BotHandleUpdateHandler == nil {
		unregistered = append(unregistered, "bot.HandleUpdateHandler")
	}
	if o.AuthLoginViaBotHandler == nil {
		unregistered = append(unregistered, "auth.LoginViaBotHandler")
	}
	if o.CatalogResolveTelegramHandler == nil {
		unregistered = append(unregistered, "catalog.ResolveTelegramHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *BirzzhaAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *BirzzhaAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {
	result := make(map[string]runtime.Authenticator)
	for name := range schemes {
		switch name {
		case "TokenHeader":
			scheme := schemes[name]
			result[name] = o.APIKeyAuthenticator(scheme.Name, scheme.In, func(token string) (interface{}, error) {
				return o.TokenHeaderAuth(token)
			})

		case "TokenQuery":
			scheme := schemes[name]
			result[name] = o.APIKeyAuthenticator(scheme.Name, scheme.In, func(token string) (interface{}, error) {
				return o.TokenQueryAuth(token)
			})

		}
	}
	return result
}

// Authorizer returns the registered authorizer
func (o *BirzzhaAPI) Authorizer() runtime.Authorizer {
	return o.APIAuthorizer
}

// ConsumersFor gets the consumers for the specified media types.
// MIME type parameters are ignored here.
func (o *BirzzhaAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {
	result := make(map[string]runtime.Consumer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONConsumer
		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result
}

// ProducersFor gets the producers for the specified media types.
// MIME type parameters are ignored here.
func (o *BirzzhaAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {
	result := make(map[string]runtime.Producer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONProducer
		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result
}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *BirzzhaAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the birzzha API
func (o *BirzzhaAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *BirzzhaAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened
	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/lot"] = catalog.NewCreateLot(o.context, o.CatalogCreateLotHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/auth"] = auth.NewCreateToken(o.context, o.AuthCreateTokenHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/bot"] = bot.NewGetBotInfo(o.context, o.BotGetBotInfoHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/topics"] = catalog.NewGetTopics(o.context, o.CatalogGetTopicsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/user"] = auth.NewGetUser(o.context, o.AuthGetUserHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/bot"] = bot.NewHandleUpdate(o.context, o.BotHandleUpdateHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/auth/bot"] = auth.NewLoginViaBot(o.context, o.AuthLoginViaBotHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/tg/resolve"] = catalog.NewResolveTelegram(o.context, o.CatalogResolveTelegramHandler)
}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *BirzzhaAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *BirzzhaAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *BirzzhaAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *BirzzhaAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}

// AddMiddlewareFor adds a http middleware to existing handler
func (o *BirzzhaAPI) AddMiddlewareFor(method, path string, builder middleware.Builder) {
	um := strings.ToUpper(method)
	if path == "/" {
		path = ""
	}
	o.Init()
	if h, ok := o.handlers[um][path]; ok {
		o.handlers[method][path] = builder(h)
	}
}
