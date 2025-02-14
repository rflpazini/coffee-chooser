package main

import (
	"context"

	"coffee-choose/internal/router"
	"coffee-choose/internal/server"
	auth2 "coffee-choose/pkg/api/auth"
	"coffee-choose/pkg/api/health-check"
	"coffee-choose/pkg/api/preferences"
	"coffee-choose/pkg/api/recommendation"
	"coffee-choose/pkg/api/varieties"
	"coffee-choose/pkg/auth"
	"coffee-choose/pkg/config"
	"coffee-choose/pkg/database"
	"coffee-choose/pkg/service/coffeeTypes"
	"coffee-choose/pkg/service/geo"
	opService "coffee-choose/pkg/service/openaiClient"
	preferencesService "coffee-choose/pkg/service/preferences"
	"coffee-choose/pkg/service/recommendations"

	"go.uber.org/dig"
)

func registration(ctx context.Context, c *dig.Container, cfg *config.Config) error {
	register := registerContainers(c)

	if err := c.Provide(auth2.Middleware, dig.Name("Route.Handler.Auth")); err != nil {
		return err
	}

	if err := cfg.Register(register); err != nil {
		return err
	}

	if err := database.Register(register); err != nil {
		return err
	}

	if err := router.Register(register); err != nil {
		return err
	}

	if err := server.Register(register); err != nil {
		return err
	}

	if err := auth.Register(register); err != nil {
		return err
	}

	// Services
	if err := preferencesService.Register(register); err != nil {
		return err
	}

	if err := geo.Register(register); err != nil {
		return err
	}

	if err := coffeeTypes.Register(register); err != nil {
		return err
	}

	if err := opService.Register(register); err != nil {
		return err
	}

	if err := recommendations.Register(register); err != nil {
		return err
	}

	// APIs registration
	if err := healthcheck.Register(c, register); err != nil {
		return err
	}
	if err := preferences.Register(c, register); err != nil {
		return err
	}
	if err := recommendation.Register(c, register); err != nil {
		return err
	}
	if err := varieties.Register(c, register); err != nil {
		return err
	}

	if err := auth2.Register(c, register); err != nil {
		return err
	}

	return nil
}

func registerContainers(c *dig.Container) func(...interface{}) error {
	return func(cts ...interface{}) error {
		for _, ct := range cts {
			if err := c.Provide(ct); err != nil {
				return err
			}
		}
		return nil
	}
}
