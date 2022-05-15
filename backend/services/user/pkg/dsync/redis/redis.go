package redis

import (
	"github.com/go-redis/redis/v8"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v8"
	"github.com/rkfcccccc/english_words/user/pkg/dsync"
)

type syncClient struct {
	rs *redsync.Redsync
}

func NewSyncClient(client *redis.Client) dsync.Client {
	return &syncClient{redsync.New(goredis.NewPool(client))}
}

func (client *syncClient) NewMutex(name string) dsync.Mutex {
	return client.rs.NewMutex(name)
}
