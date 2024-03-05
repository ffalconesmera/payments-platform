package helpers

import (
	"strings"
	"time"
)

// CustomValidation is a singleton for returning errors
type customValidation struct{}

var cValid *customValidation

func CustomValidation() *customValidation {
	if cValid == nil {
		cValid = &customValidation{}
	}

	return cValid
}

func (c *customValidation) InvalidFloat(f float64) bool {
	return f <= 0
}

func (c *customValidation) EmptyString(s string) bool {
	return strings.Trim(s, " ") == ""
}

func (c *customValidation) DateInvalid(s string) bool {
	_, err := time.Parse("2006-01-02", s)
	return err != nil
}

func (c *customValidation) DateTimeInvalid(s string) bool {
	_, err := time.Parse("2006-01-02 15:04:05", s)
	return err != nil
}
