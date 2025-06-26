package song

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"

	domain "github.com/maYkiss56/tunes/internal/domain/song"
	"github.com/maYkiss56/tunes/internal/domain/song/dto"
	"github.com/maYkiss56/tunes/internal/logger"
	"github.com/maYkiss56/tunes/internal/utilites"
)

type SongService interface {
	CreateSong(ctx context.Context, song *domain.Song) error
	GetAllSongsSortedByRating(ctx context.Context) ([]dto.Response, error)
	GetTopSongs(ctx context.Context, timeRange string, limit int) ([]dto.Response, error)
	GetAllSongs(ctx context.Context) ([]dto.Response, error)
	GetSongByID(ctx context.Context, id int) (*dto.Response, error)
	UpdateSong(ctx context.Context, id int, update dto.UpdateSongRequest) error
	DeleteSong(ctx context.Context, id int) error
}

type Handler struct {
	service SongService
	logger  *logger.Logger
}

func NewHandler(service SongService, logger *logger.Logger) *Handler {
	return &Handler{
		service: service,
		logger:  logger,
	}
}

func (h *Handler) CreateSong(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		h.logger.Error("failed to parse mulipart form", "error", err)
		utilites.RenderError(w, r, http.StatusBadRequest, "failed to parse form")
		return
	}

	title := r.FormValue("title")
	fullTitle := r.FormValue("full_title")
	file, header, err := r.FormFile("image")
	if err != nil {
		h.logger.Error("failed to get image file", "error", err)
		utilites.RenderError(w, r, http.StatusBadRequest, "failed to get image")
		return
	}

	realeseStr := r.FormValue("release_date")

	releaseDate, err := time.Parse(time.RFC3339, realeseStr)
	if err != nil {
		h.logger.Error("failed to parse release date song", "error", err)
		utilites.RenderError(w, r, http.StatusBadRequest, "invalid date")
		return
	}

	genreIDStr := r.FormValue("genre_id")
	genreID, err := strconv.Atoi(genreIDStr)
	if err != nil {
		h.logger.Error("failed to convert genre_id sring -> int", "error", err)
		utilites.RenderError(w, r, http.StatusBadRequest, "failed to get genre_id")
		return
	}

	artistIDStr := r.FormValue("artist_id")
	artistID, err := strconv.Atoi(artistIDStr)
	if err != nil {
		h.logger.Error("failed to convert artist_ID string -> int", "error", err)
		utilites.RenderError(w, r, http.StatusBadRequest, "failed to get artist_id")
		return
	}

	albumIDStr := r.FormValue("album_id")
	albumID, err := strconv.Atoi(albumIDStr)
	if err != nil {
		h.logger.Error("failed to convert artist_ID string -> int", "error", err)
		utilites.RenderError(w, r, http.StatusBadRequest, "failed to get artist_id")
		return
	}

	defer func() {
		if file != nil {
			file.Close()
		}
	}()

	var imagePath string

	if file != nil {
		imagePath, err = utilites.SaveImage(file, header, "static/uploads/songs")
		if err != nil {
			h.logger.Error("failed to save image", "error", err)
			utilites.RenderError(w, r, http.StatusInternalServerError, "failed to save image")
			return
		}
	}

	newSong, err := domain.NewSong(title, fullTitle, imagePath, releaseDate, genreID, artistID, albumID)
	if err != nil {
		h.logger.Error("invalid input song", "error", err)
		utilites.RenderError(w, r, http.StatusBadRequest, "invalid unput song")
		return
	}

	if err := h.service.CreateSong(r.Context(), newSong); err != nil {
		h.logger.Error("faile to create song", "error", err)
		utilites.RenderError(w, r, http.StatusInternalServerError, "failed to create song")
		return
	}

	res := *newSong
	res.ImageURL = utilites.GetImageURL(imagePath)

	utilites.RenderJSON(w, r, http.StatusCreated, res)
}

func (h *Handler) GetAllSongsSortedByRating(w http.ResponseWriter, r *http.Request) {
	songs, err := h.service.GetAllSongsSortedByRating(r.Context())
	if err != nil {
		h.logger.Error("failed to get songs sorted by rating", "error", err)
		utilites.RenderError(w, r, http.StatusInternalServerError, "failed to get songs")
		return
	}

	songList := make([]dto.Response, 0, len(songs))
	for _, song := range songs {
		songList = append(songList, song)
	}
	utilites.RenderJSON(w, r, http.StatusOK, songList)
}

func (h *Handler) GetTopSongs(w http.ResponseWriter, r *http.Request) {
	timeRange := r.URL.Query().Get("time_range")
	if timeRange == "" {
		timeRange = "all"
	}

	limitStr := r.URL.Query().Get("limit")
	limit := 100
	if limitStr != "" {
		var err error
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			h.logger.Error("invalid limit parameter", "error", err)
			utilites.RenderError(w, r, http.StatusBadRequest, "invalid limit parameter")
			return
		}
	}

	songs, err := h.service.GetTopSongs(r.Context(), timeRange, limit)
	if err != nil {
		h.logger.Error("failed to get top songs", "error", err)
		utilites.RenderError(w, r, http.StatusInternalServerError, "failed to get top songs")
		return
	}

	utilites.RenderJSON(w, r, http.StatusOK, songs)
}

