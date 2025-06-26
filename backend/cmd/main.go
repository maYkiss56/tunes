package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/maYkiss56/tunes/internal/app"
	"github.com/maYkiss56/tunes/internal/config"
	"github.com/maYkiss56/tunes/internal/logger"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	cfg := config.GetConfig()

	logger, err := logger.New(ctx, "logs/app.log")
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v\n", err)
	}

	application, err := app.New(cfg, logger)
	if err != nil {
		logger.Error("Application init failed", "error", err)
		os.Exit(1)
	}

	if err := application.Run(ctx); err != nil {
		logger.Error("Application run failed", "error", err)
		os.Exit(1)
	}
}
