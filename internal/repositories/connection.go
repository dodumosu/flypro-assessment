package repositories

import (
	"flypro-assessment/internal/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConnection struct {
	*gorm.DB
}

func NewConnection(cfg config.DatabaseConfig) (*DBConnection, error) {
	connectionString, err := cfg.BuildDSN()
	if err != nil {
		return nil, err
	}

	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &DBConnection{
		DB: db,
	}, nil
}
