package coffee

import (
	"context"
	"net/http"
	"time"

	"coffee-choose/pkg/service/coffee"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"go.uber.org/dig"
)

const (
	getBrewingMethod    = "get-brewing-method"
	saveBrewingMethod   = "save-brewing-method"
	deleteBrewingMethod = "delete-brewing-method"
)

type makePostParams struct {
	dig.In

	coffee.SaveBrewingMethod
}

func makePostHandler(p makePostParams) echo.HandlerFunc {
	return func(c echo.Context) error {
		r := c.Request()
		var input coffee.BrewingRequest

		if err := c.Bind(&input); err != nil {
			return err
		}
		r = r.WithContext(context.WithValue(context.Background(), saveBrewingMethod, input))

		input.UpdatedAt = time.Now()
		go func() {
			err := p.SaveBrewingMethod(r.Context(), input)
			if err != nil {
				log.Err(err).Msgf("error saving to DB: %s", err.Error())
				return
			}
		}()

		return c.NoContent(http.StatusCreated)
	}
}

type makeGetParams struct {
	dig.In

	coffee.GetBrewingMethod
}

func makeGetRequest(p makeGetParams) echo.HandlerFunc {
	return func(c echo.Context) error {
		r := c.Request()
		r = r.WithContext(context.WithValue(context.Background(), getBrewingMethod, nil))

		methods, err := p.GetBrewingMethod(r.Context())
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		if len(methods) < 1 {
			c.Response().WriteHeader(http.StatusNotFound)
		}

		return c.JSON(http.StatusOK, methods)
	}

}

type makeDeleteParams struct {
	dig.In

	coffee.DeleteBrewingMethod
}

func makeDeleteRequest(p makeDeleteParams) echo.HandlerFunc {
	return func(c echo.Context) error {
		r := c.Request()

		name := c.Param("name")
		if name == "" {
			return c.JSON(http.StatusBadRequest, "name is required")
		}
		r = r.WithContext(context.WithValue(context.Background(), deleteBrewingMethod, name))

		err := p.DeleteBrewingMethod(r.Context(), name)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.NoContent(http.StatusOK)
	}
}
