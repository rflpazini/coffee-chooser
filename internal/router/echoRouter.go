package router

import (
	"github.com/labstack/echo/v4"
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

	return router
}
