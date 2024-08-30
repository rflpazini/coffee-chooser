//go:generate easyjson -lower_camel_case $GOFILE
package auth

import (
	"coffee-choose/pkg/service/geo"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

//easyjson:json
type SessionToken struct {
	SessionID   uuid.UUID     `json:"session_id"`
	UserID      string        `json:"user_id"`
	Geolocation *geo.Location `json:"geolocation"`
	jwt.RegisteredClaims
}
