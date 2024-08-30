package preferences

import (
	"coffee-choose/internal/router"
	"github.com/labstack/echo/v4"
	"go.uber.org/dig"
)

type coffeeRoutesParams struct {
	dig.In

	*echo.Echo
	SavePreferencesHandler echo.HandlerFunc    `name:"Route.Handler.Coffee.Post"`
	AuthMiddleware         echo.MiddlewareFunc `name:"Route.Handler.Auth"`
}

func setupCoffeeRoutes(p coffeeRoutesParams) router.RouteGroup {
	coffeeRoutes := p.Echo.Group("/v1")
	coffeeRoutes.Use(p.AuthMiddleware)

	coffeeRoutes.POST("/preferences", p.SavePreferencesHandler)

	return router.RouteGroup{Group: coffeeRoutes}
}

func Register(c *dig.Container, register func(...interface{}) error) error {
	if err := c.Provide(makeCreateRequest, dig.Name("Route.Handler.Coffee.Post")); err != nil {
		return err
	}

	return register(setupCoffeeRoutes)
}
