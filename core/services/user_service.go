package services

import (
	"deadpool/core/ports"
)

type UserService struct {
	UserRepo ports.UserRepository
}
