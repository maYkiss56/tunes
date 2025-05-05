package album

import "github.com/go-chi/chi/v5"

func RegisterRoutes(r chi.Router, handler *Handler) {
	r.Route("/", func(r chi.Router) {
		r.Post("/", handler.CreateAlbum)
		r.Get("/", handler.GetAllAlbums)
		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", handler.GetAlbumByID)
			r.Patch("/", handler.UpdateAlbum)
			r.Delete("/", handler.DeleteAlbum)
		})
	})
}
