package services

import (
	"context"
	"strconv"
	"time"

	"flypro-assessment/internal/config"

	"github.com/redis/go-redis/v9"
)

type Cache interface {
	GetSymbol(ctx context.Context, key string) (string, error)
	SetSymbol(ctx context.Context, key string, value string, lifetime time.Duration) error
	GetUSDRate(ctx context.Context, key string) (float64, error)
	SetUSDRate(ctx context.Context, key string, value float64, lifetime time.Duration) error
}

type RedisCache struct {
	client *redis.Client
}

func NewRedisCache(cfg config.RedisConfig) (*RedisCache, error) {
	connectionString, err := cfg.BuildConnectionString()
	if err != nil {
		return nil, err
	}

	opt, err := redis.ParseURL(connectionString)
	if err != nil {
		return nil, err
	}
	client := redis.NewClient(opt)

	return &RedisCache{
		client: client,
	}, nil
}

func (r *RedisCache) GetSymbol(ctx context.Context, key string) (string, error) {
	result := r.client.Get(ctx, key)
	if err := result.Err(); err != nil {
		if err == redis.Nil {
			// Key does not exist
			return "", nil
		}
		// Other error occurred
		return "", err
	}
	return result.String(), nil
}

func (r *RedisCache) SetSymbol(ctx context.Context, key string, value string, lifetime time.Duration) error {
	return r.client.Set(ctx, key, value, lifetime).Err()
}

func (r *RedisCache) GetUSDRate(ctx context.Context, key string) (float64, error) {
	result := r.client.Get(ctx, key)
	if err := result.Err(); err != nil {
		if err == redis.Nil {
			// Key does not exist
			return 0, nil
		}
		// Other error occurred
		return 0, err
	}

	// Parse the string value to float64
	val, err := strconv.ParseFloat(result.Val(), 64)
	if err != nil {
		return 0, err
	}

	return val, nil
}

func (r *RedisCache) SetUSDRate(ctx context.Context, key string, value float64, lifetime time.Duration) error {
	return r.client.Set(ctx, key, strconv.FormatFloat(value, 'f', -1, 64), lifetime).Err()
}
