package ports

import "deadpool/core/domain"

type UserRepository interface {
	Create(user *domain.User) error
	FindByGoogleID(googleID string) (*domain.User, error)
}

type AuthService interface {
	GenerateToken(user *domain.User) (string, error)
	ValidateToken(token string) (*domain.User, error)
}
