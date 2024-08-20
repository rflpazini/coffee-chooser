package preferences

import (
	"net/http"

	"coffee-choose/pkg/service/geo"
	"coffee-choose/pkg/service/preferences"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"go.uber.org/dig"
)

type makePostParams struct {
	dig.In

	SaveUserPreferences preferences.SaveUserPreferences
	GeoIPService        geo.IPService
}

func makeCreateRequest(p makePostParams) echo.HandlerFunc {
	return func(c echo.Context) error {
		r := c.Request()
		var input preferences.UserPreferences

		if err := c.Bind(&input); err != nil {
			log.Error().Err(err).Msgf("binding request failed: %s", err.Error())
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		if err := c.Validate(&input); err != nil {
			log.Error().Err(err).Msgf("saving request validate failed: %s", err.Error())
			return echo.NewHTTPError(http.StatusBadRequest, "all fields must be filled")
		}

		input.IPAddress = c.RealIP()

		res := make(chan preferences.SaveResponse)
		go func(ch chan preferences.SaveResponse) {
			id, err := p.SaveUserPreferences(r.Context(), input)
			if err != nil {
				log.Error().Err(err).Msgf("error saving to DB: %s", err.Error())
				ch <- preferences.SaveResponse{Err: err.Error()}
				return
			}
			ch <- preferences.SaveResponse{ID: id}
		}(res)

		errResponse := <-res
		if errResponse.Err != "" {
			return echo.NewHTTPError(http.StatusConflict, errResponse.Err)
		}

		close(res)
		return c.JSON(http.StatusCreated, errResponse)
	}
}
