package auth

import (
	"context"
	"errors"
	"time"

	"coffee-choose/pkg/config"
	"coffee-choose/pkg/service/geo"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"go.uber.org/dig"
)

type ServiceParams struct {
	dig.In

	*config.JwtConfig
}

type CreateSessionTokenFunc func(ctx context.Context, userID string, geo *geo.Location) (string, error)

func makeCreateSessionToken(p ServiceParams) CreateSessionTokenFunc {
	return func(ctx context.Context, userID string, geo *geo.Location) (string, error) {
		claims := &SessionToken{
			SessionID:   uuid.New(),
			UserID:      userID,
			Geolocation: geo,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
				IssuedAt:  jwt.NewNumericDate(time.Now()),
				Issuer:    p.JwtConfig.Issuer,
				Audience:  jwt.ClaimStrings{p.JwtConfig.Audience},
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		return token.SignedString([]byte(p.JwtConfig.Secret))
	}
}

type ValidateSessionTokenFunc func(ctx context.Context, tokenString string) (*SessionToken, error)

func makeValidateSessionToken(p ServiceParams) ValidateSessionTokenFunc {
	return func(ctx context.Context, tokenString string) (*SessionToken, error) {
		token, err := jwt.ParseWithClaims(tokenString, &SessionToken{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("unexpected signing method")
			}
			return []byte(p.JwtConfig.Secret), nil
		})

		if err != nil {
			return nil, err
		}

		if claims, ok := token.Claims.(*SessionToken); ok && token.Valid {
			return claims, nil
		}

		return nil, errors.New("invalid token")
	}
}
