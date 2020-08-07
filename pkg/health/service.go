package health

import (
	"context"
	"database/sql"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
)

// Check result of single service
type Check struct {
	Err  error
	Took time.Duration
}

// Ok returns true if no error
func (check *Check) Ok() bool {
	return check.Err == nil
}

// Health of subservice healthcheck
type Health struct {
	Redis    Check
	Postgres Check
}

// Ok returns true, if all deps is alive.
func (health *Health) Ok() bool {
	return health.Redis.Ok() && health.Postgres.Ok()
}

// Service check if all dependencies is alive
type Service struct {
	Postgres *sql.DB
	Redis    redis.UniversalClient
	Timeout  time.Duration
}

func (srv *Service) getTimeout() time.Duration {
	if srv.Timeout == 0 {
		return time.Second * 5
	}

	return srv.Timeout
}

// Check if deps is alive
func (srv *Service) Check(ctx context.Context) *Health {
	ctx, cancel := context.WithTimeout(ctx, srv.getTimeout())
	defer cancel()

	result := &Health{}

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func(start time.Time) {
		defer wg.Done()

		result.Postgres.Err = srv.Postgres.PingContext(ctx)
		result.Postgres.Took = time.Since(start)
	}(time.Now())

	wg.Add(1)
	go func(start time.Time) {
		defer wg.Done()

		result.Redis.Err = srv.Redis.Ping(ctx).Err()
		result.Redis.Took = time.Since(start)
	}(time.Now())

	wg.Wait()

	return result
}
