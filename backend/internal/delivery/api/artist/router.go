package artist

import (
	"github.com/go-chi/chi/v5"

	"github.com/maYkiss56/tunes/internal/middleware"
)

func RegisterPublicRoutes(r chi.Router, handler *Handler) {
	r.Route("/", func(r chi.Router) {
		r.Get("/", handler.GetAllArtists)
		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", handler.GetArtistByID)
		})
	})
}

func RegisterAdminRoutes(r chi.Router, handler *Handler) {
	r.Route("/", func(r chi.Router) {
		r.Use(middleware.AuthMiddleware)
		r.Use(middleware.AdminOnlyMiddleware)

		r.Post("/", handler.CreateArtist)
		r.Route("/{id}", func(r chi.Router) {
			r.Patch("/", handler.UpdateArtist)
			r.Delete("/", handler.DeleteArtist)
		})
	})
}
