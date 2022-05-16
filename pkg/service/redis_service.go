package service

import (
	"context"
	"github.com/amrchnk/api-gateway/pkg/repository/cache"
	"time"
)

type RedisService struct {
	*cache.RedisClient
}

func NewRedisService(cache *cache.RedisClient) RedisService {
	return RedisService{cache}
}

func (r RedisService) GetFromCache(ctx context.Context, key string) ([]byte, error) {
	return r.GetFromCache(ctx, key)
}

func (r RedisService) SetInCache(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return r.SetInCache(ctx, key, value, expiration)
}
