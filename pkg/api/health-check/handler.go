package healthcheck

import (
	"net/http"

	"coffee-choose/pkg/config"
	"github.com/labstack/echo/v4"
	"go.uber.org/dig"
)

type makeHealthCheckParams struct {
	dig.In

	*config.ServerConfig
}

func makeHealthCheckHandler(p makeHealthCheckParams) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, p.ServerConfig)
	}
}
