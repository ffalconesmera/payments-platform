package helpers

import (
	"github.com/ffalconesmera/payments-platform/payments/config"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// CustomHash is a singleton for create customs hashes and indentifiers
type customHash struct{}

var hash *customHash

func CustomHash() *customHash {
	if hash == nil {
		hash = &customHash{}
	}

	return hash
}

// Generate an identifier
func (h *customHash) NewUUIDString() string {
	return uuid.New().String()
}

// Check if a json web token is valid
func (h *customHash) CheckJWToken(tokenString string) (bool, string) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return config.Config().GetJWTSecretKey(), nil
	})

	if err != nil {
		return false, "token invalid"
	}

	if !token.Valid {
		return false, "token invalid"
	}

	return true, ""
}
