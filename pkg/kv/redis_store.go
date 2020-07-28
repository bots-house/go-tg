package kv

import (
	"context"
	"encoding/json"

	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
)

type RedisStore struct {
	prefix string
	Client redis.UniversalClient
}

var _ Store = &RedisStore{}

func (store *RedisStore) getKey(key string) string {
	if store.prefix != "" {
		return store.prefix + "::" + key
	}

	return store.prefix
}

func (store *RedisStore) Set(ctx context.Context, key string, value interface{}, optss ...SetOption) error {
	opts := newOptions(optss)

	key = store.getKey(key)

	data, err := json.Marshal(value)
	if err != nil {
		return errors.Wrap(err, "marshal")
	}

	return store.Client.Set(ctx, key, data, opts.TTL).Err()
}

func (store *RedisStore) Get(ctx context.Context, key string, dst interface{}) error {
	key = store.getKey(key)

	data, err := store.Client.Get(ctx, key).Result()
	if err == redis.Nil {
		return ErrKeyNotFound
	} else if err != nil {
		return err
	}

	if dst != nil {
		if err := json.Unmarshal([]byte(data), dst); err != nil {
			return errors.Wrap(err, "unmarshal")
		}
	}

	return nil
}

func (store *RedisStore) Sub(prefix string) Store {
	return &RedisStore{
		prefix: prefix,
		Client: store.Client,
	}
}
