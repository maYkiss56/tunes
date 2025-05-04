package song

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	domain "github.com/maYkiss56/tunes/internal/domain/song"
	"github.com/maYkiss56/tunes/internal/logger"
)

type SongService interface {
	CreateSong(ctx context.Context, song *domain.Song) error
	GetAllSongs(ctx context.Context) ([]*domain.Song, error)
	GetSongByID(ctx context.Context, id int) (*domain.Song, error)
	UpdateSong(ctx context.Context, id int, update domain.UpdateSongRequest) error
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
	var req domain.CreateSongRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.logger.Error("invalid request body", "error", err)
		renderError(w, r, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := req.Validate(); err != nil {
		renderError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	newSong := domain.Song{
		Title:       req.Title,
		FullTitle:   req.FullTitle,
		ImageURL:    req.ImageURL,
		ReleaseDate: req.ReleaseDate,
	}

	if err := h.service.CreateSong(r.Context(), &newSong); err != nil {
		h.logger.Error("failed to create song", "error", err)
		renderError(w, r, http.StatusInternalServerError, "failed to create song")
		return
	}

	renderJSON(w, r, http.StatusCreated, domain.ToResponse(newSong))
}

func (h *Handler) GetAllSongs(w http.ResponseWriter, r *http.Request) {
	songs, err := h.service.GetAllSongs(r.Context())
	if err != nil {
		h.logger.Error("failed to get songs", "error", err)
		renderError(w, r, http.StatusInternalServerError, "failed to get songs")
		return
	}

	var songList []domain.Response
	for _, song := range songs {
		songList = append(songList, domain.ToResponse(*song))
	}
	renderJSON(w, r, http.StatusOK, songList)
}

func (h *Handler) GetSongByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		h.logger.Error("invalid song id", "error", err)
		renderError(w, r, http.StatusBadRequest, "invalid song id")
		return
	}

	s, err := h.service.GetSongByID(r.Context(), id)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			renderError(w, r, http.StatusNotFound, "song not found")
			return
		}
		h.logger.Error("failed to get song", "error", err)
		renderError(w, r, http.StatusInternalServerError, "failed to get song")
		return
	}

	renderJSON(w, r, http.StatusOK, domain.ToResponse(*s))
}

func (h *Handler) UpdateSong(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		h.logger.Error("invalid song id", "error", err)
		renderError(w, r, http.StatusBadRequest, "invalid song id")
		return
	}

	var req domain.UpdateSongRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.logger.Error("invalid request body", "error", err)
		renderError(w, r, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := h.service.UpdateSong(r.Context(), id, req); err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			renderError(w, r, http.StatusNotFound, "song not found")
			return
		}
		h.logger.Error("failed to update song", "error", err)
		renderError(w, r, http.StatusInternalServerError, "failed to update song")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) DeleteSong(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		h.logger.Error("invalid song id", "error", err)
		renderError(w, r, http.StatusBadRequest, "invalid song id")
		return
	}

	if err := h.service.DeleteSong(r.Context(), id); err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			renderError(w, r, http.StatusNotFound, "song not found")
			return
		}
		h.logger.Error("failed to delete song", "error", err)
		renderError(w, r, http.StatusInternalServerError, "failed to delete song")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func renderJSON(w http.ResponseWriter, _ *http.Request, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
}

func renderError(w http.ResponseWriter, r *http.Request, status int, message string) {
	renderJSON(w, r, status, map[string]string{"error": message})
}
