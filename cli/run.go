package cli

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/benbjohnson/clock"
	tgbotapi "github.com/bots-house/telegram-bot-api"
	tgme "github.com/bots-house/tg-me"
	"github.com/go-redis/redis/v8"
	"github.com/kelseyhightower/envconfig"
	"github.com/subosito/gotenv"
	"golang.org/x/sync/errgroup"

	"github.com/bots-house/birzzha/api"
	"github.com/bots-house/birzzha/bot"
	"github.com/bots-house/birzzha/core"
	"github.com/bots-house/birzzha/pkg/kv"
	"github.com/bots-house/birzzha/pkg/log"
	"github.com/bots-house/birzzha/pkg/stat"
	"github.com/bots-house/birzzha/pkg/storage"
	"github.com/bots-house/birzzha/pkg/tg"
	"github.com/bots-house/birzzha/service/admin"
	"github.com/bots-house/birzzha/service/auth"
	"github.com/bots-house/birzzha/service/catalog"
	"github.com/bots-house/birzzha/service/landing"
	"github.com/bots-house/birzzha/service/payment"
	"github.com/bots-house/birzzha/service/payment/interkassa"
	"github.com/bots-house/birzzha/service/personal"
	"github.com/bots-house/birzzha/service/posting"
	"github.com/bots-house/birzzha/service/views"
	"github.com/bots-house/birzzha/store/postgres"
	"github.com/bots-house/birzzha/worker"

	"github.com/pkg/errors"
)

var logger = log.NewLogger(true, true)

const envPrefix = "BRZ"

func Run(ctx context.Context, revision string) {

	ctx = log.WithLogger(ctx, logger)

	log.Info(ctx, "start", "revision", revision)

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
		GlobalDir:    cfg.S3GlobalDir,
	}, nil
}

