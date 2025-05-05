package api

import (
	"github.com/go-chi/chi/v5"

	albumHandler "github.com/maYkiss56/tunes/internal/delivery/api/album"
	artistHandler "github.com/maYkiss56/tunes/internal/delivery/api/artist"
	songHandler "github.com/maYkiss56/tunes/internal/delivery/api/song"
	"github.com/maYkiss56/tunes/internal/logger"
	"github.com/maYkiss56/tunes/internal/middleware"
)

func NewRouter(
	song *songHandler.Handler,
	artist *artistHandler.Handler,
	album *albumHandler.Handler,
	logger *logger.Logger,
) chi.Router {
	r := chi.NewRouter()

	r.Use(middleware.RecoverMiddleware(logger))
	songRouter := chi.NewRouter()
	songHandler.RegisterRoutes(songRouter, song)
	r.Mount("/api/songs", songRouter)

	artistRouter := chi.NewRouter()
	artistHandler.RegisterRoutes(artistRouter, artist)
	r.Mount("/api/artists", artistRouter)

	albumRouter := chi.NewRouter()
	albumHandler.RegisterRoutes(albumRouter, album)
	r.Mount("/api/albums", albumRouter)

	return r
}
