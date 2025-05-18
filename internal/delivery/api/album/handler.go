package album

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	domain "github.com/maYkiss56/tunes/internal/domain/album"
	"github.com/maYkiss56/tunes/internal/domain/album/dto"
	"github.com/maYkiss56/tunes/internal/logger"
	"github.com/maYkiss56/tunes/internal/utilites"
)

type AlbumService interface {
	CreateAlbum(ctx context.Context, album *domain.Album) error
	GetAllAlbums(ctx context.Context) ([]*domain.Album, error)
	GetAlbumByID(ctx context.Context, id int) (*domain.Album, error)
	UpdateAlbum(ctx context.Context, id int, update dto.UpdateAlbumRequest) error
	DeleteAlbum(ctx context.Context, id int) error
}

type Handler struct {
	service AlbumService
	logger  *logger.Logger
}

func NewHandler(service AlbumService, logger *logger.Logger) *Handler {
	return &Handler{
		service: service,
		logger:  logger,
	}
}

func (h *Handler) CreateAlbum(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateAlbumRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.logger.Error("invalid request body", "error", err)
		utilites.RenderError(w, r, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := req.Validate(); err != nil {
		utilites.RenderError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	newAlbum, err := domain.NewAlbum(req.Title, req.ImageURL, req.ArtistID)
	if err != nil {
		h.logger.Error("invalid input album", "error", err)
		utilites.RenderError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.service.CreateAlbum(r.Context(), newAlbum); err != nil {
		h.logger.Error("failed to create album", "error", err)
		utilites.RenderError(w, r, http.StatusInternalServerError, "failed to create album")
		return
	}

	utilites.RenderJSON(w, r, http.StatusCreated, dto.ToResponse(*newAlbum))
}

func (h *Handler) GetAllAlbums(w http.ResponseWriter, r *http.Request) {
	albums, err := h.service.GetAllAlbums(r.Context())
	if err != nil {
		h.logger.Error("failed to get albums", "error", err)
		utilites.RenderError(w, r, http.StatusInternalServerError, "failed to get albums")
		return
	}

	albumsList := make([]dto.Response, 0, len(albums))
	for _, album := range albums {
		albumsList = append(albumsList, dto.ToResponse(*album))
	}

	utilites.RenderJSON(w, r, http.StatusOK, albumsList)
}

func (h *Handler) GetAlbumByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		h.logger.Error("invalid album id", "error", err)
		utilites.RenderError(w, r, http.StatusBadRequest, "invalid album id")
		return
	}

	a, err := h.service.GetAlbumByID(r.Context(), id)
	if err != nil {
		//TODO: err not found
		h.logger.Error("failed to get album by id", "error", err)
		utilites.RenderError(w, r, http.StatusInternalServerError, "failed to get album by id")
		return
	}

	utilites.RenderJSON(w, r, http.StatusOK, dto.ToResponse(*a))
}

func (h *Handler) UpdateAlbum(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		h.logger.Error("invalid album id", "error", err)
		utilites.RenderError(w, r, http.StatusBadRequest, "invalid album id")
		return
	}

	var req dto.UpdateAlbumRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.logger.Error("invalid request body", "error", err)
		utilites.RenderError(w, r, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := req.Validate(); err != nil {
		utilites.RenderError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.service.UpdateAlbum(r.Context(), id, req); err != nil {
		//TODO: err not found
		h.logger.Error("failed to update album", "error", err)
		utilites.RenderError(w, r, http.StatusInternalServerError, "failed to update album")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) DeleteAlbum(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		h.logger.Error("invalid album id", "error", err)
		utilites.RenderError(w, r, http.StatusBadRequest, "invalid album id")
		return
	}

	if err := h.service.DeleteAlbum(r.Context(), id); err != nil {
		//TODO: err not found
		h.logger.Error("failed to delete album", "error", err)
		utilites.RenderError(w, r, http.StatusInternalServerError, "failed to delete album")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
