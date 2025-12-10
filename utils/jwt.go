package utils

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GETJWTSecret() ([]byte, error) {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return nil, errors.New("JWT_SECRET env variable is not set")
	}
	return []byte(secret), nil
}
func MakeToken(username string, expirationTime time.Time) (string, error) {
	secretKey, err := GETJWTSecret()
	if err != nil {
		return "", err
	}
	claims := jwt.MapClaims{
		"sub": username,
		"exp": expirationTime.Unix(),
		"iat": time.Now().Unix(),
		"iss": "my-gin-app",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
