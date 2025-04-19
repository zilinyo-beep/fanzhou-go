package service

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var jwtSecret = []byte("fanzhou-secret")

func GenerateJWT(phone string, userId int64) (string, error) {
	claims := jwt.MapClaims{
		"phone":  phone,
		"userId": userId,
		"exp":    time.Now().Add(24 * time.Hour).Unix(), // 24小时有效期
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}
