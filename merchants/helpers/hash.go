package helpers

import (
	"time"

	"github.com/ffalconesmera/payments-platform/merchants/config"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
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

// Generate hash for stored passwords
func (c *customHash) GenerateHashPassword(password string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), 15)
	return string(hash)
}

// Compare hash with text
func (c *customHash) CheckHashPassword(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// Generate an identifier
func (c *customHash) NewUUIDString() string {
	return uuid.New().String()
}

// Generate a json web token for authorization
func (c *customHash) CreateJWToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"exp":      time.Now().Add(time.Minute * time.Duration(config.Config().GetJWTExpiration())).Unix(),
		})

	tokenString, err := token.SignedString([]byte(config.Config().GetJWTSecretKey()))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// Check if a json web token is valid
func (c *customHash) CheckJWToken(tokenString string) (bool, string) {
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
