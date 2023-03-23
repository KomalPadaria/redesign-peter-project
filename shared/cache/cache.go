package cache

import (
	"context"
	"time"

	"github.com/eko/gocache/lib/v4/cache"
	"github.com/eko/gocache/lib/v4/store"
	gocache_store "github.com/eko/gocache/store/go_cache/v4"
	gocache "github.com/patrickmn/go-cache"
)

type Client interface {
	Set(ctx context.Context, key string, value string, options ...store.Option) error
	Get(ctx context.Context, key string) (string, error)
}

func New(defaultExpiration time.Duration, cleanupInterval time.Duration) Client {
	gocacheClient := gocache.New(defaultExpiration, cleanupInterval)
	gocacheStore := gocache_store.NewGoCache(gocacheClient)

	cacheManager := cache.New[string](gocacheStore)

	return &lcache{cacheManager}
}

type lcache struct {
	cacheManager *cache.Cache[string]
}

func (l lcache) Set(ctx context.Context, key string, value string, options ...store.Option) error {
	return l.cacheManager.Set(ctx, key, value, options...)
}

func (l lcache) Get(ctx context.Context, key string) (string, error) {
	return l.cacheManager.Get(ctx, key)
}
