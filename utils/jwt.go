package utils

import (
    "github.com/golang-jwt/jwt/v4"
    "time"
)

var jwtKey = []byte("your-secret-key")

func GenerateJWT(userID uint) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "id": userID,
        "exp":     time.Now().Add(time.Hour * 1).Unix(), // หมดอายุใน 1 ชั่วโมง
    })

    return token.SignedString(jwtKey)
}