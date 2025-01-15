package utils

import (
    "github.com/golang-jwt/jwt/v4"
    "time"
)

func GenerateJWT(userID uint, secret string) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id": userID,
        "exp":     time.Now().Add(24 * time.Hour).Unix(),
    })
    return token.SignedString([]byte(secret))
}

func DecodeJWT(tokenString, secret string) (jwt.MapClaims, error) {
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        return []byte(secret), nil
    })

    if err != nil || !token.Valid {
        return nil, err
    }

    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok {
        return nil, err
    }

    return claims, nil
}