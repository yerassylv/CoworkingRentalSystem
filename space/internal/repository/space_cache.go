package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"space/internal/entity"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisCache struct {
	client *redis.Client
	ttl    time.Duration
}

func NewRedisCache(client *redis.Client, ttl time.Duration) *RedisCache {
	return &RedisCache{client: client, ttl: ttl}
}

func (r *RedisCache) GetSpace(ctx context.Context, spaceID string) (*entity.Space, error) {
	key := r.spaceKey(spaceID)
	data, err := r.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	var space entity.Space
	if err := json.Unmarshal([]byte(data), &space); err != nil {
		return nil, err
	}
	return &space, nil
}

func (r *RedisCache) SetSpace(ctx context.Context, space *entity.Space) error {
	key := r.spaceKey(space.SpaceID)
	data, err := json.Marshal(space)
	if err != nil {
		return err
	}
	return r.client.Set(ctx, key, data, r.ttl).Err()
}

func (r *RedisCache) InvalidateSpace(ctx context.Context, spaceID string) error {
	return r.client.Del(ctx, r.spaceKey(spaceID)).Err()
}

func (r *RedisCache) spaceKey(spaceID string) string {
	return fmt.Sprintf("space_profile:%s", spaceID)
}
