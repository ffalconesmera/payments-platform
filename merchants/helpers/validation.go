package helpers

import (
	"strings"
	"time"
	"unicode"
)

// CustomValidation is an interface for returning errors
type CustomValidation interface {
	InvalidFloat(f float64) bool
	EmptyString(s string) bool
	DateInvalid(s string) bool
	PasswordInvalid(s string) bool
	PasswordInvalidMessage() string
}

type customValidation struct{}

func NewCustomValidation() *customValidation {
	return &customValidation{}
}

func (c *customValidation) InvalidFloat(f float64) bool {
	if f <= 0 {
		return true
	}

	return false
}

func (c *customValidation) EmptyString(s string) bool {
	if strings.Trim(s, " ") == "" {
		return true
	}

	return false
}

func (c *customValidation) DateInvalid(s string) bool {
	_, err := time.Parse("2006-01-02", s)

	if err != nil {
		return true
	}

	return false
}

func (c *customValidation) DateTimeInvalid(s string) bool {
	_, err := time.Parse("2006-01-02 15:04:05", s)

	if err != nil {
		return true
	}

	return false
}

func (c *customValidation) PasswordInvalid(s string) bool {
	hasUpper := false
	hasLower := false
	hasNumber := false
	hasEspecialCharacter := false
	for _, r := range s {
		if unicode.IsUpper(r) {
			hasUpper = true
		}
		if unicode.IsLower(r) {
			hasLower = true
		}
		if r == '.' || r == '_' {
			hasEspecialCharacter = true
		}
		if unicode.IsNumber(r) {
			hasNumber = true
		}
	}

	if len(s) < 8 || !hasUpper || !hasLower || !hasNumber || !hasEspecialCharacter {
		return true
	}

	return false
}

func (c *customValidation) PasswordInvalidMessage() string {
	return "must have at least 8 characters, one uppercase letter, one lowercase letter, one number and one special character . _"
}
