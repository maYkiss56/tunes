package app

import (
	"context"
	"fmt"

	"github.com/maYkiss56/tunes/internal/config"
	"github.com/maYkiss56/tunes/internal/delivery/api"
	"github.com/maYkiss56/tunes/internal/delivery/api/artist"
	"github.com/maYkiss56/tunes/internal/delivery/api/song"
	"github.com/maYkiss56/tunes/internal/logger"
	"github.com/maYkiss56/tunes/internal/repository"
	"github.com/maYkiss56/tunes/internal/server"
	"github.com/maYkiss56/tunes/internal/service"
	"github.com/maYkiss56/tunes/pkg/client/postgresql"
)

type App struct {
	cfg        *config.Config
	httpServer *server.HTTPServer
	db         *postgresql.PgClient
	logger     *logger.Logger
}

func New(cfg *config.Config, logger *logger.Logger) (*App, error) {

	pgCfg := postgresql.NewPgConfig(
		cfg.PostgreSQL.Username,
		cfg.PostgreSQL.Password,
		cfg.PostgreSQL.Host,
		cfg.PostgreSQL.Port,
		cfg.PostgreSQL.Database,
		cfg.PostgreSQL.SSLMode,
	)
	logger.Info("init pgConfig")
	dbClient, err := postgresql.NewClient(context.Background(), pgCfg)
	if err != nil {
		logger.Error("Failed to init database client", "error", err)
		return nil, fmt.Errorf("db client init failed: %w", err)
	}

	logger.Info("try get pool")
	pool := dbClient.GetPool()

	artistRepo := repository.NewArtistRepository(pool, logger)
	artistService := service.NewArtistService(artistRepo, logger)
	artistHandler := artist.NewHandler(artistService, logger)

	songRepo := repository.NewSongRepository(pool, logger)
	songService := service.NewSongService(songRepo, logger)
	songHandler := song.NewHandler(songService, logger)

	router := api.NewRouter(songHandler, artistHandler, logger)

	httpServer, err := server.NewHTTPServer(cfg, logger, router)
	if err != nil {
		logger.Error("Failed to create HTTP server", "error", err)
		return nil, fmt.Errorf("http server creation failed: %w", err)
	}

	return &App{
		cfg:        cfg,
		httpServer: httpServer,
		db:         dbClient,
		logger:     logger,
	}, nil
}

func (a *App) Run(ctx context.Context) error {
	defer a.logger.Shutdown()
	defer a.db.Close()

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
