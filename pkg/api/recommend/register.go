package recommend

import (
	"coffee-choose/internal/router"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/dig"
)

type coffeeRoutesParams struct {
	dig.In

	*echo.Echo
	SavePreferencesHandler echo.HandlerFunc `name:"Route.Handler.Coffee.Post"`
}

func setupCoffeeRoutes(p coffeeRoutesParams) router.RouteGroup {
	coffeeRoutes := p.Echo.Group("/v1", middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},                    // Allow only requests from this origin
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE}, // Allow specific HTTP methods
	}))

	coffeeRoutes.POST("/preferences", p.SavePreferencesHandler)

	return router.RouteGroup{Group: coffeeRoutes}
}

func Register(c *dig.Container, register func(...interface{}) error) error {
	if err := c.Provide(makeCreateRequest, dig.Name("Route.Handler.Coffee.Post")); err != nil {
		return err
	}

	return register(setupCoffeeRoutes)
}
