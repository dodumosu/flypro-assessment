package services

import (
	"context"
	"testing"
	"time"

	"flypro-assessment/internal/config"

	"github.com/stretchr/testify/assert"
)

func TestRedisCache_InterfaceImplementation(t *testing.T) {
	// Test that the struct implements the interface correctly
	assert.Implements(t, (*Cache)(nil), new(RedisCache))
}

func TestNewRedisCache(t *testing.T) {
	// Test with valid config
	cfg := config.RedisConfig{
		Host: "localhost",
		Port: 6379,
		DB:   0,
	}

	cache, err := NewRedisCache(cfg)
	assert.NoError(t, err)
	assert.NotNil(t, cache)
	assert.NotNil(t, cache.client)

	// Test with invalid config (negative port)
	invalidCfg := config.RedisConfig{
		Host: "localhost",
		Port: -1,
		DB:   0,
	}

	invalidCache, err := NewRedisCache(invalidCfg)
	assert.Error(t, err)
	assert.Nil(t, invalidCache)

	// Test with another invalid config (empty host)
	invalidCfg2 := config.RedisConfig{
		Host: "",
		Port: 6379,
		DB:   0,
	}

	invalidCache2, err := NewRedisCache(invalidCfg2)
	assert.Error(t, err)
	assert.Nil(t, invalidCache2)
}

// Example tests that would require a real Redis instance
func TestRedisCache_GetSymbol_Example(t *testing.T) {
	// This is just an example of how the test would look
	// In a real scenario, you would need to start a Redis instance for testing
	t.Skip("Skipping integration test - requires Redis instance")

	cfg := config.RedisConfig{
		Host: "localhost",
		Port: 6379,
		DB:   0,
	}

	cache, err := NewRedisCache(cfg)
	assert.NoError(t, err)

	ctx := context.Background()
	result, err := cache.GetSymbol(ctx, "test-key")
	assert.NoError(t, err)
	assert.Empty(t, result) // Key doesn't exist yet
}

func TestRedisCache_SetSymbol_Example(t *testing.T) {
	// This is just an example of how the test would look
	// In a real scenario, you would need to start a Redis instance for testing
	t.Skip("Skipping integration test - requires Redis instance")

	cfg := config.RedisConfig{
		Host: "localhost",
		Port: 6379,
		DB:   0,
	}

	cache, err := NewRedisCache(cfg)
	assert.NoError(t, err)

	ctx := context.Background()
	err = cache.SetSymbol(ctx, "test-key", "test-value", time.Minute)
	assert.NoError(t, err)
}

func TestRedisCache_GetUSDRate_Example(t *testing.T) {
	// This is just an example of how the test would look
	// In a real scenario, you would need to start a Redis instance for testing
	t.Skip("Skipping integration test - requires Redis instance")

	cfg := config.RedisConfig{
		Host: "localhost",
		Port: 6379,
		DB:   0,
	}

	cache, err := NewRedisCache(cfg)
	assert.NoError(t, err)

	ctx := context.Background()
	result, err := cache.GetUSDRate(ctx, "test-rate-key")
	assert.NoError(t, err)
	assert.Equal(t, float64(0), result) // Key doesn't exist yet
}

func TestRedisCache_SetUSDRate_Example(t *testing.T) {
	// This is just an example of how the test would look
	// In a real scenario, you would need to start a Redis instance for testing
	t.Skip("Skipping integration test - requires Redis instance")

	cfg := config.RedisConfig{
		Host: "localhost",
		Port: 6379,
		DB:   0,
	}

	cache, err := NewRedisCache(cfg)
	assert.NoError(t, err)

	ctx := context.Background()
	err = cache.SetUSDRate(ctx, "test-rate-key", 1.2345, time.Minute)
	assert.NoError(t, err)
}
