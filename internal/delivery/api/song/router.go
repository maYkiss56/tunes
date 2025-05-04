package song

import (
	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(r chi.Router, handler *Handler) {
	r.Route("/songs", func(r chi.Router) {
		r.Post("/", handler.CreateSong)
		r.Get("/", handler.GetAllSongs)
		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", handler.GetSongByID)
			r.Patch("/", handler.UpdateSong)
			r.Delete("/", handler.DeleteSong)
		})
	})
}
