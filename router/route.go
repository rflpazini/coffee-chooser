package router

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/dig"
)

// RouteGroup di output for aggregating *echo.Group in RouteGroups
type RouteGroup struct {
	dig.Out

	Group *echo.Group `group:"echo-groups"`
}

// RouteGroups aggregate *echo.Group for di injection
type RouteGroups struct {
	dig.In

	Groups []*echo.Group `group:"echo-groups"`
}