func newParser(client *http.Client) core.LotExtraResourceParser {
	return personal.LotExtraResourceParser{Telegram: &tgme.Parser{Client: client}}
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

func newRedis(ctx context.Context, cfg Config) (redis.UniversalClient, error) {
	opts, err := redis.ParseURL(cfg.Redis)
	if err != nil {
		return nil, errors.Wrap(err, "parse url")
	}

	// remove heroku fake username
	opts.Username = ""

	opts.PoolSize = cfg.RedisMaxIdleConns

	rds := redis.NewClient(opts)

	_, err = rds.Ping(ctx).Result()
	if err != nil {
		return nil, errors.Wrap(err, "ping db")
	}

	return rds, nil
}

func parseConfig(config string) (Config, error) {
	var cfg Config

	// load envs
	if config != "" {
		if err := gotenv.Load(config); err != nil {
			return cfg, errors.Wrap(err, "load env")
		}
	}

	if err := envconfig.Process(envPrefix, &cfg); err != nil {
		_ = envconfig.Usage(envPrefix, &cfg)
		return cfg, err
	}

	return cfg, nil
}

func run(ctx context.Context) error {

	// parse flags
	var (
		flagWorker, flagServer bool
		flagConfig             string
	)

	flag.BoolVar(&flagWorker, "worker", false, "run only server")
	flag.BoolVar(&flagServer, "server", false, "run only worker")
	flag.StringVar(&flagConfig, "config", "", "path to env file")

	flag.Parse()

	// parse config
	cfg, err := parseConfig(flagConfig)
	if err != nil {
		return errors.Wrap(err, "parse config")
	}

	log.Info(ctx, "open db", "dsn", cfg.Database)

	//create time zone
	timezone, err := time.LoadLocation(cfg.Timezone)
	if err != nil {
		return errors.Wrapf(err, "failed to load timezone=%s", timezone)
	}

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

	log.Info(ctx, "open redis", "dsn", cfg.Redis)
	rds, err := newRedis(ctx, cfg)
	if err != nil {
		return errors.Wrap(err, "open redis")
	}
	defer rds.Close()

	kvStore := &kv.RedisStore{
		Client: rds,
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

	botLinkBuilder := bot.NewLinkBuilder(tgClient.Self.UserName)

	authSrv := &auth.Service{
		UserStore: pg.User,
		Config: auth.Config{
			BotToken:           cfg.BotToken,
			WidgetInfoLifeTime: cfg.BotWidgetAuthLifeTime,

			TokenSecret:   cfg.TokenSecret,
			TokenLifeTime: cfg.TokenLifeTime,
		},
		Clock:          clock.New(),
		Storage:        strg,
		BotLinkBuilder: botLinkBuilder,
		Notifications:  notifications,
	}

	resolver := &tg.BotResolver{
		Client: tgClient,
		PublicURL: func(id string) string {
			return fmt.Sprintf("%s://%s/v1/tg/file/%s", cfg.DomainProto, cfg.Domain, id)
		},
	}

	resolverCache := tg.NewResolverCache(resolver, time.Minute*30)

	proxyDoer, err := newProxyDoer(ctx, cfg)
	if err != nil {
		return errors.Wrap(err, "new proxy doer")
	}
	telemetr := &stat.TelegramTelemetr{
		Doer: proxyDoer,
	}

	catalogSrv := &catalog.Service{
		Topic:        pg.Topic,
		Lot:          pg.Lot,
		LotTopic:     pg.LotTopic,
		Resolver:     resolverCache,
		Storage:      strg,
		Txier:        pg.Tx,
		User:         pg.User,
		Favorite:     pg.Favorite,
		LotFile:      pg.LotFile,
		TelegramStat: telemetr,
	}

	postingSrv := &posting.Service{
		Lot:      pg.Lot,
		Settings: pg.Settings,
		Topic:    pg.Topic,
		User:     pg.User,
		Post:     pg.Post,
		Txier:    pg.Tx,
		TgClient: tgClient,
	}

	adminSrv := &admin.Service{
		Review:            pg.Review,
		Settings:          pg.Settings,
		Topic:             pg.Topic,
		LotTopic:          pg.LotTopic,
		User:              pg.User,
		Lot:               pg.Lot,
		LotFile:           pg.LotFile,
		Txier:             pg.Tx,
		Favorite:          pg.Favorite,
		LotCanceledReason: pg.LotCanceledReason,
		Landing:           pg.Landing,
		Storage:           strg,
		BotLinkBuilder:    botLinkBuilder,
		AvatarResolver: tg.AvatarResolver{
			Client: http.DefaultClient,
		},
		Posting: postingSrv,
		Post:    pg.Post,
	}

	landingSrv := &landing.Service{
		Review:   pg.Review,
		Settings: pg.Settings,
		Resolver: resolverCache,
		Landing:  pg.Landing,
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
		LotFavorite:       pg.Favorite,
		Storage:           strg,
		Settings:          pg.Settings,
		TelegramStat:      telemetr,
		Gateways:          gateways,
		AdminNotify:       notifications,
		LotCanceledReason: pg.LotCanceledReason,
		LotFile:           pg.LotFile,
		Parser:            newParser(&http.Client{}),
	}

	viewsSrv := &views.Service{
		Lot:                pg.Lot,
		Txier:              pg.Tx,
		KV:                 kvStore.Sub("views"),
		SiteViewExpiration: cfg.SiteViewExpiration,
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
		Logger:       log.GetLogger(ctx),
		Views:        viewsSrv,
	}

	// setup server
	var (
		srv *http.Server
		wrk *worker.Worker
	)

	isRunBoth := !flagServer && !flagWorker

	if flagServer || isRunBoth {
		srv = newServer(cfg, handler.Make())
	}

	if flagWorker || isRunBoth {
		wrk = &worker.Worker{
			Landing:  pg.Landing,
			Lot:      pg.Lot,
			Settings: pg.Settings,

			Resolver: resolver,

			Storage: strg,

			SiteStat: &stat.SiteYandexMetrika{
				CounterID: cfg.YandexMetrikaCounterID,
				Doer:      proxyDoer,
			},

			TelegramStat: telemetr,

			Location: timezone,

			UpdateLandingSpec: cfg.WorkerUpdateLandingCron,
			PublishPostsSpec:  cfg.WorkerPublishPostsCron,
			UpdateLotsSpec:    cfg.WorkerUpdateLotListCron,
			Posting:           postingSrv,
		}
	}

	go func() {
		<-ctx.Done()
		log.Warn(ctx, "shutdown signal received")

		shutdownCtx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		if srv != nil {
			log.Info(ctx, "shutdown server")
			if err := srv.Shutdown(shutdownCtx); err != nil {
				log.Warn(ctx, "shutdown error", "err", err)
			}
		}
	}()

	g, ctx := errgroup.WithContext(ctx)

	if srv != nil {
		g.Go(func() error {
			log.Info(ctx, "start server", "addr", cfg.Addr, "webhook_domain", cfg.BotWebhookDomain)
			if err := srv.ListenAndServe(); err != http.ErrServerClosed {
				return errors.Wrap(err, "run server")
			}
			return nil
		})
	}

	if wrk != nil {
		g.Go(func() error {
			if err := wrk.Run(ctx); err != nil {
				return errors.Wrap(err, "run worker")
			}

			return nil
		})
	}

	return g.Wait()
}

func newServer(cfg Config, handler http.Handler) *http.Server {
	baseCtx := context.Background()
	baseCtx = log.WithLogger(baseCtx, logger)
	baseCtx = log.WithPrefix(baseCtx, "scope", "server")

	return &http.Server{
		Addr:    cfg.Addr,
		Handler: handler,
		BaseContext: func(_ net.Listener) context.Context {
			return baseCtx
		},
	}
}

func newProxyDoer(ctx context.Context, cfg Config) (*http.Client, error) {
	if cfg.Proxy != "" {
		log.Info(ctx, "use proxy for stats service", "dsn", cfg.Proxy)
		u, err := url.Parse(cfg.Proxy)
		if err != nil {
			return nil, errors.Wrap(err, "parse proxy url")
		}

		return &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyURL(u),
			},
		}, nil
	}

	return http.DefaultClient, nil
}
