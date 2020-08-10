package rate

import (
	"net/http"

	"github.com/go-redis/redis/v7"
	"github.com/pkg/errors"

	"github.com/bots-house/birzzha/pkg/log"
	"github.com/go-openapi/runtime/middleware"
	"github.com/ulule/limiter/v3/drivers/middleware/stdlib"

	"github.com/ulule/limiter/v3"
	redisstore "github.com/ulule/limiter/v3/drivers/store/redis"
)

const (
	reachedLimitErr = `{"code": "reached_limit", "description": "Вы достигли лимита по этому запросу. Попробуйте повторить попытку позже"}`
	internalErr     = `{"code": "internal_server_error", "description": "went wrong"}`
)

type Limitter struct {
	client redis.UniversalClient
}

func NewLimitter(redisHost string, maxIdleConns int) (*Limitter, error) {
	rds, err := newRedisV7(redisHost, maxIdleConns)
	if err != nil {
		return nil, errors.Wrap(err, "init limiter")
	}

	return &Limitter{client: rds}, nil
}

func (l Limitter) Default(suffix string) middleware.Builder {
	store, _ := redisstore.NewStoreWithOptions(l.client, limiter.StoreOptions{Prefix: limiter.DefaultPrefix + ":" + suffix})
	rate, _ := limiter.NewRateFromFormatted("20-M")

	return stdlib.NewMiddleware(limiter.New(store, rate, limiter.WithTrustForwardHeader(true)),
		stdlib.WithLimitReachedHandler(reachedLimitError),
		stdlib.WithErrorHandler(internalLimitError)).
		Handler
}

func reachedLimitError(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "appilication/json")
	w.WriteHeader(http.StatusTooManyRequests)
	_, _ = w.Write([]byte(reachedLimitErr))
}

func internalLimitError(w http.ResponseWriter, r *http.Request, err error) {
	log.Error(r.Context(), "limiter internal error", "error", err)
	w.Header().Set("Content-Type", "appilication/json")
	w.WriteHeader(http.StatusInternalServerError)
	_, _ = w.Write([]byte(internalErr))
}

func newRedisV7(host string, maxIdleConns int) (redis.UniversalClient, error) {
	opts, err := redis.ParseURL(host)
	if err != nil {
		return nil, errors.Wrap(err, "parse url (client v7)")
	}

	// remove heroku fake username
	opts.Username = ""

	opts.PoolSize = maxIdleConns

	rds := redis.NewClient(opts)

	_, err = rds.Ping().Result()
	if err != nil {
		return nil, errors.Wrap(err, "ping db (client v7)")
	}

	return rds, nil
}
