package services

import (
	"context"
	"fmt"
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
	symbolKey := fmt.Sprintf("symbol:%s", key)
	result := r.client.Get(ctx, symbolKey)
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
	symbolKey := fmt.Sprintf("symbol:%s", key)
	return r.client.Set(ctx, symbolKey, value, lifetime).Err()
}

func (r *RedisCache) GetUSDRate(ctx context.Context, key string) (float64, error) {
	rateKey := fmt.Sprintf("rate:%s", key)
	result := r.client.Get(ctx, rateKey)
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
	rateKey := fmt.Sprintf("rate:%s", key)
	return r.client.Set(ctx, rateKey, strconv.FormatFloat(value, 'f', -1, 64), lifetime).Err()
}

func (r *RedisCache) Close() {
	r.client.Conn().Close()
}
