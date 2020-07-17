package cli

import (
	"context"
	"database/sql"
	"fmt"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/benbjohnson/clock"
	tgbotapi "github.com/bots-house/telegram-bot-api"
	"github.com/kelseyhightower/envconfig"

	"github.com/bots-house/birzzha/api"
	"github.com/bots-house/birzzha/bot"
	"github.com/bots-house/birzzha/pkg/log"
	"github.com/bots-house/birzzha/pkg/storage"
	"github.com/bots-house/birzzha/pkg/tg"
	"github.com/bots-house/birzzha/service/admin"
	"github.com/bots-house/birzzha/service/auth"
	"github.com/bots-house/birzzha/service/catalog"
	"github.com/bots-house/birzzha/service/landing"
	"github.com/bots-house/birzzha/service/payment"
	"github.com/bots-house/birzzha/service/payment/interkassa"
	"github.com/bots-house/birzzha/service/personal"
	"github.com/bots-house/birzzha/store/postgres"

	"github.com/pkg/errors"
)

var logger = log.NewLogger(true, true)

const envPrefix = "BRZ"

func Run(ctx context.Context) {

	ctx = log.WithLogger(ctx, logger)
	if err := run(ctx); err != nil {
		log.Error(ctx, "fatal error", "err", err)
		os.Exit(1)
	}
}

func newStorage(cfg Config) (*storage.Space, error) {
	s3Config := &aws.Config{
		Credentials: credentials.NewStaticCredentials(cfg.S3AccessKey, cfg.S3SecretKey, ""),
		Endpoint:    aws.String(cfg.S3Endpoint),
		Region:      aws.String("us-east-1"),
	}

	newSession, err := session.NewSession(s3Config)
	if err != nil {
		return nil, err
	}
	s3Client := s3.New(newSession)

	return &storage.Space{
		Client:       s3Client,
		Bucket:       cfg.S3Bucket,
		PublicPrefix: cfg.S3PublicPrefix,
	}, nil
}

func newGatewayRegistry(ctx context.Context, cfg Config) *payment.GatewayRegistry {
	var gateways []payment.Gateway

	if cfg.InterkassaCheckoutID != "" {
		gateways = append(gateways, &interkassa.Gateway{
			CheckoutID:      cfg.InterkassaCheckoutID,
			SecretKey:       cfg.InterkassaSecretKey,
			TestSecretKey:   cfg.InterkassaTestSecretKey,
			NotificationURL: fmt.Sprintf("https://%s/v1/webhooks/gateways/interkassa", cfg.BotWebhookDomain),
			SuccessURL:      cfg.getSiteFullPath(cfg.SitePathPaymentSuccess),
			PendingURL:      cfg.getSiteFullPath(cfg.SitePathPaymentPending),
			FailedURL:       cfg.getSiteFullPath(cfg.SitePathPaymentFailed),
		})

		if cfg.InterkassaTestSecretKey != "" {
			log.Warn(ctx, "interkassa was configured in test mode")
		}
	} else {
		log.Warn(ctx, "interkassa gateway is not configured")
	}

	return payment.NewGatewayRegistry(gateways...)
}

