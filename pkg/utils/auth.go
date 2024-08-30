package utils

import (
	"encoding/base64"
	"fmt"
)

const (
	WebDevice = "web"
)

func CreateAuthSubjectForDevice(country, uuid string) string {
	encodedUUID := base64.StdEncoding.EncodeToString([]byte(uuid))
	encodedUUID = encodedUUID[:len(encodedUUID)-2]

	return fmt.Sprintf("%s:%s:%s", WebDevice, country, encodedUUID)
}
