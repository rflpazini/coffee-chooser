package router

import (
	"regexp"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"go.uber.org/dig"
)

type newEchoRouterParams struct {
	dig.In
}

func newEchoRouter(p newEchoRouterParams) *echo.Echo {
	router := echo.New()
	router.JSONSerializer = easyJSONSerializer{}

	router.Logger.SetLevel(log.OFF)
	router.Debug = false
	router.HideBanner = true
	router.Validator = &CustomValidator{Validator: validator.New()}

	// Create a regex to match any subdomain of chooser.cafe
	chooserCafeRegex := regexp.MustCompile(`^https?:\/\/([a-zA-Z0-9-]+\.)*chooser\.cafe$`)

	// CORS middleware with custom AllowOriginFunc
	router.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOriginFunc: func(origin string) (bool, error) {
			return chooserCafeRegex.MatchString(origin) || origin == "http://localhost:3000", nil
		},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}))

	return router
}
