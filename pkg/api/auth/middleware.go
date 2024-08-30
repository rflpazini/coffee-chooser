package auth

import (
	"context"
	"net/http"

	authService "coffee-choose/pkg/auth"
	"github.com/labstack/echo/v4"
	"go.uber.org/dig"
)

type MiddlewareParams struct {
	dig.In

	ValidateSessionTokenFunc authService.ValidateSessionTokenFunc
}

func Middleware(p MiddlewareParams) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return echo.NewHTTPError(http.StatusUnauthorized, "missing authorization header")
			}

			tokenString := authHeader[len("Bearer "):]

			claims, err := p.ValidateSessionTokenFunc(c.Request().Context(), tokenString)
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "invalid or expired token")
			}

			ctx := context.WithValue(c.Request().Context(), "session", claims)
			c.SetRequest(c.Request().WithContext(ctx))
			return next(c)
		}
	}
}
