package kv

import (
	"context"
	"errors"
	"time"
)

var (
	ErrKeyNotFound = errors.New("key not found")
)

type Store interface {
	Set(ctx context.Context, key string, value interface{}, opts ...SetOption) error
	Get(ctx context.Context, key string, dst interface{}) error
	Sub(prefix string) Store
}

type options struct {
	TTL time.Duration
}

func newOptions(opts []SetOption) *options {
	result := &options{}

	for _, opt := range opts {
		opt(result)
	}

	return result
}

type SetOption func(opts *options)

func Expiration(v time.Duration) SetOption {
	return func(opts *options) {
		opts.TTL = v
	}
}
