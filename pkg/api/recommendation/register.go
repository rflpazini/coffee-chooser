package recommendation

import (
	"coffee-choose/internal/router"
	"github.com/labstack/echo/v4"
	"go.uber.org/dig"
)

type coffeeRoutesParams struct {
	dig.In

	*echo.Echo
	SaveRecommendationHandler echo.HandlerFunc `name:"Route.Handler.Coffee.Recommendation.Get"`
}

func setupCoffeeRecommendationRoutes(p coffeeRoutesParams) router.RouteGroup {
	coffeeRoutes := p.Echo.Group("/v1")

	coffeeRoutes.GET("/recommendation", p.SaveRecommendationHandler)

	return router.RouteGroup{Group: coffeeRoutes}
}

func Register(c *dig.Container, register func(...interface{}) error) error {
	if err := c.Provide(makeRecommendationGet, dig.Name("Route.Handler.Coffee.Recommendation.Get")); err != nil {
		return err
	}

	return register(setupCoffeeRecommendationRoutes)
}
