package song

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	e "github.com/maYkiss56/tunes/internal/domain/errors"
	domain "github.com/maYkiss56/tunes/internal/domain/song"
	"github.com/maYkiss56/tunes/internal/domain/song/dto"
	"github.com/maYkiss56/tunes/internal/logger"
	"github.com/maYkiss56/tunes/internal/utilites"
)

type SongService interface {
	CreateSong(ctx context.Context, song *domain.Song) error
	GetAllSongs(ctx context.Context) ([]*domain.Song, error)
	GetSongByID(ctx context.Context, id int) (*domain.Song, error)
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
	var req dto.CreateSongRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.logger.Error("invalid request body", "error", err)
		utilites.RenderError(w, r, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := req.Validate(); err != nil {
		utilites.RenderError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	newSong, err := domain.NewSong(req.Title, req.FullTitle, req.ImageURL, req.ReleaseDate)
	if err != nil {
		h.logger.Error("invalid input song", "error", err)
		utilites.RenderError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.service.CreateSong(r.Context(), newSong); err != nil {
		h.logger.Error("failed to create song", "error", err)
		utilites.RenderError(w, r, http.StatusInternalServerError, "failed to create song")
		return
	}

	utilites.RenderJSON(w, r, http.StatusCreated, dto.ToResponse(*newSong))
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
		songList = append(songList, dto.ToResponse(*song))
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
		if errors.Is(err, e.ErrNotFound) {
			utilites.RenderError(w, r, http.StatusNotFound, "song not found")
			return
		}
		h.logger.Error("failed to get song", "error", err)
		utilites.RenderError(w, r, http.StatusInternalServerError, "failed to get song")
		return
	}

	utilites.RenderJSON(w, r, http.StatusOK, dto.ToResponse(*s))
}

func (h *Handler) UpdateSong(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		h.logger.Error("invalid song id", "error", err)
		utilites.RenderError(w, r, http.StatusBadRequest, "invalid song id")
		return
	}

	var req dto.UpdateSongRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.logger.Error("invalid request body", "error", err)
		utilites.RenderError(w, r, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := h.service.UpdateSong(r.Context(), id, req); err != nil {
		if errors.Is(err, e.ErrNotFound) {
			utilites.RenderError(w, r, http.StatusNotFound, "song not found")
			return
		}
		h.logger.Error("failed to update song", "error", err)
		utilites.RenderError(w, r, http.StatusInternalServerError, "failed to update song")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) DeleteSong(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		h.logger.Error("invalid song id", "error", err)
		utilites.RenderError(w, r, http.StatusBadRequest, "invalid song id")
		return
	}

	if err := h.service.DeleteSong(r.Context(), id); err != nil {
		if errors.Is(err, e.ErrNotFound) {
			utilites.RenderError(w, r, http.StatusNotFound, "song not found")
			return
		}
		h.logger.Error("failed to delete song", "error", err)
		utilites.RenderError(w, r, http.StatusInternalServerError, "failed to delete song")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
