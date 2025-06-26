package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	albumHandler "github.com/maYkiss56/tunes/internal/delivery/api/album"
	artistHandler "github.com/maYkiss56/tunes/internal/delivery/api/artist"
	genreHandler "github.com/maYkiss56/tunes/internal/delivery/api/genre"
	reviewHandler "github.com/maYkiss56/tunes/internal/delivery/api/review"
	songHandler "github.com/maYkiss56/tunes/internal/delivery/api/song"
	userHandler "github.com/maYkiss56/tunes/internal/delivery/api/user"
	"github.com/maYkiss56/tunes/internal/logger"
	"github.com/maYkiss56/tunes/internal/middleware"
)

func NewRouter(
	user *userHandler.Handler,
	song *songHandler.Handler,
	artist *artistHandler.Handler,
	album *albumHandler.Handler,
	genre *genreHandler.Handler,
	review *reviewHandler.Handler,
	logger *logger.Logger,
) chi.Router {
	r := chi.NewRouter()

	r.Use(middleware.RecoverMiddleware(logger))

	r.Mount("/uploads/", http.StripPrefix("/uploads/", http.FileServer(http.Dir("static/uploads"))))

	userRouter := chi.NewRouter()
	userHandler.RegisterRoutes(userRouter, user)
	r.Mount("/api", userRouter)

	songRouter := chi.NewRouter()
	songHandler.RegisterPublicRoutes(songRouter, song)
	r.Mount("/api/songs", songRouter)

	songAdminRouter := chi.NewRouter()
	songHandler.RegisterAdminRoutes(songAdminRouter, song)
	r.Mount("/api/admin/songs", songAdminRouter)

	artistRouter := chi.NewRouter()
	artistHandler.RegisterPublicRoutes(artistRouter, artist)
	r.Mount("/api/artists", artistRouter)

	artistAdminRouter := chi.NewRouter()
	artistHandler.RegisterAdminRoutes(artistAdminRouter, artist)
	r.Mount("/api/admin/artists", artistAdminRouter)

	albumRouter := chi.NewRouter()
	albumHandler.RegisterPublicRoutes(albumRouter, album)
	r.Mount("/api/albums", albumRouter)

	albumAdminRouter := chi.NewRouter()
	albumHandler.RegisterAdminRoutes(albumAdminRouter, album)
	r.Mount("/api/admin/albums", albumAdminRouter)

	genreRouter := chi.NewRouter()
	genreHandler.RegisterPublicRoutes(genreRouter, genre)
	r.Mount("/api/genres", genreRouter)

	genreAdminRouter := chi.NewRouter()
	genreHandler.RegisterAdminRoutes(genreAdminRouter, genre)
	r.Mount("/api/admin/genres", genreAdminRouter)

	reviewRouter := chi.NewRouter()
	reviewHandler.RegisterPublicRoutes(reviewRouter, review)
	r.Mount("/api/reviews", reviewRouter)
	return r
}
