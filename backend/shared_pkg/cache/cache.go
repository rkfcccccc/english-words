package cache

import (
	"context"
	"errors"
	"time"
)

var ErrCacheMiss = errors.New("cache: key is missing")

type Repository interface {
	Get(ctx context.Context, key string, value interface{}) error

	Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error

	Del(ctx context.Context, key string) error
}
