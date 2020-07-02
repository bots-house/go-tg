package api

import (
	"net/http"

	"github.com/bots-house/birzzha/api/authz"
	"github.com/bots-house/birzzha/api/gen/restapi"
	"github.com/bots-house/birzzha/api/gen/restapi/operations"
	authops "github.com/bots-house/birzzha/api/gen/restapi/operations/auth"
	catalogops "github.com/bots-house/birzzha/api/gen/restapi/operations/catalog"
	personalops "github.com/bots-house/birzzha/api/gen/restapi/operations/personal_area"
	webhookops "github.com/bots-house/birzzha/api/gen/restapi/operations/webhook"
	"github.com/bots-house/birzzha/pkg/storage"
	"github.com/bots-house/birzzha/pkg/tg"
	"github.com/bots-house/birzzha/service/landing"
	"github.com/bots-house/birzzha/service/payment"
	"github.com/bots-house/birzzha/service/personal"

	botops "github.com/bots-house/birzzha/api/gen/restapi/operations/bot"
	landingops "github.com/bots-house/birzzha/api/gen/restapi/operations/landing"
	kitlog "github.com/go-kit/kit/log"

	"github.com/go-http-utils/etag"

	"github.com/bots-house/birzzha/bot"
	"github.com/bots-house/birzzha/service/auth"
	"github.com/bots-house/birzzha/service/catalog"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
)

type Handler struct {
	Auth         *auth.Service
	Catalog      *catalog.Service
	Personal     *personal.Service
	Bot          *bot.Bot
	BotFileProxy *tg.FileProxy
	Storage      storage.Storage
	Gateways     *payment.GatewayRegistry
	Landing      *landing.Service
	Logger       kitlog.Logger
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
	// auth
	api.AuthCreateTokenHandler = authops.CreateTokenHandlerFunc(h.createToken)
	api.AuthGetUserHandler = authops.GetUserHandlerFunc(h.getUser)
	api.AuthLoginViaBotHandler = authops.LoginViaBotHandlerFunc(h.loginViaBot)

	// bot
	api.BotHandleUpdateHandler = botops.HandleUpdateHandlerFunc(h.handleBotUpdate)
	api.BotGetBotInfoHandler = botops.GetBotInfoHandlerFunc(h.getBotInfo)
	//api.BotGetFileHandler = botops.GetFileHandlerFunc(h.getBotFile)

	// catalog
	api.CatalogGetTopicsHandler = catalogops.GetTopicsHandlerFunc(h.getTopics)
	api.CatalogResolveTelegramHandler = catalogops.ResolveTelegramHandlerFunc(h.resolveTelegram)
	api.CatalogGetFilterBoundariesHandler = catalogops.GetFilterBoundariesHandlerFunc(h.getFilterBoundaries)
	api.CatalogGetLotsHandler = catalogops.GetLotsHandlerFunc(h.getLots)
	api.CatalogGetLotHandler = catalogops.GetLotHandlerFunc(h.getLot)

	// landing
	api.LandingGetReviewsHandler = landingops.GetReviewsHandlerFunc(h.getReviews)

	// personal
	api.PersonalAreaCreateLotHandler = personalops.CreateLotHandlerFunc(h.createLot)
	api.PersonalAreaGetApplicationInoviceHandler = personalops.GetApplicationInoviceHandlerFunc(h.getApplicationInvoice)
	api.PersonalAreaCreateApplicationPaymentHandler = personalops.CreateApplicationPaymentHandlerFunc(h.createApplicationPayment)
	api.PersonalAreaGetPaymentStatusHandler = personalops.GetPaymentStatusHandlerFunc(h.getPaymentStatus)
	api.PersonalAreaGetUserLotsHandler = personalops.GetUserLotsHandlerFunc(h.getOwnedLots)

	// webhook
	api.WebhookHandleGatewayNotificationHandler = webhookops.HandleGatewayNotificationHandlerFunc(h.handleGatewayNotification)
}

func (h Handler) setupMiddleware(api *operations.BirzzhaAPI) {
	etagMiddleware := middleware.Builder(func(handler http.Handler) http.Handler {
		return etag.Handler(handler, false)
	})

	api.AddMiddlewareFor(http.MethodGet, "/bot", etagMiddleware)
}

func (h Handler) wrapMiddleware(handler http.Handler) http.Handler {
	// handler = common.WrapMiddlewareFS(handler, h.Service.Config.MediaStoragePath)
	handler = h.wrapMiddlewareLogger(handler)

	handler = h.wrapMiddlewareRecovery(handler)
	fileProxyWrapper := newFileProxyWrapper(h.BotFileProxy)

	handler = fileProxyWrapper(handler)

	return handler
}

func (h Handler) Make() http.Handler {
	api := h.newAPI()

	h.setupProducersAndConsumers(api)
	h.setupHandlers(api)
	h.setupMiddleware(api)
	h.setupAuth(api)

	return h.wrapMiddleware(api.Serve(nil))
}
