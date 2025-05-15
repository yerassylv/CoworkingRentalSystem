package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
	"user/internal/entity"

	"github.com/go-redis/redis/v8"
)

type RedisCache struct {
	client *redis.Client
	ttl    time.Duration
}

func NewRedisCache(client *redis.Client, ttl time.Duration) *RedisCache {
	return &RedisCache{client: client, ttl: ttl}
}

func (r *RedisCache) GetUserProfile(ctx context.Context, userID string) (*entity.User, error) {
	key := r.userKey(userID)
	data, err := r.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return nil, nil // cache miss
	} else if err != nil {
		return nil, err
	}

	var user entity.User
	if err := json.Unmarshal([]byte(data), &user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *RedisCache) SetUserProfile(ctx context.Context, user *entity.User) error {
	key := r.userKey(user.UserID)
	data, err := json.Marshal(user)
	if err != nil {
		return err
	}
	return r.client.Set(ctx, key, data, r.ttl).Err()
}

func (r *RedisCache) InvalidateUserProfile(ctx context.Context, userID string) error {
	return r.client.Del(ctx, r.userKey(userID)).Err()
}

func (r *RedisCache) userKey(userID string) string {
	return fmt.Sprintf("user_profile:%s", userID)
}
