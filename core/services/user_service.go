package services

import (
    "deadpool/core/domain"
    "deadpool/core/ports"
    "errors"
)

type UserService struct {
    UserRepo ports.UserRepository
}

func NewUserService(userRepo ports.UserRepository) *UserService {
    return &UserService{
        UserRepo:  userRepo,
    }
}

func (s *UserService) GetUserByID(userID uint) (*domain.User, error) {
    user, err := s.UserRepo.FindByID(userID)
    if err != nil {
        return nil, err
    }

    if user == nil {
        return nil, errors.New("user not found")
    }

    return user, nil
}

func (s *UserService) CreateUser(user *domain.User) error {
    return s.UserRepo.Create(user)
}

func (s *UserService) UpdateUser(user *domain.User) error {
    return s.UserRepo.Update(user)
}


func (s *UserService) DeleteUser(userID uint) error {
    return s.UserRepo.Delete(userID)
}
