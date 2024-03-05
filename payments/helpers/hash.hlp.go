package helpers

import (
	"log"

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
		return []byte(config.GetJWTSecretKey()), nil
	})

	if err != nil {
		log.Println(err)
		return false, "token invalid"
	}

	if !token.Valid {
		log.Println("asdsadsa")
		return false, "token invalid"
	}

	return true, ""
}
