package auth

import (
	"aurora-graph/account/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTCustomClaims struct {
	UserID uint64 `json:"user_id"`
	jwt.RegisteredClaims
}

func GenerateToken(userID uint64)(string, error) {
	claims := &JWTCustomClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: config.Issuer,
			IssuedAt: jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(config.SecretKey))
}