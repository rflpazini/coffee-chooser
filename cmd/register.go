package main

import (
	"context"

	"coffee-choose/internal/router"
	"coffee-choose/internal/server"
	"coffee-choose/pkg/api/health-check"
	"coffee-choose/pkg/api/recommend"
	"coffee-choose/pkg/config"
	"coffee-choose/pkg/database"
	recommendService "coffee-choose/pkg/service/recommend"
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

	if err := recommendService.Register(register); err != nil {
		return err
	}

	// APIs registration
	if err := healthcheck.Register(c, register); err != nil {
		return err
	}
	if err := recommend.Register(c, register); err != nil {
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
