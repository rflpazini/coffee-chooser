package varieties

import (
	"coffee-choose/internal/router"
	"github.com/labstack/echo/v4"
	"go.uber.org/dig"
)

type coffeeVarietyRoutesParams struct {
	dig.In

	*echo.Echo
	GetAllCoffeeVarietiesHandler  echo.HandlerFunc `name:"Route.Handler.CoffeeVariety.GetAll"`
	GetCoffeeVarietyByNameHandler echo.HandlerFunc `name:"Route.Handler.CoffeeVariety.GetByName"`
	PostCoffeeVarietiesHandler    echo.HandlerFunc `name:"Route.Handler.CoffeeVariety.Post"`
}

func setupCoffeeVarietyRoutes(p coffeeVarietyRoutesParams) router.RouteGroup {
	coffeeVarietyRoutes := p.Echo.Group("/v1")

	coffeeVarietyRoutes.GET("/varieties", p.GetAllCoffeeVarietiesHandler)
	coffeeVarietyRoutes.GET("/varieties/:name", p.GetCoffeeVarietyByNameHandler)

	coffeeVarietyRoutes.POST("/varieties", p.PostCoffeeVarietiesHandler)

	return router.RouteGroup{Group: coffeeVarietyRoutes}
}

func Register(c *dig.Container, register func(...interface{}) error) error {
	if err := c.Provide(makeGetAllCoffeeVarietiesHandler, dig.Name("Route.Handler.CoffeeVariety.GetAll")); err != nil {
		return err
	}
	if err := c.Provide(makeGetCoffeeByNameHandler, dig.Name("Route.Handler.CoffeeVariety.GetByName")); err != nil {
		return err
	}
	if err := c.Provide(makePostCoffeeVarietyHandler, dig.Name("Route.Handler.CoffeeVariety.Post")); err != nil {
		return err
	}

	return register(setupCoffeeVarietyRoutes)
}
