package varieties

import (
	"net/http"

	"coffee-choose/pkg/service/coffeeTypes"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"go.uber.org/dig"
)

type makeCoffeeVarietyHandlerParams struct {
	dig.In

	GetAllCoffeeVarieties coffeeTypes.GetAllCoffeeVarieties
}

func makeGetAllCoffeeVarietiesHandler(p makeCoffeeVarietyHandlerParams) echo.HandlerFunc {
	return func(c echo.Context) error {
		coffeeVarieties, err := p.GetAllCoffeeVarieties(c.Request().Context())
		if err != nil {
			log.Error().Err(err).Msg("error retrieving all coffee varieties from DB")
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, coffeeVarieties)
	}
}
