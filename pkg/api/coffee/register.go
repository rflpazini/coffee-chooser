package coffee

import (
	"coffee-choose/router"
	"github.com/labstack/echo/v4"
	"go.uber.org/dig"
)

type healthCheckParams struct {
	dig.In

	*echo.Echo
	BrewingGetHandler    echo.HandlerFunc `name:"Route.Handler.Brewing.Get"`
	BrewingPostHandler   echo.HandlerFunc `name:"Route.Handler.Brewing.Post"`
	BrewingDeleteHandler echo.HandlerFunc `name:"Route.Handler.Brewing.Delete"`
}

func setupBrewingRoutes(p healthCheckParams) router.RouteGroup {
	brewingRoutes := p.Echo.Group("/v1")

	brewingRoutes.GET("/brewing", p.BrewingGetHandler)
	brewingRoutes.POST("/brewing", p.BrewingPostHandler)
	brewingRoutes.DELETE("/brewing/:name", p.BrewingDeleteHandler)

	return router.RouteGroup{Group: brewingRoutes}
}

func Register(c *dig.Container, register func(...interface{}) error) error {
	if err := c.Provide(makeGetRequest, dig.Name("Route.Handler.Brewing.Get")); err != nil {
		return err
	}

	if err := c.Provide(makePostHandler, dig.Name("Route.Handler.Brewing.Post")); err != nil {
		return err
	}

	if err := c.Provide(makeDeleteRequest, dig.Name("Route.Handler.Brewing.Delete")); err != nil {
		return err
	}

	return register(setupBrewingRoutes)
}
