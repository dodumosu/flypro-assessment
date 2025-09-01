package config

import (
	"fmt"
	"log/slog"
	"os"
	"strconv"
	"sync"

	"github.com/caarlos0/env/v11"
	"github.com/charmbracelet/log"
	"github.com/joho/godotenv"
)

type APIConfig struct {
	// description for OpenAPI docs
	Description string `env:"API_DESCRIPTION" envDefault:"App"`
	// default list size
	ListSize int `env:"API_LIST_SIZE" envDefault:"50"`
	// version string for OpenAPI docs
	Version string `env:"API_VERSION" envDefault:"1.0.0"`
}

type LogStyle string

const (
	ColouredLogStyle LogStyle = "colour"
	JSONLogStyle     LogStyle = "json"
	PlainLogStyle    LogStyle = "plain"
)

// logging config
type LogConfig struct {
	LogLevel string   `env:"LOG_LEVEL" envDefault:"info"`
	Style    LogStyle `env:"LOG_STYLE" envDefault:"json"`
}

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
	AllowedOrigins []string `env:"ALLOWED_ORIGINS"`
	Host           string   `env:"HOST" envDefault:"0.0.0.0"`
	Port           int      `env:"PORT" envDefault:"8080"`
}

// combined settings
type Settings struct {
	APISettings APIConfig
	Logging     LogConfig
	RatesConfig RateAPIConfig `envPrefix:"APILAYER_EXCHANGE_RATES_API_KEY"`
	Server      ServerConfig  `envPrefix:"SERVER_"`
}

// root logger
var rootLogger *slog.Logger

func createLogger(cfg LogConfig) *slog.Logger {
	var logger *slog.Logger
	var logLevel slog.Level
	logParseError := logLevel.UnmarshalText([]byte(cfg.LogLevel))
	if logParseError != nil {
		logLevel = slog.LevelInfo
	}

	switch cfg.Style {
	case ColouredLogStyle:
		lvl, err := log.ParseLevel(cfg.LogLevel)
		if err != nil {
			lvl = log.InfoLevel
		}
		charmLogger := log.NewWithOptions(os.Stdout, log.Options{
			ReportTimestamp: true,
			ReportCaller:    true,
			Level:           lvl,
		})
		logger = slog.New(charmLogger)
	case JSONLogStyle:
		logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			AddSource: true,
			Level:     logLevel,
		}))
	case PlainLogStyle:
		logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			AddSource: true,
			Level:     logLevel,
		}))
	}

	return logger
}

func GetRootLogger(cfg LogConfig) *slog.Logger {
	loggerLoad.Do(func() {
		rootLogger = createLogger(cfg)
	})

	return rootLogger
}

// global settings instance
var settings Settings

// ensure that settings and logger are loaded/created once
var settingsLoad sync.Once
var loggerLoad sync.Once

func LoadConfig(paths ...string) (err error) {
	settingsLoad.Do(func() {
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
