package main

import (
	"context"

	"coffee-choose/server"
	"github.com/rs/zerolog/log"
	"go.uber.org/dig"
)

func main() {
	container := dig.New()

	if err := func() error {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		cfg, err := NewConfig()
		if err != nil {
			log.Error().Err(err).Msg("config load failure")
			return err
		}

		if err := registration(ctx, container, cfg); err != nil {
			log.Error().Err(err).Msg("components registration failure")
			return err
		}

		if err := container.Invoke(func(start server.Start) error {
			return start(ctx, cancel)
		}); err != nil {
			log.Error().
				Err(err).
				Msg("httpserver crash")
			return err
		}

		return nil
	}(); err != nil {
		panic(err)
	}

}
