package utils

import "github.com/google/uuid"

func CreateUserId() string {
	return "cc-" + uuid.NewString()
}
