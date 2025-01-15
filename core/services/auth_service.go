package services

import (
    "deadpool/core/ports"
    "deadpool/infrastructure/utils"
    "errors"
)

type AuthService struct {
    JWTSecret string
    UserRepo  ports.UserRepository
}

func NewAuthService(jwtSecret string, userRepo ports.UserRepository) *AuthService {
    return &AuthService{
        JWTSecret: jwtSecret,
        UserRepo:  userRepo,
    }
}

// ValidateToken ตรวจสอบ JWT และคืนค่า userID
func (s *AuthService) ValidateToken(token string) (uint, error) {
    claims, err := utils.DecodeJWT(token, s.JWTSecret)
    if err != nil {
        return 0, errors.New("invalid token")
    }

    userID, ok := claims["user_id"].(float64)
    if !ok {
        return 0, errors.New("invalid token payload")
    }

    return uint(userID), nil
}

// GenerateToken สร้าง JWT สำหรับผู้ใช้
func (s *AuthService) GenerateToken(userID uint) (string, error) {
    return utils.GenerateJWT(userID, s.JWTSecret)
}

// RefreshToken สร้าง Token ใหม่จาก Token เก่า
func (s *AuthService) RefreshToken(oldToken string) (string, error) {
    claims, err := utils.DecodeJWT(oldToken, s.JWTSecret)
    if err != nil {
        return "", errors.New("invalid token")
    }

    userID, ok := claims["user_id"].(float64)
    if !ok {
        return "", errors.New("invalid token payload")
    }

    return utils.GenerateJWT(uint(userID), s.JWTSecret)
}

// HasRole ตรวจสอบว่าผู้ใช้มีสิทธิ์ที่ต้องการหรือไม่
// func (s *AuthService) HasRole(userID uint, requiredRole string) (bool, error) {
//     user, err := s.UserRepo.FindByID(userID)
//     if err != nil {
//         return false, err
//     }

//     for _, role := range user.Roles {
//         if role == requiredRole {
//             return true, nil
//         }
//     }

//     return false, nil
// }