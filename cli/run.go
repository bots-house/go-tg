package cli

import (
	"context"
	"database/sql"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/benbjohnson/clock"
	"github.com/bots-house/birzzha/api"
	"github.com/bots-house/birzzha/bot"
	"github.com/bots-house/birzzha/pkg/log"
	"github.com/bots-house/birzzha/pkg/storage"
	"github.com/bots-house/birzzha/service/auth"
	"github.com/bots-house/birzzha/service/catalog"
	"github.com/bots-house/birzzha/store/postgres"
	tgbotapi "github.com/bots-house/telegram-bot-api"
	"github.com/kelseyhightower/envconfig"

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

func run(ctx context.Context) error {
	// parse config
	var cfg Config

	if err := envconfig.Process(envPrefix, &cfg); err != nil {
		envconfig.Usage(envPrefix, &cfg)
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

	strg := &storage.FS{
		Path: "storage/",
	}

	tgClient, err := tgbotapi.NewBotAPI(cfg.BotToken)
	if err != nil {
		return errors.Wrap(err, "new bot api client")
	}

	authSrv := &auth.Service{
		UserStore: pg.User,
		Config: auth.Config{
			BotToken:           cfg.BotToken,
			WidgetInfoLifeTime: cfg.BotWidgetAuthLifeTime,

			TokenSecret:   cfg.TokenSecret,
			TokenLifeTime: cfg.TokenLifeTime,
		},
		Clock:   clock.New(),
		Storage: strg,
		Bot:     tgClient,
	}

	catalogSrv := &catalog.Service{
		Topic: pg.Topic,
	}

	bot := bot.New(tgClient, authSrv)

	if err := bot.SetWebhookIfNeed(ctx, cfg.BotWebhookDomain, cfg.BotWebhookPath); err != nil {
		return errors.Wrap(err, "set bot webhook")
	}

	handler := api.Handler{
		Auth:    authSrv,
		Bot:     bot,
		Catalog: catalogSrv,
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
