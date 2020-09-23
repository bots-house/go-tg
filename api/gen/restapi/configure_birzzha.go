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
	"github.com/bots-house/birzzha/api/gen/restapi/operations/admin"
	"github.com/bots-house/birzzha/api/gen/restapi/operations/auth"
	"github.com/bots-house/birzzha/api/gen/restapi/operations/bot"
	"github.com/bots-house/birzzha/api/gen/restapi/operations/catalog"
	"github.com/bots-house/birzzha/api/gen/restapi/operations/health"
	"github.com/bots-house/birzzha/api/gen/restapi/operations/landing"
	"github.com/bots-house/birzzha/api/gen/restapi/operations/personal_area"
	"github.com/bots-house/birzzha/api/gen/restapi/operations/webhook"
)

//go:generate swagger generate server --target ../../gen --name Birzzha --spec ../../../../../../../tmp/swagger.yml079597412 --principal authz.Identity --exclude-main

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
	api.MultipartformConsumer = runtime.DiscardConsumer
	api.UrlformConsumer = runtime.DiscardConsumer

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
	if api.AdminAdminCancelLotHandler == nil {
		api.AdminAdminCancelLotHandler = admin.AdminCancelLotHandlerFunc(func(params admin.AdminCancelLotParams, principal *authz.Identity) middleware.Responder {
			return middleware.NotImplemented("operation admin.AdminCancelLot has not yet been implemented")
		})
	}
	if api.AdminAdminCreateCouponHandler == nil {
		api.AdminAdminCreateCouponHandler = admin.AdminCreateCouponHandlerFunc(func(params admin.AdminCreateCouponParams, principal *authz.Identity) middleware.Responder {
			return middleware.NotImplemented("operation admin.AdminCreateCoupon has not yet been implemented")
		})
	}
	if api.AdminAdminCreateLotCanceledReasonHandler == nil {
		api.AdminAdminCreateLotCanceledReasonHandler = admin.AdminCreateLotCanceledReasonHandlerFunc(func(params admin.AdminCreateLotCanceledReasonParams, principal *authz.Identity) middleware.Responder {
			return middleware.NotImplemented("operation admin.AdminCreateLotCanceledReason has not yet been implemented")
		})
	}
	if api.AdminAdminCreatePostHandler == nil {
		api.AdminAdminCreatePostHandler = admin.AdminCreatePostHandlerFunc(func(params admin.AdminCreatePostParams, principal *authz.Identity) middleware.Responder {
			return middleware.NotImplemented("operation admin.AdminCreatePost has not yet been implemented")
		})
	}
	if api.AdminAdminCreateTopicHandler == nil {
		api.AdminAdminCreateTopicHandler = admin.AdminCreateTopicHandlerFunc(func(params admin.AdminCreateTopicParams, principal *authz.Identity) middleware.Responder {
			return middleware.NotImplemented("operation admin.AdminCreateTopic has not yet been implemented")
		})
	}
	if api.AdminAdminDeclineLotHandler == nil {
		api.AdminAdminDeclineLotHandler = admin.AdminDeclineLotHandlerFunc(func(params admin.AdminDeclineLotParams, principal *authz.Identity) middleware.Responder {
			return middleware.NotImplemented("operation admin.AdminDeclineLot has not yet been implemented")
		})
	}
	if api.AdminAdminDeleteCouponHandler == nil {
		api.AdminAdminDeleteCouponHandler = admin.AdminDeleteCouponHandlerFunc(func(params admin.AdminDeleteCouponParams, principal *authz.Identity) middleware.Responder {
			return middleware.NotImplemented("operation admin.AdminDeleteCoupon has not yet been implemented")
		})
	}
	if api.AdminAdminDeleteLotCanceledReasonHandler == nil {
		api.AdminAdminDeleteLotCanceledReasonHandler = admin.AdminDeleteLotCanceledReasonHandlerFunc(func(params admin.AdminDeleteLotCanceledReasonParams, principal *authz.Identity) middleware.Responder {
			return middleware.NotImplemented("operation admin.AdminDeleteLotCanceledReason has not yet been implemented")
		})
	}
	if api.AdminAdminDeletePostHandler == nil {
		api.AdminAdminDeletePostHandler = admin.AdminDeletePostHandlerFunc(func(params admin.AdminDeletePostParams, principal *authz.Identity) middleware.Responder {
			return middleware.NotImplemented("operation admin.AdminDeletePost has not yet been implemented")
		})
	}
	if api.AdminAdminDeleteReviewHandler == nil {
		api.AdminAdminDeleteReviewHandler = admin.AdminDeleteReviewHandlerFunc(func(params admin.AdminDeleteReviewParams, principal *authz.Identity) middleware.Responder {
			return middleware.NotImplemented("operation admin.AdminDeleteReview has not yet been implemented")
		})
	}
	if api.AdminAdminDeleteTopicHandler == nil {
		api.AdminAdminDeleteTopicHandler = admin.AdminDeleteTopicHandlerFunc(func(params admin.AdminDeleteTopicParams, principal *authz.Identity) middleware.Responder {
			return middleware.NotImplemented("operation admin.AdminDeleteTopic has not yet been implemented")
		})
	}
	if api.AdminAdminGetCouponsHandler == nil {
		api.AdminAdminGetCouponsHandler = admin.AdminGetCouponsHandlerFunc(func(params admin.AdminGetCouponsParams, principal *authz.Identity) middleware.Responder {
			return middleware.NotImplemented("operation admin.AdminGetCoupons has not yet been implemented")
		})
	}
	if api.AdminAdminGetLotHandler == nil {
		api.AdminAdminGetLotHandler = admin.AdminGetLotHandlerFunc(func(params admin.AdminGetLotParams, principal *authz.Identity) middleware.Responder {
			return middleware.NotImplemented("operation admin.AdminGetLot has not yet been implemented")
		})
	}
	if api.AdminAdminGetLotStatusesHandler == nil {
		api.AdminAdminGetLotStatusesHandler = admin.AdminGetLotStatusesHandlerFunc(func(params admin.AdminGetLotStatusesParams, principal *authz.Identity) middleware.Responder {
			return middleware.NotImplemented("operation admin.AdminGetLotStatuses has not yet been implemented")
		})
	}
	if api.AdminAdminGetLotsHandler == nil {
		api.AdminAdminGetLotsHandler = admin.AdminGetLotsHandlerFunc(func(params admin.AdminGetLotsParams, principal *authz.Identity) middleware.Responder {
			return middleware.NotImplemented("operation admin.AdminGetLots has not yet been implemented")
		})
	}
	if api.AdminAdminGetPostTextHandler == nil {
		api.AdminAdminGetPostTextHandler = admin.AdminGetPostTextHandlerFunc(func(params admin.AdminGetPostTextParams, principal *authz.Identity) middleware.Responder {
			return middleware.NotImplemented("operation admin.AdminGetPostText has not yet been implemented")
		})
	}
	if api.AdminAdminGetPostsHandler == nil {
		api.AdminAdminGetPostsHandler = admin.AdminGetPostsHandlerFunc(func(params admin.AdminGetPostsParams, principal *authz.Identity) middleware.Responder {
			return middleware.NotImplemented("operation admin.AdminGetPosts has not yet been implemented")
		})
	}
	if api.AdminAdminGetReviewsHandler == nil {
		api.AdminAdminGetReviewsHandler = admin.AdminGetReviewsHandlerFunc(func(params admin.AdminGetReviewsParams, principal *authz.Identity) middleware.Responder {
			return middleware.NotImplemented("operation admin.AdminGetReviews has not yet been implemented")
		})
	}
	if api.AdminAdminGetSettingsHandler == nil {
		api.AdminAdminGetSettingsHandler = admin.AdminGetSettingsHandlerFunc(func(params admin.AdminGetSettingsParams, principal *authz.Identity) middleware.Responder {
			return middleware.NotImplemented("operation admin.AdminGetSettings has not yet been implemented")
		})
	}
	if api.AdminAdminGetUsersHandler == nil {
		api.AdminAdminGetUsersHandler = admin.AdminGetUsersHandlerFunc(func(params admin.AdminGetUsersParams, principal *authz.Identity) middleware.Responder {
			return middleware.NotImplemented("operation admin.AdminGetUsers has not yet been implemented")
		})
	}
	if api.AdminAdminRefreshLotHandler == nil {
		api.AdminAdminRefreshLotHandler = admin.AdminRefreshLotHandlerFunc(func(params admin.AdminRefreshLotParams, principal *authz.Identity) middleware.Responder {
			return middleware.NotImplemented("operation admin.AdminRefreshLot has not yet been implemented")
		})
	}
	if api.AdminAdminSendPostPreviewHandler == nil {
		api.AdminAdminSendPostPreviewHandler = admin.AdminSendPostPreviewHandlerFunc(func(params admin.AdminSendPostPreviewParams, principal *authz.Identity) middleware.Responder {
			return middleware.NotImplemented("operation admin.AdminSendPostPreview has not yet been implemented")
		})
	}
	if api.AdminAdminUpdateCouponHandler == nil {
		api.AdminAdminUpdateCouponHandler = admin.AdminUpdateCouponHandlerFunc(func(params admin.AdminUpdateCouponParams, principal *authz.Identity) middleware.Responder {
			return middleware.NotImplemented("operation admin.AdminUpdateCoupon has not yet been implemented")
		})
	}
	if api.AdminAdminUpdateLotHandler == nil {
		api.AdminAdminUpdateLotHandler = admin.AdminUpdateLotHandlerFunc(func(params admin.AdminUpdateLotParams, principal *authz.Identity) middleware.Responder {
			return middleware.NotImplemented("operation admin.AdminUpdateLot has not yet been implemented")
		})
	}
	if api.AdminAdminUpdateLotCanceledReasonHandler == nil {
		api.AdminAdminUpdateLotCanceledReasonHandler = admin.AdminUpdateLotCanceledReasonHandlerFunc(func(params admin.AdminUpdateLotCanceledReasonParams, principal *authz.Identity) middleware.Responder {
			return middleware.NotImplemented("operation admin.AdminUpdateLotCanceledReason has not yet been implemented")
		})
	}
	if api.AdminAdminUpdatePostHandler == nil {
		api.AdminAdminUpdatePostHandler = admin.AdminUpdatePostHandlerFunc(func(params admin.AdminUpdatePostParams, principal *authz.Identity) middleware.Responder {
			return middleware.NotImplemented("operation admin.AdminUpdatePost has not yet been implemented")
		})
	}
	if api.AdminAdminUpdateReviewHandler == nil {
		api.AdminAdminUpdateReviewHandler = admin.AdminUpdateReviewHandlerFunc(func(params admin.AdminUpdateReviewParams, principal *authz.Identity) middleware.Responder {
			return middleware.NotImplemented("operation admin.AdminUpdateReview has not yet been implemented")
		})
	}
	if api.AdminAdminUpdateSettingsChannelHandler == nil {
		api.AdminAdminUpdateSettingsChannelHandler = admin.AdminUpdateSettingsChannelHandlerFunc(func(params admin.AdminUpdateSettingsChannelParams, principal *authz.Identity) middleware.Responder {
			return middleware.NotImplemented("operation admin.AdminUpdateSettingsChannel has not yet been implemented")
		})
	}
	if api.AdminAdminUpdateSettingsGarantHandler == nil {
		api.AdminAdminUpdateSettingsGarantHandler = admin.AdminUpdateSettingsGarantHandlerFunc(func(params admin.AdminUpdateSettingsGarantParams, principal *authz.Identity) middleware.Responder {
			return middleware.NotImplemented("operation admin.AdminUpdateSettingsGarant has not yet been implemented")
		})
	}
	if api.AdminAdminUpdateSettingsLandingHandler == nil {
		api.AdminAdminUpdateSettingsLandingHandler = admin.AdminUpdateSettingsLandingHandlerFunc(func(params admin.AdminUpdateSettingsLandingParams, principal *authz.Identity) middleware.Responder {
			return middleware.NotImplemented("operation admin.AdminUpdateSettingsLanding has not yet been implemented")
		})
	}
	if api.AdminAdminUpdateSettingsPricesHandler == nil {
		api.AdminAdminUpdateSettingsPricesHandler = admin.AdminUpdateSettingsPricesHandlerFunc(func(params admin.AdminUpdateSettingsPricesParams, principal *authz.Identity) middleware.Responder {
			return middleware.NotImplemented("operation admin.AdminUpdateSettingsPrices has not yet been implemented")
		})
	}
	if api.AdminAdminUpdateTopicHandler == nil {
		api.AdminAdminUpdateTopicHandler = admin.AdminUpdateTopicHandlerFunc(func(params admin.AdminUpdateTopicParams, principal *authz.Identity) middleware.Responder {
			return middleware.NotImplemented("operation admin.AdminUpdateTopic has not yet been implemented")
		})
	}
	if api.PersonalAreaCancelLotHandler == nil {
		api.PersonalAreaCancelLotHandler = personal_area.CancelLotHandlerFunc(func(params personal_area.CancelLotParams, principal *authz.Identity) middleware.Responder {
			return middleware.NotImplemented("operation personal_area.CancelLot has not yet been implemented")
		})
	}
	if api.PersonalAreaChangeLotPriceHandler == nil {
		api.PersonalAreaChangeLotPriceHandler = personal_area.ChangeLotPriceHandlerFunc(func(params personal_area.ChangeLotPriceParams, principal *authz.Identity) middleware.Responder {
			return middleware.NotImplemented("operation personal_area.ChangeLotPrice has not yet been implemented")
		})
	}
	if api.PersonalAreaCreateApplicationPaymentHandler == nil {
		api.PersonalAreaCreateApplicationPaymentHandler = personal_area.CreateApplicationPaymentHandlerFunc(func(params personal_area.CreateApplicationPaymentParams, principal *authz.Identity) middleware.Responder {
			return middleware.NotImplemented("operation personal_area.CreateApplicationPayment has not yet been implemented")
		})
	}
	if api.PersonalAreaCreateChangePricePaymentHandler == nil {
		api.PersonalAreaCreateChangePricePaymentHandler = personal_area.CreateChangePricePaymentHandlerFunc(func(params personal_area.CreateChangePricePaymentParams, principal *authz.Identity) middleware.Responder {
			return middleware.NotImplemented("operation personal_area.CreateChangePricePayment has not yet been implemented")
		})
	}
	if api.PersonalAreaCreateLotHandler == nil {
		api.PersonalAreaCreateLotHandler = personal_area.CreateLotHandlerFunc(func(params personal_area.CreateLotParams, principal *authz.Identity) middleware.Responder {
			return middleware.NotImplemented("operation personal_area.CreateLot has not yet been implemented")
		})
	}
	if api.AuthCreateTokenHandler == nil {
		api.AuthCreateTokenHandler = auth.CreateTokenHandlerFunc(func(params auth.CreateTokenParams) middleware.Responder {
			return middleware.NotImplemented("operation auth.CreateToken has not yet been implemented")
		})
	}
	if api.PersonalAreaGetApplicationInoviceHandler == nil {
		api.PersonalAreaGetApplicationInoviceHandler = personal_area.GetApplicationInoviceHandlerFunc(func(params personal_area.GetApplicationInoviceParams, principal *authz.Identity) middleware.Responder {
			return middleware.NotImplemented("operation personal_area.GetApplicationInovice has not yet been implemented")
		})
	}
	if api.BotGetBotInfoHandler == nil {
		api.BotGetBotInfoHandler = bot.GetBotInfoHandlerFunc(func(params bot.GetBotInfoParams) middleware.Responder {
			return middleware.NotImplemented("operation bot.GetBotInfo has not yet been implemented")
		})
	}
	if api.PersonalAreaGetChangePriceInvoiceHandler == nil {
		api.PersonalAreaGetChangePriceInvoiceHandler = personal_area.GetChangePriceInvoiceHandlerFunc(func(params personal_area.GetChangePriceInvoiceParams, principal *authz.Identity) middleware.Responder {
			return middleware.NotImplemented("operation personal_area.GetChangePriceInvoice has not yet been implemented")
		})
	}
	if api.PersonalAreaGetCouponHandler == nil {
		api.PersonalAreaGetCouponHandler = personal_area.GetCouponHandlerFunc(func(params personal_area.GetCouponParams, principal *authz.Identity) middleware.Responder {
			return middleware.NotImplemented("operation personal_area.GetCoupon has not yet been implemented")
		})
	}
	if api.CatalogGetDailyCoverageHandler == nil {
		api.CatalogGetDailyCoverageHandler = catalog.GetDailyCoverageHandlerFunc(func(params catalog.GetDailyCoverageParams) middleware.Responder {
			return middleware.NotImplemented("operation catalog.GetDailyCoverage has not yet been implemented")
		})
	}
	if api.PersonalAreaGetFavoriteLotsHandler == nil {
		api.PersonalAreaGetFavoriteLotsHandler = personal_area.GetFavoriteLotsHandlerFunc(func(params personal_area.GetFavoriteLotsParams, principal *authz.Identity) middleware.Responder {
			return middleware.NotImplemented("operation personal_area.GetFavoriteLots has not yet been implemented")
		})
	}
	if api.CatalogGetFilterBoundariesHandler == nil {
		api.CatalogGetFilterBoundariesHandler = catalog.GetFilterBoundariesHandlerFunc(func(params catalog.GetFilterBoundariesParams) middleware.Responder {
			return middleware.NotImplemented("operation catalog.GetFilterBoundaries has not yet been implemented")
		})
	}
	if api.HealthGetHealthHandler == nil {
		api.HealthGetHealthHandler = health.GetHealthHandlerFunc(func(params health.GetHealthParams) middleware.Responder {
			return middleware.NotImplemented("operation health.GetHealth has not yet been implemented")
		})
	}
	if api.LandingGetLandingHandler == nil {
		api.LandingGetLandingHandler = landing.GetLandingHandlerFunc(func(params landing.GetLandingParams) middleware.Responder {
			return middleware.NotImplemented("operation landing.GetLanding has not yet been implemented")
		})
	}
	if api.CatalogGetLotHandler == nil {
		api.CatalogGetLotHandler = catalog.GetLotHandlerFunc(func(params catalog.GetLotParams, principal *authz.Identity) middleware.Responder {
			return middleware.NotImplemented("operation catalog.GetLot has not yet been implemented")
		})
	}
	if api.PersonalAreaGetLotCanceledReasonsHandler == nil {
		api.PersonalAreaGetLotCanceledReasonsHandler = personal_area.GetLotCanceledReasonsHandlerFunc(func(params personal_area.GetLotCanceledReasonsParams) middleware.Responder {
			return middleware.NotImplemented("operation personal_area.GetLotCanceledReasons has not yet been implemented")
		})
	}
	if api.CatalogGetLotsHandler == nil {
		api.CatalogGetLotsHandler = catalog.GetLotsHandlerFunc(func(params catalog.GetLotsParams, principal *authz.Identity) middleware.Responder {
			return middleware.NotImplemented("operation catalog.GetLots has not yet been implemented")
		})
	}
	if api.PersonalAreaGetPaymentStatusHandler == nil {
		api.PersonalAreaGetPaymentStatusHandler = personal_area.GetPaymentStatusHandlerFunc(func(params personal_area.GetPaymentStatusParams, principal *authz.Identity) middleware.Responder {
			return middleware.NotImplemented("operation personal_area.GetPaymentStatus has not yet been implemented")
		})
	}
	if api.LandingGetReviewsHandler == nil {
		api.LandingGetReviewsHandler = landing.GetReviewsHandlerFunc(func(params landing.GetReviewsParams) middleware.Responder {
			return middleware.NotImplemented("operation landing.GetReviews has not yet been implemented")
		})
	}
	if api.CatalogGetSimilarLotsHandler == nil {
		api.CatalogGetSimilarLotsHandler = catalog.GetSimilarLotsHandlerFunc(func(params catalog.GetSimilarLotsParams) middleware.Responder {
			return middleware.NotImplemented("operation catalog.GetSimilarLots has not yet been implemented")
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
	if api.PersonalAreaGetUserLotsHandler == nil {
		api.PersonalAreaGetUserLotsHandler = personal_area.GetUserLotsHandlerFunc(func(params personal_area.GetUserLotsParams, principal *authz.Identity) middleware.Responder {
			return middleware.NotImplemented("operation personal_area.GetUserLots has not yet been implemented")
		})
	}
	if api.WebhookHandleGatewayNotificationHandler == nil {
		api.WebhookHandleGatewayNotificationHandler = webhook.HandleGatewayNotificationHandlerFunc(func(params webhook.HandleGatewayNotificationParams) middleware.Responder {
			return middleware.NotImplemented("operation webhook.HandleGatewayNotification has not yet been implemented")
		})
	}
	if api.WebhookHandleGatewayUnitpayNotificationHandler == nil {
		api.WebhookHandleGatewayUnitpayNotificationHandler = webhook.HandleGatewayUnitpayNotificationHandlerFunc(func(params webhook.HandleGatewayUnitpayNotificationParams) middleware.Responder {
			return middleware.NotImplemented("operation webhook.HandleGatewayUnitpayNotification has not yet been implemented")
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
	if api.CatalogToggleLotFavoriteHandler == nil {
		api.CatalogToggleLotFavoriteHandler = catalog.ToggleLotFavoriteHandlerFunc(func(params catalog.ToggleLotFavoriteParams, principal *authz.Identity) middleware.Responder {
			return middleware.NotImplemented("operation catalog.ToggleLotFavorite has not yet been implemented")
		})
	}
	if api.AdminToggleUserAdminHandler == nil {
		api.AdminToggleUserAdminHandler = admin.ToggleUserAdminHandlerFunc(func(params admin.ToggleUserAdminParams, principal *authz.Identity) middleware.Responder {
			return middleware.NotImplemented("operation admin.ToggleUserAdmin has not yet been implemented")
		})
	}
	if api.PersonalAreaUploadLotFileHandler == nil {
		api.PersonalAreaUploadLotFileHandler = personal_area.UploadLotFileHandlerFunc(func(params personal_area.UploadLotFileParams) middleware.Responder {
			return middleware.NotImplemented("operation personal_area.UploadLotFile has not yet been implemented")
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
