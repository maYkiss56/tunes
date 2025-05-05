package artist

import "github.com/go-chi/chi/v5"

func RegisterRoutes(r chi.Router, handler *Handler) {
	r.Route("/", func(r chi.Router) {
		r.Post("/", handler.CreateArtist)
		r.Get("/", handler.GetAllArtists)
		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", handler.GetArtistByID)
			r.Patch("/", handler.UpdateArtist)
			r.Delete("/", handler.DeleteArtist)
		})
	})
}
