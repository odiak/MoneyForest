package util

import (
	uuid "github.com/satori/go.uuid"
)

func StringPtr(s string) *string {
	return &s
}

func UUIDPtr(u uuid.UUID) *uuid.UUID {
	return &u
}
