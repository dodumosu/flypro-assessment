package repositories

import (
	"context"
	"errors"

	"flypro-assessment/internal/models"

	"gorm.io/gorm"
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

func (u userRepository) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	db := u.db.WithContext(ctx)
	var user models.User
	result := db.Where("lower(email) = lower(?)", email).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, ErrEntityDoesNotExist
		}
		return nil, result.Error
	}
	return &user, nil
}

func (u userRepository) Create(ctx context.Context, instance *models.User) (*models.User, error) {
	db := u.db.WithContext(ctx)
	result := db.Create(instance)
	if result.Error != nil {
		// if a unique key (email/id) is specified
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return nil, ErrEntityExists
		}
	}

	return u.GetByEmail(ctx, instance.Email)
}

func (u userRepository) Get(ctx context.Context, id int) (*models.User, error) {
	db := u.db.WithContext(ctx)
	var user models.User

	result := db.Where("id = ?", id).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, ErrEntityDoesNotExist
		}
		return nil, result.Error
	}

	return &user, nil
}
