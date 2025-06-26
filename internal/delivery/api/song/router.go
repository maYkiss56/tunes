package song

import (
	"github.com/go-chi/chi/v5"

	"github.com/maYkiss56/tunes/internal/middleware"
)

func RegisterPublicRoutes(r chi.Router, handler *Handler) {
	r.Route("/", func(r chi.Router) {
		r.Get("/", handler.GetAllSongs)
		r.Get("/sorted-by-rating", handler.GetAllSongsSortedByRating)
		r.Get("/top", handler.GetTopSongs)
		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", handler.GetSongByID)
		})
	})
}

func RegisterAdminRoutes(r chi.Router, handler *Handler) {
	r.Route("/", func(r chi.Router) {
		r.Use(middleware.AuthMiddleware)
		r.Use(middleware.AdminOnlyMiddleware)

		r.Post("/", handler.CreateSong)
		r.Route("/{id}", func(r chi.Router) {
			r.Patch("/", handler.UpdateSong)
			r.Delete("/", handler.DeleteSong)
		})
	})
}
