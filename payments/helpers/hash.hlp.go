package helpers

import (
	"github.com/ffalconesmera/payments-platform/payments/config"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// CustomHash is a singleton for create customs hashes and indentifiers
// Generate an identifier
func NewUUIDString() string {
	return uuid.New().String()
}

// Check if a json web token is valid
func CheckJWToken(tokenString string) (bool, string) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return config.GetJWTSecretKey(), nil
	})

	if err != nil {
		return false, "token invalid"
	}

	if !token.Valid {
		return false, "token invalid"
	}

	return true, ""
}
