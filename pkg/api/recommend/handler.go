package recommend

import (
	"net/http"
	"strings"
	"time"

	"coffee-choose/pkg/service/recommend"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"go.uber.org/dig"
)

type makePostParams struct {
	dig.In

	SaveUserPreferences recommend.SaveUserPreferences
}

func makeCreateRequest(p makePostParams) echo.HandlerFunc {
	return func(c echo.Context) error {
		r := c.Request()
		var input recommend.UserPreferences

		if err := c.Bind(&input); err != nil {
			log.Error().Err(err).Msgf("binding request failed: %s", err.Error())
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		if err := c.Validate(&input); err != nil {
			log.Error().Err(err).Msgf("saving request validate failed: %s", err.Error())
			return echo.NewHTTPError(http.StatusBadRequest, "all fields must be filled")
		}

		input.FlavorNotes = strings.ToLower(input.FlavorNotes)
		input.UpdatedAt = time.Now()
		input.CreatedAt = time.Now()

		res := make(chan recommend.SaveResponse)
		go func(ch chan recommend.SaveResponse) {
			id, err := p.SaveUserPreferences(r.Context(), input)
			if err != nil {
				log.Error().Err(err).Msgf("error saving to DB: %s", err.Error())
				ch <- recommend.SaveResponse{Err: err.Error()}
				return
			}
			ch <- recommend.SaveResponse{ID: id}
		}(res)

		errResponse := <-res
		if errResponse.Err != "" {
			return echo.NewHTTPError(http.StatusConflict, errResponse.Err)
		}

		close(res)
		return c.JSON(http.StatusCreated, errResponse)
	}
}
