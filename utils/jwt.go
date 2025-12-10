package utils

import (
	"math/rand"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

/*func GETJWTSecret() ([]byte, error) {
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
}*/

func GenerateToken(username string) (string, error) {
	jti := time.Now().UnixNano() + rand.Int63()
	claims := jwt.MapClaims{
		"username": username,
		"jti":      jti,
		"exp":      time.Now().Add(time.Hour * 1).Unix(),
		"iat":      time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("lanshan"))
}
func GeneratefreshToken(username string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24 * 7).Unix(),
		"iat":      time.Now().Unix(),
		"type":     "fresh",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("lanshan"))
}
func ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte("lanshan"), nil
	})
}
