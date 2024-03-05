package helpers

import (
	"strings"
	"time"
	"unicode"
)

// CustomValidation is a singleton for returning errors
func InvalidFloat(f float64) bool {
	return f <= 0
}

func EmptyString(s string) bool {
	return strings.Trim(s, " ") == ""
}

func DateInvalid(s string) bool {
	_, err := time.Parse("2006-01-02", s)
	return err != nil
}

func DateTimeInvalid(s string) bool {
	_, err := time.Parse("2006-01-02 15:04:05", s)
	return err != nil
}

func PasswordInvalid(s string) bool {
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

func PasswordInvalidMessage() string {
	return "must have at least 8 characters, one uppercase letter, one lowercase letter, one number and one special character . _"
}
