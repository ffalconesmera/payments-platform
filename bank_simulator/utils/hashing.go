package utils

import (
	"github.com/google/uuid"
)

func NewUUIDString() string {
	return uuid.New().String()
}
