package app

import (
	"context"
	"fmt"

	"github.com/maYkiss56/tunes/internal/config"
	"github.com/maYkiss56/tunes/internal/logger"
	"github.com/maYkiss56/tunes/internal/server"
)

type App struct {
	cfg        *config.Config
	httpServer *server.HTTPServer
	logger     *logger.Logger
}

func New(cfg *config.Config, logger *logger.Logger) (*App, error) {
	httpServer, err := server.NewHTTPServer(cfg, logger)
	if err != nil {
		logger.Error("Failed to create HTTP server", "error", err)
		return nil, fmt.Errorf("http server creation failed: %w", err)
	}

	return &App{
		cfg:        cfg,
		logger:     logger,
		httpServer: httpServer,
	}, nil
}

func (a *App) Run(ctx context.Context) error {
	defer a.logger.Shutdown()

	a.logger.Info(
		"Starting application",
	)

	serverErr := make(chan error, 1)

	go a.httpServer.Start(serverErr)

	select {
	case err := <-serverErr:
		a.logger.Error("Server stopped with error", "error", err)
		return err
	case <-ctx.Done():
		a.logger.Info("Received shutdown signal")
		shutdownCtx, cancel := context.WithTimeout(
			context.Background(),
			a.cfg.HTTP.GracefullTimeout,
		)
		defer cancel()

		if err := a.httpServer.Shutdown(shutdownCtx); err != nil {
			a.logger.Error("Greceful shutdown failed", "error", err)
			return fmt.Errorf("shutdown failed: %w", err)
		}
		a.logger.Info("Application stopped gracefully")
		return nil
	}
}
