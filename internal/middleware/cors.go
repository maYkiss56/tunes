package middleware

import (
	"net/http"

	"github.com/rs/cors"

	"github.com/maYkiss56/tunes/internal/config"
)

func NewCORSHandler(cfg *config.Config, h http.Handler) http.Handler {
	c := cors.New(cors.Options{
		AllowedOrigins:   cfg.HTTP.CORS.AllowedOrigins,
		AllowedMethods:   cfg.HTTP.CORS.AllowedMethods,
		AllowedHeaders:   cfg.HTTP.CORS.AllowedHeaders,
		AllowCredentials: cfg.HTTP.CORS.AllowCredentials,
	})
	return c.Handler(h)
}
