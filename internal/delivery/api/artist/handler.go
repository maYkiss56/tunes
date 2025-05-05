package artist

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	domain "github.com/maYkiss56/tunes/internal/domain/artist"
	"github.com/maYkiss56/tunes/internal/domain/artist/dto"
	"github.com/maYkiss56/tunes/internal/logger"
	"github.com/maYkiss56/tunes/internal/utilites"
)

type ArtistService interface {
	CreateArtist(ctx context.Context, artist *domain.Artist) error
	GetAllArtists(ctx context.Context) ([]*domain.Artist, error)
	GetArtistByID(ctx context.Context, id int) (*domain.Artist, error)
	UpdateArtist(ctx context.Context, id int, update dto.UpdateArtistRequest) error
	DeleteArtist(ctx context.Context, id int) error
}

type Handler struct {
	service ArtistService
	logger  *logger.Logger
}

func NewHandler(service ArtistService, logger *logger.Logger) *Handler {
	return &Handler{
		service: service,
		logger:  logger,
	}
}

func (h *Handler) CreateArtist(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateArtistRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.logger.Error("invalid request body", "error", err)
		utilites.RenderError(w, r, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := req.Validate(); err != nil {
		utilites.RenderError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	newArtist, err := domain.NewArtist(req.Nickname, req.BIO, req.Country)
	if err != nil {
		h.logger.Error("invalid input artist", "error", err)
		utilites.RenderError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.service.CreateArtist(r.Context(), newArtist); err != nil {
		h.logger.Error("failed to create artist", "error", err)
		utilites.RenderError(w, r, http.StatusInternalServerError, "failed to create artist")
		return
	}

	utilites.RenderJSON(w, r, http.StatusCreated, dto.ToResponse(*newArtist))
}

func (h *Handler) GetAllArtists(w http.ResponseWriter, r *http.Request) {
	artists, err := h.service.GetAllArtists(r.Context())
	if err != nil {
		h.logger.Error("failed to get artists", "error", err)
		utilites.RenderError(w, r, http.StatusInternalServerError, "failed to get artists")
		return
	}

	artistList := make([]dto.Response, 0, len(artists))
	for _, artist := range artists {
		artistList = append(artistList, dto.ToResponse(*artist))
	}
	utilites.RenderJSON(w, r, http.StatusOK, artistList)
}

func (h *Handler) GetArtistByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		h.logger.Error("invalid artist id", "error", err)
		utilites.RenderError(w, r, http.StatusBadRequest, "invalid artist id")
		return
	}

	a, err := h.service.GetArtistByID(r.Context(), id)
	if err != nil {
		//TODO: not found err
		h.logger.Error("failed to get artist by id", "error", err)
		utilites.RenderError(w, r, http.StatusInternalServerError, "failed to get artist by id")
		return
	}

	utilites.RenderJSON(w, r, http.StatusOK, dto.Response(*a))
}

func (h *Handler) UpdateArtist(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		h.logger.Error("invalid artist id", "error", err)
		utilites.RenderError(w, r, http.StatusBadRequest, "invalid artist id")
		return
	}

	var req dto.UpdateArtistRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.logger.Error("invalid request body", "error", err)
		utilites.RenderError(w, r, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := req.Validate(); err != nil {
		utilites.RenderError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.service.UpdateArtist(r.Context(), id, req); err != nil {
		//TODO: not found err
		h.logger.Error("failed to update artist", "error", err)
		utilites.RenderError(w, r, http.StatusInternalServerError, "failed to update artist")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) DeleteArtist(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		h.logger.Error("invalid artist id", "error", err)
		utilites.RenderError(w, r, http.StatusBadRequest, "invalid artist id")
		return
	}

	if err := h.service.DeleteArtist(r.Context(), id); err != nil {
		//TODO: err not found
		h.logger.Error("failed to delete artist", "error", err)
		utilites.RenderError(w, r, http.StatusInternalServerError, "failed to delete artist")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
