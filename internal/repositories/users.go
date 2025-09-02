package repositories

import (
	"context"

	"flypro-assessment/internal/models"
)

type UserRepository interface {
	Create(ctx context.Context, instance *models.User) (*models.User, error)
	Get(ctx context.Context, id int) (*models.User, error)
}

type userRepository struct {
	db *DBConnection
}

func NewUserRepository(db *DBConnection) UserRepository {
	return &userRepository{db: db}
}

func (u userRepository) Create(ctx context.Context, instance *models.User) (*models.User, error) {
	return nil, nil
}

func (u userRepository) Get(ctx context.Context, id int) (*models.User, error) {
	return nil, nil
}
