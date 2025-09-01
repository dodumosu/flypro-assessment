package config

import (
	"fmt"
	"os"
	"sync"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

// server configuration
type ServerConfig struct {
	Host string `env:"HOST" envDefault:"0.0.0.0"`
	Port int    `env:"PORT" envDefault:"8080"`
}

// combined settings
type Settings struct {
	Server ServerConfig `envPrefix:"SERVER_"`
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
