package helpers

import (
	"strings"
	"time"
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
