package cache

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"sync"
	"time"
)

var (
	redisClient *RedisClient
	lock        = &sync.Mutex{}
)

type RedisClient struct {
	client *redis.Client
}

type Options struct {
	Addr     string
	Password string
	DB       int
}

func NewRedisClient(ctx context.Context, options *Options) (*RedisClient, error) {
	lock.Lock()
	defer lock.Unlock()

	if redisClient == nil {
		client, err := getRedisClient(ctx, options)
		if err != nil {
			return nil, fmt.Errorf("failed to initialize redis client, error is: %s", err)
		}
		redisClient = &RedisClient{
			client: client,
		}
		return redisClient, nil
	}

	return redisClient, nil
}

func getRedisClient(ctx context.Context, options *Options) (*redis.Client, error) {
	opts := redis.Options{
		Addr:     options.Addr,
		Password: options.Password,
		DB:       options.DB,
	}
	client := redis.NewClient(&opts)

	if err := client.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("redis client with ttl %s failed to ping address %s, error is: %s",
			opts.IdleTimeout, opts.Addr, err)
	}

	return client, nil
}

func (rc *RedisClient) SetInCache(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return rc.client.Set(ctx, key, value, expiration).Err()
}

func (rc *RedisClient) GetFromCache(ctx context.Context, key string) ([]byte, error) {
	return rc.client.Get(ctx, key).Bytes()
}

func (rc *RedisClient) DeleteFromCache(ctx context.Context, key string) (int64, error) {
	return rc.client.Del(ctx, key).Result()
}
