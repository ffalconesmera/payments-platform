package helpers

import (
	"errors"
	"fmt"
	"strings"
	"time"
	"unicode"
)

// CustomError is an interface for returning errors
type CustomError interface {
	NewError(message string) error
	EmptyMesage(caption string) error
	InvalidFloat(f float64, caption string) error
	EmptyString(s string, caption string) error
	DateInvalid(s string, caption string) error
	PasswordInvalid(s string, caption string) error
}

type customError struct{}

func NewCustomError() *customError {
	return &customError{}
}

func (c *customError) NewError(message string) error {
	return errors.New(message)
}

func (c *customError) EmptyMesage(caption string) error {
	return c.NewError(fmt.Sprintf("%s is empty", caption))
}

func (c *customError) InvalidFloat(f float64, caption string) error {
	if f <= 0 {
		return c.NewError(fmt.Sprintf("%s: could not be zero", caption))
	}

	return nil
}

func (c *customError) EmptyString(s string, caption string) error {
	if strings.Trim(s, " ") == "" {
		return c.NewError(fmt.Sprintf("%s: could not be empty", caption))
	}

	return nil
}

func (c *customError) DateInvalid(s string, caption string) error {
	_, err := time.Parse("2006-01-02", s)

	if err != nil {
		return c.NewError(fmt.Sprintf("%s:  format is invalid (yyyy-mm-dd). %s", caption, err))
	}

	return nil
}

func (c *customError) DateTimeInvalid(s string, caption string) error {
	_, err := time.Parse("2006-01-02 15:04:05", s)

	if err != nil {
		return c.NewError(fmt.Sprintf("%s: format is invalid (yyyy-mm-dd hh:ii:ss). %s", caption, err))
	}

	return nil
}

func (c *customError) PasswordInvalid(s string, caption string) error {
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
		return c.NewError(fmt.Sprintf("%s: %s", caption, "must have at least 8 characters, one uppercase letter, one lowercase letter, one number and one special character . _"))
	}

	return nil
}
