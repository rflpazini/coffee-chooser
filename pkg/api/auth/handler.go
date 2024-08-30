package auth

import (
	"context"
	"net/http"

	"coffee-choose/pkg/auth"
	"coffee-choose/pkg/config"
	"coffee-choose/pkg/service/geo"
	"coffee-choose/pkg/utils"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"go.uber.org/dig"
)

type makeSessionTokenHandlerParams struct {
	dig.In

	*config.JwtConfig
	Geo                    geo.IPService
	CreateSessionTokenFunc auth.CreateSessionTokenFunc
}

func makeSessionTokenHandler(p makeSessionTokenHandlerParams) echo.HandlerFunc {
	return func(c echo.Context) error {
		r := c.Request()
		ctx := context.WithValue(r.Context(), "version", "v1-start")
		r = r.WithContext(ctx)

		userID := c.FormValue("user_id")
		if userID == "" {
			userID = utils.CreateUserId()
		}

		geolocation, _ := p.Geo.GetLocation(r.Context(), c.RealIP())
		sessionID := uuid.NewString()

		token, err := p.CreateSessionTokenFunc(r.Context(), userID, geolocation)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "could not create session token")
		}

		rsp := &Response{
			SessionID: sessionID,
			Token:     token,
		}

		return c.JSON(http.StatusOK, rsp)
	}
}
