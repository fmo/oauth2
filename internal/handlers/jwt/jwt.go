// Package jwtutil is wrapper of jwt library
package jwtutil

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userID, clientID string) (string, error) {
	claims := jwt.MapClaims{
		"sub": userID,
		"iss": "localhost:8080",
		"aud": clientID,
		"exp": time.Now().Add(time.Hour).Unix(),
		"iat": time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte("my-secret"))
}
