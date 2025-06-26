package review

import (
	"github.com/go-chi/chi/v5"

	"github.com/maYkiss56/tunes/internal/middleware"
)

func RegisterPublicRoutes(r chi.Router, handler *Handler) {
	r.Get("/", handler.GetAllReviews)
	r.Get("/{id}", handler.GetReviewByID)

	r.Group(func(r chi.Router) {
		r.Use(middleware.AuthMiddleware)

		r.Post("/", handler.CreateReview)
		r.Get("/user/{id}", handler.GetAllReviewsByUserID)
		r.Patch("/{id}", handler.UpdateReview)
		r.Delete("/{id}", handler.DeleteReview)
	})
}
