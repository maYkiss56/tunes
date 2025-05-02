package app

import (
	"context"
	"fmt"

	"github.com/maYkiss56/tunes/internal/config"
	"github.com/maYkiss56/tunes/internal/server"
)

type App struct {
	cfg        *config.Config
	httpServer *server.HTTPServer
}

func New(cfg *config.Config) (*App, error) {
	httpServer, err := server.NewHTTPServer(cfg)
	if err != nil {
		return nil, err
	}

	return &App{
		cfg:        cfg,
		httpServer: httpServer,
	}, nil
}

func (a *App) Run(ctx context.Context) error {
	serverErr := make(chan error, 1)

	go a.httpServer.Start(serverErr)

	select {
	case err := <-serverErr:
		return err
	case <-ctx.Done():
		fmt.Println("Context canceled, shutting down")
		shutdownCtx, cancel := context.WithTimeout(
			context.Background(),
			a.cfg.HTTP.GracefullTimeout,
		)
		defer cancel()
		return a.httpServer.Shutdown(shutdownCtx)
	}
}
