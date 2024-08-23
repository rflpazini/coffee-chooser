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

	GetAllCoffeeVarieties  coffeeTypes.GetAllCoffeeVarieties
	GetCoffeeVarietyByName coffeeTypes.GetCoffeeVarietyByName
	AddCoffeeVariety       coffeeTypes.PostCoffeeVariety
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

func makeGetCoffeeByNameHandler(p makeCoffeeVarietyHandlerParams) echo.HandlerFunc {
	return func(c echo.Context) error {
		name := c.Param("name")
		if name == "" {
			return echo.NewHTTPError(http.StatusBadRequest, "Variety name is required")
		}

		coffeeVariety, err := p.GetCoffeeVarietyByName(c.Request().Context(), name)
		if err != nil {
			log.Error().Err(err).Msg("Error retrieving coffee variety by name from DB")
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		if coffeeVariety == nil {
			return echo.NewHTTPError(http.StatusNotFound, "Coffee variety not found")
		}

		return c.JSON(http.StatusOK, coffeeVariety)
	}
}

func makePostCoffeeVarietyHandler(p makeCoffeeVarietyHandlerParams) echo.HandlerFunc {
	return func(c echo.Context) error {
		var input coffeeTypes.CoffeeVariety
		if err := c.Bind(&input); err != nil {
			log.Error().Err(err).Msg("failed to bind input")
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid input")
		}

		if err := c.Validate(&input); err != nil {
			log.Error().Err(err).Msg("validation failed")
			return echo.NewHTTPError(http.StatusBadRequest, "Validation failed")
		}

		updatedId, err := p.AddCoffeeVariety(c.Request().Context(), input)
		if err != nil {
			log.Error().Err(err).Msg("failed to add coffee variety to DB")
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to add coffee variety")
		}

		return c.JSON(http.StatusOK, echo.Map{"id": updatedId})
	}
}
