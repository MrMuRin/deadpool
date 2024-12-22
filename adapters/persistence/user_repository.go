package persistence

import (
	"deadpool/core/domain"
	"errors"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) FindByGoogleID(googleID string) (*domain.User, error) {
	var user domain.User
	result := r.db.Where("google_id = ?", googleID).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &user, result.Error
}

func (r *UserRepository) Create(user *domain.User) error {
	return r.db.Create(user).Error
}
