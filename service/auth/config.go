package auth

import (
	"crypto/sha256"
	"sync"
	"time"
)

type Config struct {
	BotToken           string
	WidgetInfoLifeTime time.Duration
	TokenSecret        string
	TokenLifeTime      time.Duration

	botTokenHash []byte
	botTokenOnce sync.Once
}

func (cfg *Config) getBotTokenHash() []byte {
	cfg.botTokenOnce.Do(func() {
		hsh := sha256.New()
		hsh.Write([]byte(cfg.BotToken))
		cfg.botTokenHash = hsh.Sum(nil)
	})

	return cfg.botTokenHash
}
