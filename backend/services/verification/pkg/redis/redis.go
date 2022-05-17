package redis

import (
	"fmt"

	"github.com/go-redis/redis/v8"
)

func NewClient(host, port string) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Password: "",
		DB:       0,
	})
}
