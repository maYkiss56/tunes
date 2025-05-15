package album

import (
	"github.com/go-chi/chi/v5"

	"github.com/maYkiss56/tunes/internal/middleware"
)

func RegisterPublicRoutes(r chi.Router, handler *Handler) {
	r.Route("/", func(r chi.Router) {
		r.Get("/", handler.GetAllAlbums)
		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", handler.GetAlbumByID)
		})
	})

}

func RegisterAdminRoutes(r chi.Router, handler *Handler) {
	r.Route("/", func(r chi.Router) {
		r.Use(middleware.AuthMiddleware)
		r.Use(middleware.AdminOnlyMiddleware)

		r.Post("/", handler.CreateAlbum)
		r.Route("/{id}", func(r chi.Router) {
			r.Patch("/", handler.UpdateAlbum)
			r.Delete("/", handler.DeleteAlbum)
		})
	})
}
