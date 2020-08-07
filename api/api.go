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

	"github.com/bots-house/birzzha/pkg/health"
	"github.com/bots-house/birzzha/pkg/storage"
	"github.com/bots-house/birzzha/pkg/tg"
	"github.com/bots-house/birzzha/service/admin"
	"github.com/bots-house/birzzha/service/landing"
	"github.com/bots-house/birzzha/service/payment"
	"github.com/bots-house/birzzha/service/personal"
	"github.com/bots-house/birzzha/service/views"

	adminops "github.com/bots-house/birzzha/api/gen/restapi/operations/admin"
	botops "github.com/bots-house/birzzha/api/gen/restapi/operations/bot"
	healthops "github.com/bots-house/birzzha/api/gen/restapi/operations/health"

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
	Admin        *admin.Service
	Catalog      *catalog.Service
	Personal     *personal.Service
	Bot          *bot.Bot
	BotFileProxy *tg.FileProxy
	Storage      storage.Storage
	Gateways     *payment.GatewayRegistry
	Landing      *landing.Service
	Views        *views.Service
	Logger       kitlog.Logger
	Health       *health.Service
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
	api.CatalogGetDailyCoverageHandler = catalogops.GetDailyCoverageHandlerFunc(h.getDailyCoverage)
	api.CatalogGetTopicsHandler = catalogops.GetTopicsHandlerFunc(h.getTopics)
	api.CatalogResolveTelegramHandler = catalogops.ResolveTelegramHandlerFunc(h.resolveTelegram)
	api.CatalogGetFilterBoundariesHandler = catalogops.GetFilterBoundariesHandlerFunc(h.getFilterBoundaries)
	api.CatalogGetLotsHandler = catalogops.GetLotsHandlerFunc(h.getLots)
	api.CatalogGetLotHandler = catalogops.GetLotHandlerFunc(h.getLot)
	api.CatalogGetSimilarLotsHandler = catalogops.GetSimilarLotsHandlerFunc(h.getSimilarLots)
	api.CatalogToggleLotFavoriteHandler = catalogops.ToggleLotFavoriteHandlerFunc(h.toggleLotFavorite)

	// landing
	api.LandingGetReviewsHandler = landingops.GetReviewsHandlerFunc(h.getReviews)
	api.LandingGetLandingHandler = landingops.GetLandingHandlerFunc(h.getLanding)

	// personal
	api.PersonalAreaCreateLotHandler = personalops.CreateLotHandlerFunc(h.createLot)
	api.PersonalAreaGetApplicationInoviceHandler = personalops.GetApplicationInoviceHandlerFunc(h.getApplicationInvoice)
	api.PersonalAreaCreateApplicationPaymentHandler = personalops.CreateApplicationPaymentHandlerFunc(h.createApplicationPayment)
	api.PersonalAreaGetPaymentStatusHandler = personalops.GetPaymentStatusHandlerFunc(h.getPaymentStatus)
	api.PersonalAreaGetUserLotsHandler = personalops.GetUserLotsHandlerFunc(h.getOwnedLots)
	api.PersonalAreaGetLotCanceledReasonsHandler = personalops.GetLotCanceledReasonsHandlerFunc(h.getLotCanceledReasons)
	api.PersonalAreaCancelLotHandler = personalops.CancelLotHandlerFunc(h.cancelLot)
	api.PersonalAreaUploadLotFileHandler = personalops.UploadLotFileHandlerFunc(h.uploadLotFile)
	api.PersonalAreaGetFavoriteLotsHandler = personalops.GetFavoriteLotsHandlerFunc(h.getFavoriteLots)
	api.PersonalAreaChangeLotPriceHandler = personalops.ChangeLotPriceHandlerFunc(h.changeLotPrice)
	api.PersonalAreaGetChangePriceInvoiceHandler = personalops.GetChangePriceInvoiceHandlerFunc(h.getChangePriceInvoice)
	api.PersonalAreaCreateChangePricePaymentHandler = personalops.CreateChangePricePaymentHandlerFunc(h.createChangePricePayment)

	// webhook
	api.WebhookHandleGatewayNotificationHandler = webhookops.HandleGatewayNotificationHandlerFunc(h.handleGatewayNotification)

	//admin
	api.AdminAdminDeleteReviewHandler = adminops.AdminDeleteReviewHandlerFunc(h.adminDeleteReview)
	api.AdminAdminUpdateReviewHandler = adminops.AdminUpdateReviewHandlerFunc(h.adminUpdateReview)
	api.AdminAdminGetReviewsHandler = adminops.AdminGetReviewsHandlerFunc(h.adminGetReviews)

	api.AdminAdminGetUsersHandler = adminops.AdminGetUsersHandlerFunc(h.adminGetUsers)
	api.AdminToggleUserAdminHandler = adminops.ToggleUserAdminHandlerFunc(h.toggleUserAdmin)
	api.AdminAdminGetLotStatusesHandler = adminops.AdminGetLotStatusesHandlerFunc(h.adminGetLotStatuses)
	api.AdminAdminGetLotsHandler = adminops.AdminGetLotsHandlerFunc(h.adminGetLots)
	api.AdminAdminGetPostTextHandler = adminops.AdminGetPostTextHandlerFunc(h.adminGetPostText)
	api.AdminAdminSendPostPreviewHandler = adminops.AdminSendPostPreviewHandlerFunc(h.adminSendPostPreview)
	api.AdminAdminUpdateLotHandler = adminops.AdminUpdateLotHandlerFunc(h.adminUpdateLot)
	api.AdminAdminDeclineLotHandler = adminops.AdminDeclineLotHandlerFunc(h.adminDeclineLot)
	api.AdminAdminGetLotHandler = adminops.AdminGetLotHandlerFunc(h.adminGetLot)

	api.AdminAdminGetSettingsHandler = adminops.AdminGetSettingsHandlerFunc(h.adminGetSettings)
	api.AdminAdminUpdateSettingsLandingHandler = adminops.AdminUpdateSettingsLandingHandlerFunc(h.adminUpdateSettingsLanding)
	api.AdminAdminUpdateSettingsPricesHandler = adminops.AdminUpdateSettingsPricesHandlerFunc(h.adminUpdateSettingsPrices)
	api.AdminAdminUpdateSettingsChannelHandler = adminops.AdminUpdateSettingsChannelHandlerFunc(h.adminUpdateSettingsChannel)

	api.AdminAdminCreateTopicHandler = adminops.AdminCreateTopicHandlerFunc(h.adminCreateTopic)
	api.AdminAdminUpdateTopicHandler = adminops.AdminUpdateTopicHandlerFunc(h.adminUpdateTopic)
	api.AdminAdminDeleteTopicHandler = adminops.AdminDeleteTopicHandlerFunc(h.adminDeleteTopic)

	api.AdminAdminCreateLotCanceledReasonHandler = adminops.AdminCreateLotCanceledReasonHandlerFunc(h.adminCreateLotCanceledReason)
	api.AdminAdminUpdateLotCanceledReasonHandler = adminops.AdminUpdateLotCanceledReasonHandlerFunc(h.adminUpdateLotCanceledReason)
	api.AdminAdminDeleteLotCanceledReasonHandler = adminops.AdminDeleteLotCanceledReasonHandlerFunc(h.adminDeleteLotCanceledReason)

	api.AdminAdminCreatePostHandler = adminops.AdminCreatePostHandlerFunc(h.adminCreatePost)

	// health
	api.HealthGetHealthHandler = healthops.GetHealthHandlerFunc(h.getHealth)
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
