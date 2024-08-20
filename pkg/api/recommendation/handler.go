package recommendation

import (
	"net/http"

	"coffee-choose/pkg/service/recommendation"
	"github.com/labstack/echo/v4"
	"go.uber.org/dig"
)

type makeRecommendationPostParams struct {
	dig.In
}

func makeRecommendationGet() echo.HandlerFunc {
	return func(c echo.Context) error {

		r := &recommendation.Response{
			Variety:     "arara",
			Description: "arara is a brazilian coffee variety known for its vibrant fruity and floral profile, with a balanced sweetness and a full-bodied taste.",
		}
		return c.JSON(http.StatusOK, r)
	}
}
