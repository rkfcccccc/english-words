package vocabulary

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

type redisVocabularyState struct {
	redis *redis.Client
}

func NewRedisVocabularyState(redis *redis.Client) State {
	return &redisVocabularyState{redis}
}

func (c *redisVocabularyState) Get(ctx context.Context, userId int) (int, error) {
	n, err := c.redis.Get(ctx, fmt.Sprintf("vocabulary_state:challenge:%d", userId)).Int()
	if err == redis.Nil {
		return 0, nil
	}

	return int(n), err
}

func (c *redisVocabularyState) Increment(ctx context.Context, userId int) (int, error) {
	n, err := c.redis.Incr(ctx, fmt.Sprintf("vocabulary_state:challenge:%d", userId)).Result()
	return int(n), err
}
