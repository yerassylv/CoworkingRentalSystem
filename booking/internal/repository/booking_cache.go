package repository

import (
	"booking/internal/entity"
	"context"
	"encoding/json"
	"fmt"
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

func (r *RedisCache) GetBookings(ctx context.Context, userID string) ([]*entity.Booking, error) {
	key := r.key(userID)
	data, err := r.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return nil, nil // Cache miss
	} else if err != nil {
		return nil, err
	}

	var bookings []*entity.Booking
	if err := json.Unmarshal([]byte(data), &bookings); err != nil {
		return nil, err
	}
	return bookings, nil
}

func (r *RedisCache) SetBookings(ctx context.Context, userID string, bookings []*entity.Booking) error {
	key := r.key(userID)
	data, err := json.Marshal(bookings)
	if err != nil {
		return err
	}
	return r.client.Set(ctx, key, data, r.ttl).Err()
}

func (r *RedisCache) InvalidateBookings(ctx context.Context, userID string) error {
	return r.client.Del(ctx, r.key(userID)).Err()
}

func (r *RedisCache) key(userID string) string {
	return fmt.Sprintf("user_bookings:%s", userID)
}
