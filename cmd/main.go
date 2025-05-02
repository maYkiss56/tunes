package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/maYkiss56/tunes/internal/app"
	"github.com/maYkiss56/tunes/internal/config"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	cfg := config.GetConfig()

	application, err := app.New(cfg)
	if err != nil {
		log.Printf("app init failed: %v\n", err)
		os.Exit(1)
	}

	if err := application.Run(ctx); err != nil {
		log.Printf("app run failed: %v\n", err)
		os.Exit(1)
	}
}
