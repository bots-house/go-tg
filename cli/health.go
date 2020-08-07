package cli

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"net/http"

	"github.com/bots-house/birzzha/api/gen/models"
	"github.com/bots-house/birzzha/pkg/log"
	"github.com/pkg/errors"
)

func getHealthcheckEmoji(v bool) string {
	if v {
		return "✅"
	} else {
		return "💩"
	}
}

func runHealthcheck(ctx context.Context, cfg Config) error {
	host, port, err := net.SplitHostPort(cfg.Addr)
	if err != nil {
		return errors.Wrap(err, "parse host and port")
	}

	if host == "" {
		host = "localhost"
	}

	u := fmt.Sprintf("http://%s:%s/v1/health", host, port)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
	if err != nil {
		return errors.Wrap(err, "new request")
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return errors.Wrap(err, "do request")
	}
	defer res.Body.Close()

	var result models.Health

	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return errors.Wrap(err, "decode body")
	}

	log.Info(ctx, "check",
		"service", "postgres",
		"result", getHealthcheckEmoji(*result.Postgres.Ok),
		"took", result.Postgres.Took,
		"err", result.Postgres.Err,
	)

	log.Info(ctx, "check",
		"service", "redis",
		"result", getHealthcheckEmoji(*result.Redis.Ok),
		"took", result.Redis.Took,
		"err", result.Redis.Err,
	)

	if *result.Postgres.Ok && *result.Redis.Ok {
		return nil
	}

	return errors.New("service unavailable")
}