func (h *Handler) GetAllSongs(w http.ResponseWriter, r *http.Request) {
	songs, err := h.service.GetAllSongs(r.Context())
	if err != nil {
		h.logger.Error("failed to get songs", "error", err)
		utilites.RenderError(w, r, http.StatusInternalServerError, "failed to get songs")
		return
	}

	songList := make([]dto.Response, 0, len(songs))
	for _, song := range songs {
		songList = append(songList, song)
	}
	utilites.RenderJSON(w, r, http.StatusOK, songList)
}

func (h *Handler) GetSongByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		h.logger.Error("invalid song id", "error", err)
		utilites.RenderError(w, r, http.StatusBadRequest, "invalid song id")
		return
	}

	s, err := h.service.GetSongByID(r.Context(), id)
	if err != nil {
		h.logger.Error("failed to get song", "error", err)
		utilites.RenderError(w, r, http.StatusInternalServerError, "failed to get song")
		return
	}

	utilites.RenderJSON(w, r, http.StatusOK, *s)
}

func (h *Handler) UpdateSong(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		h.logger.Error("invalid song id", "error", err)
		utilites.RenderError(w, r, http.StatusBadRequest, "invalid song id")
		return
	}

	// max 10 MB
	if err = r.ParseMultipartForm(10 << 20); err != nil {
		h.logger.Error("failed to parse form", "error", err)
		utilites.RenderError(w, r, http.StatusBadRequest, "failed to parse form")
		return
	}

	req := dto.UpdateSongRequest{}

	if title := r.FormValue("title"); title != "" {
		req.Title = &title
	}
	if fullTitle := r.FormValue("full_title"); fullTitle != "" {
		req.FullTitle = &fullTitle
	}
	if realeseStr := r.FormValue("release_date"); realeseStr != "" {
		releaseDate, err := time.Parse(time.RFC3339, realeseStr)
		if err != nil {
			h.logger.Error("failed to parse release date song", "error", err)
			utilites.RenderError(w, r, http.StatusBadRequest, "invalid date")
			return
		}
		req.ReleaseDate = &releaseDate
	}

	file, header, err := r.FormFile("image")
	if err == nil {
		defer file.Close()
		imagePath, err := utilites.SaveImage(file, header, "static/uploads/songs")
		if err != nil {
			h.logger.Error("failed to save image", "error", err)
			utilites.RenderError(w, r, http.StatusInternalServerError, "failed to save image")
			return
		}
		req.ImageURL = &imagePath
	}

	if genreIDstr := r.FormValue("genre_id"); genreIDstr != "" {
		genreID, err := strconv.Atoi(genreIDstr)
		if err != nil {
			h.logger.Error("failed to convert genreID string -> int")
			utilites.RenderError(w, r, http.StatusInternalServerError, "failed to get genre_id")
			return
		}
		req.GenreID = &genreID
	}

	if artistIDStr := r.FormValue("artist_id"); artistIDStr != "" {
		artistID, err := strconv.Atoi(artistIDStr)
		if err != nil {
			h.logger.Error("failed to convert artistID string -> int")
			utilites.RenderError(w, r, http.StatusInternalServerError, "failed to get artist_id")
			return
		}
		req.ArtistID = &artistID
	}

	if albumIDStr := r.FormValue("album_id"); albumIDStr != "" {

		albumID, err := strconv.Atoi(albumIDStr)
		if err != nil {
			h.logger.Error("failed to convert albumID string -> int")
			utilites.RenderError(w, r, http.StatusInternalServerError, "failed to get album_id")
			return
		}
		req.AlbumID = &albumID
	}

	if err := h.service.UpdateSong(r.Context(), id, req); err != nil {
		h.logger.Error("failed to update song", "error", err)
		utilites.RenderError(w, r, http.StatusInternalServerError, "failed to update song")
		return
	}

	updatedSong, err := h.service.GetSongByID(r.Context(), id)
	if err != nil {
		h.logger.Error("failed to get updated song", "error", err)
		utilites.RenderError(w, r, http.StatusInternalServerError, "failed to get updated song")
		return
	}

	utilites.RenderJSON(w, r, http.StatusOK, *updatedSong)
}

func (h *Handler) DeleteSong(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		h.logger.Error("invalid song id", "error", err)
		utilites.RenderError(w, r, http.StatusBadRequest, "invalid song id")
		return
	}

	if err := h.service.DeleteSong(r.Context(), id); err != nil {
		h.logger.Error("failed to delete song", "error", err)
		utilites.RenderError(w, r, http.StatusInternalServerError, "failed to delete song")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
