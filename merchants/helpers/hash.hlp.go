package helpers

import (
	"time"

	"github.com/ffalconesmera/payments-platform/merchants/config"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// CustomHash is a singleton for create customs hashes and indentifiers

// Generate hash for stored passwords
func GenerateHashPassword(password string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), 15)
	return string(hash)
}

// Compare hash with text
func CheckHashPassword(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// Generate an identifier
func NewUUIDString() string {
	return uuid.New().String()
}

// Generate a json web token for authorization
func CreateJWToken(merchantCode string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"merchant_code": merchantCode,
			"exp":           time.Now().Add(time.Minute * time.Duration(config.GetJWTExpiration())).Unix(),
		})

	tokenString, err := token.SignedString([]byte(config.GetJWTSecretKey()))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
