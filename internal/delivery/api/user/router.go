package user

import (
	"github.com/go-chi/chi/v5"

	"github.com/maYkiss56/tunes/internal/middleware"
)

func RegisterRoutes(r chi.Router, handler *Handler) {
	r.Route("/auth", func(r chi.Router) {
		r.Post("/register", handler.RegisterUser)
		r.Post("/login", handler.LoginUser)
	})

	r.Route("/profile", func(r chi.Router) {
		r.Use(middleware.AuthMiddleware)
		r.Get("/", handler.ProfileUser)
		r.Post("/logout", handler.LogoutUser)
	})
}
