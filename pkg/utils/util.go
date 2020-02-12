package utils

import "github.com/google/uuid"

// NewUUID as string
func NewUUID() string {
	return uuid.New().String()
}
