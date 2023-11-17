package coffee

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"time"

	"coffee-choose/pkg/service/coffee"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/dig"
)

const (
	getBrewingMethod       = "get-brewing-method"
	getBrewingMethodByName = "get-brewing-method-by-name"
	saveBrewingMethod      = "save-brewing-method"
	deleteBrewingMethod    = "delete-brewing-method"
)

type makeGetParams struct {
	dig.In

	coffee.GetBrewingMethod
	coffee.GetBrewingMethodByName
}

func makeGetRequest(p makeGetParams) echo.HandlerFunc {
	return func(c echo.Context) error {
		r := c.Request()
		name := strings.ToLower(c.QueryParam("name"))

		if name != "" {
			r = r.WithContext(context.WithValue(context.Background(), getBrewingMethodByName, name))

			method, err := p.GetBrewingMethodByName(r.Context(), name)
			if errors.Is(err, mongo.ErrNoDocuments) {
				return echo.NewHTTPError(http.StatusNotFound, "Couldn't find brewing method with name: "+name)
			}

			return c.JSON(http.StatusOK, method)
		}

		r = r.WithContext(context.WithValue(context.Background(), getBrewingMethod, nil))
		methods, err := p.GetBrewingMethod(r.Context())
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, methods)
	}
}

type makePostParams struct {
	dig.In

	coffee.SaveBrewingMethod
}

func makeCreateRequest(p makePostParams) echo.HandlerFunc {
	return func(c echo.Context) error {
		r := c.Request()
		var input coffee.BrewingRequest

		if err := c.Bind(&input); err != nil {
			return err
		}
		r = r.WithContext(context.WithValue(context.Background(), saveBrewingMethod, input))

		input.Name = strings.ToLower(input.Name)
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

type makeDeleteParams struct {
	dig.In

	coffee.DeleteBrewingMethod
}

func makeDeleteByNameRequest(p makeDeleteParams) echo.HandlerFunc {
	return func(c echo.Context) error {
		r := c.Request()

		name := strings.ToLower(c.QueryParam("name"))
		if name == "" {
			return c.JSON(http.StatusBadRequest, "name is required")
		}
		r = r.WithContext(context.WithValue(context.Background(), deleteBrewingMethod, name))

		err := p.DeleteBrewingMethod(r.Context(), name)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.NoContent(http.StatusOK)
	}
}
