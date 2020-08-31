package cli

import (
	"strings"
	"time"
)

type Config struct {
	Env string `default:"local" split_words:"true"`

	Domain      string `default:"localhost:8000"`
	DomainProto string `default:"http" split_words:"true"`

	Database             string `default:"postgres://birzzha:birzzha@localhost/birzzha?sslmode=disable"`
	DatabaseMaxOpenConns int    `default:"10" split_words:"true"`
	DatabaseMaxIdleConns int    `default:"0" split_words:"true"`

	Redis             string `default:"redis://localhost:6379"`
	RedisMaxOpenConns int    `default:"10" split_words:"true"`
	RedisMaxIdleConns int    `default:"0" split_words:"true"`

	TokenSecret   string        `required:"true" split_words:"true"`
	TokenLifeTime time.Duration `default:"24h" split_words:"true"`

	S3Endpoint  string `required:"true" split_words:"true"`
	S3AccessKey string `required:"true" split_words:"true"`
	S3SecretKey string `required:"true" split_words:"true"`
	S3Bucket    string `required:"true" split_words:"true"`
	S3Secure    bool   `required:"true" split_words:"true"`
	S3GlobalDir string `default:"" split_words:"true"`

	S3PublicPrefix string `required:"true" split_words:"true"`

	SentryDSN     string `default:"" split_words:"true"`
	SentryEnabled bool   `default:"false" split_words:"true"`

	BotToken              string        `required:"true" split_words:"true"`
	BotWebhookDomain      string        `required:"true" split_words:"true"`
	BotWebhookPath        string        `default:"/" split_words:"true"`
	BotWidgetAuthLifeTime time.Duration `default:"1m" split_words:"true"`

	InterkassaCheckoutID    string `split_words:"true"`
	InterkassaSecretKey     string `split_words:"true"`
	InterkassaTestSecretKey string ` split_words:"true"`

	UnitPayPublicKey string `split_words:"true"`
	UnitPaySecretKey string `split_words:"true"`

	Site string `default:"https://dev.birzzha.me/" split_words:"true"`

	SiteViewExpiration time.Duration `default:"24h" split_words:"true"`

	SitePathSellChannel string `default:"/channels/sell" split_words:"true"`
	SitePathListChannel string `default:"/channels" split_words:"true"`

	SitePathPaymentSuccess string `split_words:"true" default:"/payment/success"`
	SitePathPaymentFailed  string `split_words:"true" default:"/payment/failed"`
	SitePathPaymentPending string `split_words:"true" default:"/payment/pending"`

	FileProxyCachePath string `default:".cache" split_words:"true"`

	AdminNotificationsChannelID int64 `required:"true" split_words:"true"`

	Addr string `default:":8000"`

	// Proxy used for Yandex Metrika and Telemetr
	Proxy string `split_words:"true"`

	YandexMetrikaCounterID int `required:"true" split_words:"true"`

	// Run each 15 minutes
	WorkerUpdateLandingCron string `default:"0,15,30,45 * * * *" split_words:"true"`

	WorkerPublishPostsCron  string `default:"*/1 * * * *" split_words:"true"`
	WorkerUpdateLotListCron string `default:"0 0 * * *" split_words:"true"`

	Timezone string `default:"Europe/Kiev"`
}

func (cfg Config) getSiteFullPath(path string) string {
	site := strings.TrimPrefix(cfg.Site, "/")
	path = strings.TrimPrefix(path, "/")
	return strings.Join([]string{site, path}, "/")
}
