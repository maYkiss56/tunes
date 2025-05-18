package genre

import (
	"github.com/go-chi/chi/v5"

	"github.com/maYkiss56/tunes/internal/middleware"
)

func RegisterPublicRoutes(r chi.Router, handler *Handler) {
	r.Route("/", func(r chi.Router) {
		r.Get("/", handler.GetAllGenre)
		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", handler.GetGenreByID)
		})
	})
}

func RegisterAdminRoutes(r chi.Router, handler *Handler) {
	r.Route("/", func(r chi.Router) {
		r.Use(middleware.AuthMiddleware)
		r.Use(middleware.AdminOnlyMiddleware)

		r.Post("/", handler.CreateGenre)
		r.Route("/{id}", func(r chi.Router) {
			r.Patch("/", handler.UpdateGenre)
			r.Delete("/", handler.DeleteGenre)
		})
	})
}
