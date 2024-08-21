package varieties

import (
	"coffee-choose/internal/router"
	"github.com/labstack/echo/v4"
	"go.uber.org/dig"
)

type coffeeVarietyRoutesParams struct {
	dig.In

	*echo.Echo
	GetAllCoffeeVarietiesHandler echo.HandlerFunc `name:"Route.Handler.CoffeeVariety.GetAll"`
}

func setupCoffeeVarietyRoutes(p coffeeVarietyRoutesParams) router.RouteGroup {
	coffeeVarietyRoutes := p.Echo.Group("/v1")

	coffeeVarietyRoutes.GET("/varieties", p.GetAllCoffeeVarietiesHandler)

	return router.RouteGroup{Group: coffeeVarietyRoutes}
}

func Register(c *dig.Container, register func(...interface{}) error) error {
	if err := c.Provide(makeGetAllCoffeeVarietiesHandler, dig.Name("Route.Handler.CoffeeVariety.GetAll")); err != nil {
		return err
	}

	return register(setupCoffeeVarietyRoutes)
}
