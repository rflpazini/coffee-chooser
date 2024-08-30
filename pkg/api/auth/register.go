package auth

import (
	"coffee-choose/internal/router"
	"github.com/labstack/echo/v4"
	"go.uber.org/dig"
)

type sessionTokenParams struct {
	dig.In

	*echo.Echo
	SessionTokenHandler echo.HandlerFunc `name:"Route.Handler.SessionToken"`
}

func setupSessionTokenRoutes(p sessionTokenParams) router.RouteGroup {
	authGroup := p.Echo.Group("/v1/auth")

	authGroup.GET("/start", p.SessionTokenHandler)

	return router.RouteGroup{Group: authGroup}
}

func Register(c *dig.Container, register func(...interface{}) error) error {
	if err := c.Provide(makeSessionTokenHandler, dig.Name("Route.Handler.SessionToken")); err != nil {
		return err
	}
	return register(setupSessionTokenRoutes)
}
