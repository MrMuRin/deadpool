package ports

import "deadpool/core/domain"

type UserRepository interface {
	Create(user *domain.User) error
	FindByGoogleID(googleID string) (*domain.User, error)
	FindByID(userID uint) (*domain.User, error)
	FindByEmail(email string) (*domain.User, error)
	Update(user *domain.User) error
	Delete(userID uint) error
}

type AuthService interface {
	GenerateToken(user *domain.User) (string, error)
	ValidateToken(token string) (*domain.User, error)
	RefreshToken(oldToken string) (string, error)
}
