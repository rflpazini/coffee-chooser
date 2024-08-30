package auth

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type SessionToken struct {
	SessionID   uuid.UUID `json:"session_id"`
	UserID      string    `json:"user_id"`
	Geolocation string    `json:"geolocation"`
	jwt.RegisteredClaims
}
