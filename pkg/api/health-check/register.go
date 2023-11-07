package healthcheck

import (
	"net/http"

	"coffee-choose/router"
	"github.com/labstack/echo/v4"
	"go.uber.org/dig"
)

type healthCheckParams struct {
	dig.In

	*echo.Echo
	HealthCheckHandler echo.HandlerFunc `name:"Route.Handler.HealthCheck"`
}

func setupRoutes(p healthCheckParams) router.RouteGroup {
	health := p.Echo.Group("/healthcheck")

	health.GET("/info", p.HealthCheckHandler)
	health.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"ping": "pong"})
	})

	return router.RouteGroup{Group: health}
}

func Register(c *dig.Container, register func(...interface{}) error) error {
	if err := c.Provide(makeHealthCheckHandler, dig.Name("Route.Handler.HealthCheck")); err != nil {
		return err
	}
	return register(setupRoutes)
}
