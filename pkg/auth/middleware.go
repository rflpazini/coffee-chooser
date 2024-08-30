package auth

import (
	"errors"
	"time"

	"coffee-choose/pkg/config"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"go.uber.org/dig"
)

type CreateSessionTokenParams struct {
	dig.In

	*config.JwtConfig
}

func CreateSessionToken(p CreateSessionTokenParams, userID, geolocation string) (string, error) {
	claims := &SessionToken{
		SessionID:   uuid.New(),
		UserID:      userID,
		Geolocation: geolocation,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // Token expires after 24 hours
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    p.Issuer,
			Audience:  jwt.ClaimStrings{p.Audience},
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(p.Secret)
}

func ValidateSessionToken(p CreateSessionTokenParams, tokenString string) (*SessionToken, error) {
	token, err := jwt.ParseWithClaims(tokenString, &SessionToken{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return p.Secret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*SessionToken); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
