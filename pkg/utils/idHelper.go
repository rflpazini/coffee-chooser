package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"github.com/google/uuid"
)

func CreateUserId() string {
	return "cc-" + uuid.NewString()
}

func GenerateClientID(userAgent, ipAddress, additionalInfo string) string {
	rawData := fmt.Sprintf("%s|%s|%s", userAgent, ipAddress, additionalInfo)
	hash := sha256.New()
	hash.Write([]byte(rawData))

	return hex.EncodeToString(hash.Sum(nil))
}
