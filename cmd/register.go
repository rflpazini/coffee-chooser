package main

import (
	"context"

	"coffee-choose/pkg/api/coffee"
	"coffee-choose/pkg/api/health-check"
	"coffee-choose/pkg/config"
	"coffee-choose/pkg/database"
	"coffee-choose/router"
	"coffee-choose/server"
	"go.uber.org/dig"
)

func registration(ctx context.Context, c *dig.Container, cfg *config.Config) error {
	register := registerContainers(c)

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

	// APIs registration
	if err := healthcheck.Register(c, register); err != nil {
		return err
	}
	if err := coffee.Register(c, register); err != nil {
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
