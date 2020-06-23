package cli

import "time"

type Config struct {
	Database             string `default:"postgres://birzzha:birzzha@localhost/birzzha?sslmode=disable"`
	DatabaseMaxOpenConns int    `default:"10" split_words:"true"`
	DatabaseMaxIdleConns int    `default:"0" split_words:"true"`

	TokenSecret   string        `required:"true" split_words:"true"`
	TokenLifeTime time.Duration `default:"24h" split_words:"true"`

	S3Endpoint  string `required:"true" split_words:"true"`
	S3AccessKey string `required:"true" split_words:"true"`
	S3SecretKey string `required:"true" split_words:"true"`
	S3Bucket    string `required:"true" split_words:"true"`
	S3Secure    bool   `required:"true" split_words:"true"`

	S3PublicPrefix string `required:"true" split_words:"true"`

	BotToken              string        `required:"true" split_words:"true"`
	BotWebhookDomain      string        `required:"true" split_words:"true"`
	BotWebhookPath        string        `default:"/" split_words:"true"`
	BotWidgetAuthLifeTime time.Duration `default:"1m" split_words:"true"`

	Addr string `default:":8000"`
}
