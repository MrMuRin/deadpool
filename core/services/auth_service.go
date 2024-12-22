package services

import (
	"deadpool/core/domain"
	"deadpool/core/ports"
)

type AuthService struct {
	UserRepo ports.UserRepository
}

func (s *AuthService) LoginWithGoogle(userInfo map[string]interface{}) (*domain.User, error) {
	googleID := userInfo["id"].(string)
	email := userInfo["email"].(string)
	name := userInfo["name"].(string)
	avatar := userInfo["picture"].(string)

	// ตรวจสอบว่าผู้ใช้มีอยู่ในระบบแล้วหรือไม่
	user, err := s.UserRepo.FindByGoogleID(googleID)
	if err != nil {
		return nil, err
	}

	if user == nil {
		// สร้างผู้ใช้ใหม่
		user = &domain.User{
			GoogleID: googleID,
			Name:     name,
			Email:    email,
			Avatar:   avatar,
		}
		if err := s.UserRepo.Create(user); err != nil {
			return nil, err
		}
	}

	return user, nil
}
