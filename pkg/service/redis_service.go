package service

import (
	"context"
	"github.com/amrchnk/api-gateway/pkg/repository/cache"
	"time"
)

type RedisService struct {
	cache *cache.RedisClient
}

func NewRedisService(cache *cache.RedisClient) RedisService {
	return RedisService{cache}
}

func (r RedisService) GetFromCache(ctx context.Context, key string) ([]byte, error) {
	return r.cache.GetFromCache(ctx, key)
}

func (r RedisService) SetInCache(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return r.cache.SetInCache(ctx, key, value, expiration)
}

func (r RedisService) DeleteFromCache(ctx context.Context, key string) (int64, error) {
	return r.cache.DeleteFromCache(ctx, key)
}
