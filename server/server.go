package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"coffee-choose/pkg/config"
	"coffee-choose/router"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"go.uber.org/dig"
)

const (
	maxHeaderBytes = 1 << 20
)

type makeServerParams struct {
	dig.In

	*echo.Echo
	*config.ServerConfig
	router.RouteGroups
}

type server interface {
	Shutdown(ctx context.Context) error
	ListenAndServe() error
}

func makeServer(p makeServerParams) server {
	httpServer := &http.Server{
		Addr:           "localhost:" + p.Port,
		ReadTimeout:    time.Second * p.ReadTimeout,
		WriteTimeout:   time.Second * p.WriteTimeout,
		MaxHeaderBytes: maxHeaderBytes,
		Handler:        p.Echo,
	}

	log.Info().
		Dict("port", zerolog.Dict().Str("data.addr", httpServer.Addr)).
		Msg("server startup")

	return httpServer
}

type Start func(ctx context.Context, cancel context.CancelFunc) error

func makeStart(s server) Start {
	return func(ctx context.Context, cancel context.CancelFunc) error {
		defer cancel()
		go func() {
			if err := s.ListenAndServe(); err != nil {
				log.Fatal().Msgf("Error starting Server: ", err)
			}
		}()

		quit := make(chan os.Signal, 1)
		signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

		<-quit

		log.Info().Msg("Server Exited Properly")
		if err := s.Shutdown(ctx); err != nil {
			log.Fatal().Msgf("Server Shutdown Failed: ", err)
		}

		return nil
	}
}