func run(ctx context.Context) error {
	// load .env
	// if err := godotenv.Load(".env.local"); err != nil {
	// 	return errors.Wrap(err, "load env")
	// }

	// parse config
	var cfg Config

	if err := envconfig.Process(envPrefix, &cfg); err != nil {
		_ = envconfig.Usage(envPrefix, &cfg)
		return errors.Wrap(err, "parse config from env")
	}

	log.Info(ctx, "open db", "dsn", cfg.Database)

	// open and ping db
	db, err := sql.Open("postgres", cfg.Database)
	if err != nil {
		return errors.Wrap(err, "open db")
	}
	defer db.Close()

	log.Debug(ctx, "ping database")
	if err := db.PingContext(ctx); err != nil {
		return errors.Wrap(err, "ping db")
	}

	db.SetMaxOpenConns(cfg.DatabaseMaxOpenConns)
	db.SetMaxIdleConns(cfg.DatabaseMaxIdleConns)

	// create abstraction around db and apply migrations
	pg := postgres.NewPostgres(db)

	log.Info(ctx, "migrate database")
	if err := pg.Migrator().Up(ctx); err != nil {
		return errors.Wrap(err, "migrate db")
	}

	strg, err := newStorage(cfg)
	if err != nil {
		return errors.Wrap(err, "init file storage")
	}

	tgClient, err := tgbotapi.NewBotAPI(cfg.BotToken)
	if err != nil {
		return errors.Wrap(err, "new bot api client")
	}

	tgFileProxy, err := tg.NewFileProxy(cfg.FileProxyCachePath, tgClient)
	if err != nil {
		return errors.Wrap(err, "init file proxy cache")
	}

	var notifications *admin.Notifications

	if cfg.AdminNotificationsChannelID != 0 {
		notifications = admin.NewNotifications(tgClient, cfg.AdminNotificationsChannelID)
		defer notifications.Close()
	}

	authSrv := &auth.Service{
		UserStore: pg.User,
		Config: auth.Config{
			BotToken:           cfg.BotToken,
			WidgetInfoLifeTime: cfg.BotWidgetAuthLifeTime,

			TokenSecret:   cfg.TokenSecret,
			TokenLifeTime: cfg.TokenLifeTime,
		},
		Clock:         clock.New(),
		Storage:       strg,
		Bot:           tgClient,
		Notifications: notifications,
	}

	resolver := &tg.BotResolver{
		Client: tgClient,
		PublicURL: func(id string) string {
			return fmt.Sprintf("%s://%s/v1/tg/file/%s", cfg.DomainProto, cfg.Domain, id)
		},
	}

	resolverCache := tg.NewResolverCache(resolver, time.Minute*30)

	catalogSrv := &catalog.Service{
		Topic:    pg.Topic,
		Lot:      pg.Lot,
		LotTopic: pg.LotTopic,
		Resolver: resolverCache,
		Storage:  strg,
		Txier:    pg.Tx,
		User:     pg.User,
		Favorite: pg.Favorite,
		LotFile:  pg.LotFile,
	}

	adminSrv := &admin.Service{
		Review:            pg.Review,
		Settings:          pg.Settings,
		LotTopic:          pg.Topic,
		User:              pg.User,
		Lot:               pg.Lot,
		LotFile:           pg.LotFile,
		LotCanceledReason: pg.LotCanceledReason,
		Storage:           strg,
		AvatarResolver: tg.AvatarResolver{
			Client: &http.Client{},
		},
	}

	landingSrv := &landing.Service{
		Review: pg.Review,
	}

	bot := bot.New(bot.Config{
		Site:            cfg.Site,
		PathSellChannel: cfg.SitePathSellChannel,
		PathListChannel: cfg.SitePathListChannel,
	}, tgClient, authSrv, adminSrv)

	if err := bot.SetWebhookIfNeed(ctx, cfg.BotWebhookDomain, cfg.BotWebhookPath); err != nil {
		return errors.Wrap(err, "set bot webhook")
	}

	gateways := newGatewayRegistry(ctx, cfg)

	personalSrv := &personal.Service{
		Lot:               pg.Lot,
		Resolver:          resolver,
		Payment:           pg.Payment,
		Txier:             pg.Tx,
		Storage:           strg,
		Settings:          pg.Settings,
		Gateways:          gateways,
		AdminNotify:       notifications,
		LotCanceledReason: pg.LotCanceledReason,
		LotFile:           pg.LotFile,
	}

	handler := api.Handler{
		Auth:         authSrv,
		Admin:        adminSrv,
		Bot:          bot,
		BotFileProxy: tgFileProxy,
		Personal:     personalSrv,
		Catalog:      catalogSrv,
		Storage:      strg,
		Gateways:     gateways,
		Landing:      landingSrv,
		Logger:       log.Logger(ctx),
	}

	server := newServer(cfg, handler.Make())

	go func() {
		<-ctx.Done()

		shutdownCtx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		log.Info(ctx, "shutdown server")
		if err := server.Shutdown(shutdownCtx); err != nil {
			log.Warn(ctx, "shutdown error", "err", err)
		}
	}()

	log.Info(ctx, "start server", "addr", cfg.Addr, "webhook_domain", cfg.BotWebhookDomain)
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		return errors.Wrap(err, "listen and serve")
	}

	return nil
}

func newServer(cfg Config, handler http.Handler) *http.Server {
	baseCtx := context.Background()
	baseCtx = log.WithLogger(baseCtx, logger)

	return &http.Server{
		Addr:    cfg.Addr,
		Handler: handler,
		BaseContext: func(_ net.Listener) context.Context {
			return baseCtx
		},
	}
}
