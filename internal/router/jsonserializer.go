package router

import (
	"github.com/labstack/echo/v4"
	"github.com/mailru/easyjson"
)

type easyJSONSerializer struct{}

func (d easyJSONSerializer) Serialize(c echo.Context, i any, indent string) error {
	m, ok := i.(easyjson.Marshaler)
	if !ok {
		return echo.DefaultJSONSerializer{}.Serialize(c, i, indent)
	}

	_, err := easyjson.MarshalToWriter(m, c.Response())
	return err
}

func (d easyJSONSerializer) Deserialize(c echo.Context, i any) error {
	u, ok := i.(easyjson.Unmarshaler)
	if !ok {
		return echo.DefaultJSONSerializer{}.Deserialize(c, i)
	}
	return easyjson.UnmarshalFromReader(c.Request().Body, u)
}
