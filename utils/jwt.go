package utils

import (
	"github.com/golang-jwt/jwt"
)

func GenerateJWTToken(claims jwt.MapClaims) (string, error) {
	// Your secret key for signing the token (keep it secret!)
	secretKey := []byte("master-beast")
	// secretKey := os.Getenv("MY_SECRET")

	// Create a new token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
