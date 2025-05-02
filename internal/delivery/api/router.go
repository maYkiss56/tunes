package api

import (
	"github.com/go-chi/chi/v5"

	songHandler "github.com/maYkiss56/tunes/internal/delivery/api/song"
	"github.com/maYkiss56/tunes/internal/logger"
	"github.com/maYkiss56/tunes/internal/middleware"
)

func NewRouter(song *songHandler.Handler, logger *logger.Logger) chi.Router {
	r := chi.NewRouter()

	r.Use(middleware.RecoverMiddleware(logger))
	songRouter := chi.NewRouter()
	songHandler.RegisterRoutes(songRouter, song)
	r.Mount("/api", songRouter)

	return r
}
