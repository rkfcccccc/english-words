package redcache

import (
	"context"
	"time"

	rcache "github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
	"github.com/rkfcccccc/english_words/dictionary/pkg/cache"
)

type repository struct {
	cache *rcache.Cache
}

func NewCacheRepository(client *redis.Client) cache.Repository {
	return &repository{
		rcache.New(&rcache.Options{
			Redis:      client,
			LocalCache: rcache.NewTinyLFU(1<<20, time.Minute),
		}),
	}
}

func (repo *repository) Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	return repo.cache.Set(&rcache.Item{
		Ctx:   ctx,
		Key:   key,
		Value: value,
		TTL:   ttl,
	})
}

func (repo *repository) Get(ctx context.Context, key string, value interface{}) error {
	err := repo.cache.Get(ctx, key, &value)
	if err == rcache.ErrCacheMiss {
		return cache.ErrCacheMiss
	}

	return err
}

func (repo *repository) Del(ctx context.Context, key string) error {
	return repo.cache.Delete(ctx, key)
}
