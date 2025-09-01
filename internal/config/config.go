package config

import (
	"fmt"
	"os"
	"strconv"
	"sync"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

// exchange rates API configuration
type RateAPIConfig struct {
	APIKey string `env:"API_KEY"`
}

// Redis config
type RedisConfig struct {
	Host     string `env:"REDIS_HOST" envDefault:"localhost"`
	Port     int    `env:"REDIS_PORT" envDefault:"6379"`
	DB       int    `env:"REDIS_DB" envDefault:"0"`
	User     string `env:"REDIS_USER" envDefault:""`
	Password string `env:"REDIS_PASSWORD" envDefault:""`
	UseSSL   *bool  `env:"REDIS_USE_SSL" envDefault:"false"`
}

func (c *RedisConfig) BuildConnectionString() (string, error) {
	// Validate required fields
	if c.Host == "" {
		return "", fmt.Errorf("REDIS_HOST is required")
	}

	if c.Port <= 0 || c.Port > 65535 {
		return "", fmt.Errorf("REDIS_PORT must be between 1 and 65535")
	}

	if c.DB < 0 {
		return "", fmt.Errorf("REDIS_DB cannot be negative")
	}

	// Determine scheme based on SSL setting
	scheme := "redis"
	if c.UseSSL != nil && *c.UseSSL {
		scheme = "rediss"
	}

	urlBuilder := NewURLBuilder()
	urlBuilder.WithScheme(scheme).WithHost(c.Host).WithPort(c.Port)

	if c.User != "" || c.Password != "" {
		urlBuilder.WithCredentials(c.User, c.Password)
	}

	urlBuilder.WithPath(strconv.Itoa(c.DB))

	return urlBuilder.Build()
}

// server configuration
type ServerConfig struct {
	Host string `env:"HOST" envDefault:"0.0.0.0"`
	Port int    `env:"PORT" envDefault:"8080"`
}

// combined settings
type Settings struct {
	RatesConfig RateAPIConfig `envPrefix:"APILAYER_EXCHANGE_RATES_API_KEY"`
	Server      ServerConfig  `envPrefix:"SERVER_"`
}

// global settings instance
var settings Settings

// ensure that settings are loaded once
var onceLoad sync.Once

func LoadConfig(paths ...string) (err error) {
	onceLoad.Do(func() {
		for _, path := range paths {
			if path == "" {
				continue
			}

			if loadErr := godotenv.Load(path); loadErr != nil && !os.IsNotExist(loadErr) {
				err = loadErr
				return
			}
		}
		err = env.Parse(&settings)
	})
	return err
}

func GetSettings(paths ...string) Settings {
	err := LoadConfig(paths...)
	if err != nil {
		panic(fmt.Sprintf("Configuration load error: %s", err))
	}

	return settings
}
