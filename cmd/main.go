package main

import (
	"context"

	"coffee-choose/internal/server"
	"coffee-choose/pkg/database"
	"coffee-choose/pkg/utils"
	"github.com/rs/zerolog/log"
	"go.uber.org/dig"
)

func main() {
	container := dig.New()

	if err := func() error {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		cfg, err := utils.NewConfig()
		if err != nil {
			log.Error().Err(err).Msg("config load failure")
			return err
		}

		if err := registration(ctx, container, cfg); err != nil {
			log.Error().Err(err).Msg("components registration failure")
			return err
		}

		if err := container.Invoke(func(mongodb database.MongoDB, ping database.Ping, disconnect database.Disconnect) error {
			if err := ping(); err != nil {
				log.Err(err).Msg("failed to ping database")
				panic(err)
			}
			return nil
		}); err != nil {
			log.Err(err).Msg("could not retrieve client")
		}

		if err := container.Invoke(func(start server.Start) error {
			return start(ctx, cancel)
		}); err != nil {
			err := container.Invoke(func(databaseDisconnect database.Disconnect) error {
				err = databaseDisconnect()
				if err != nil {
					log.Error().Err(err).Msg("Failed to disconnect from database")
				}
				return err
			})

			return err
		}

		return nil
	}(); err != nil {
		panic(err)
	}
}
