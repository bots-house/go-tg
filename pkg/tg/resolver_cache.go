package tg

import (
	"context"
	"strconv"
	"time"

	"github.com/patrickmn/go-cache"
	"github.com/pkg/errors"
)

type ResolverCache struct {
	resolver Resolver
	cache    *cache.Cache
}

func NewResolverCache(resolver Resolver, exp time.Duration) *ResolverCache {
	return &ResolverCache{
		resolver: resolver,
		cache: cache.New(
			exp,
			1*time.Hour,
		),
	}
}

func (r *ResolverCache) get(typ queryType, val string) (*ResolveResult, bool) {
	result, ok := r.cache.Get(strconv.Itoa(int(typ)) + val)
	if ok {
		return result.(*ResolveResult), true
	}
	return nil, false

}

func (r *ResolverCache) getByID(id int64) (*ResolveResult, bool) {
	key := strconv.FormatInt(id, 10)

	result, ok := r.cache.Get(key)
	if ok {
		return result.(*ResolveResult), true
	}
	return nil, false
}

func (r *ResolverCache) save(typ queryType, val string, result *ResolveResult) {
	r.cache.SetDefault(strconv.Itoa(int(typ))+val, result)

	if c := result.Channel; c != nil {
		key := strconv.FormatInt(c.ID, 10)
		r.cache.SetDefault(key, result)
	}
}

func (r *ResolverCache) Resolve(ctx context.Context, query string) (*ResolveResult, error) {
	qt, val := parseResolveQuery(query)

	result, exists := r.get(qt, val)
	if !exists {
		var err error
		result, err = r.resolver.Resolve(ctx, query)
		if err != nil {
			return nil, errors.Wrap(err, "resolve")
		}
		r.save(qt, val, result)
	}

	return result, nil
}

func (r *ResolverCache) ResolveByID(ctx context.Context, id int64) (*ResolveResult, error) {
	result, exists := r.getByID(id)
	if exists {
		return result, nil
	}

	return r.resolver.ResolveByID(ctx, id)
}
