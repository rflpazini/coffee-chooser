package coffee

import (
	"coffee-choose/router"
	"github.com/labstack/echo/v4"
	"go.uber.org/dig"
)

type healthCheckParams struct {
	dig.In

	*echo.Echo
	CoffeeHandler echo.HandlerFunc `name:"Route.Handler.Coffee.Post"`
}

func setupRoutes(p healthCheckParams) router.RouteGroup {
	coffee := p.Echo.Group("/v1")

	coffee.POST("/coffee", p.CoffeeHandler)

	return router.RouteGroup{Group: coffee}
}

func Register(c *dig.Container, register func(...interface{}) error) error {
	if err := c.Provide(makeTestHandler, dig.Name("Route.Handler.Coffee.Post")); err != nil {
		return err
	}
	return register(setupRoutes)
}
