package genre

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	domain "github.com/maYkiss56/tunes/internal/domain/genre"
	"github.com/maYkiss56/tunes/internal/domain/genre/dto"
	"github.com/maYkiss56/tunes/internal/logger"
	"github.com/maYkiss56/tunes/internal/utilites"
)

type GenreService interface {
	CreateGenre(ctx context.Context, genre *domain.Genre) error
	GetAllGenre(ctx context.Context) ([]*domain.Genre, error)
	GetGenreByID(ctx context.Context, id int) (*domain.Genre, error)
	UpdateGenre(ctx context.Context, id int, update dto.UpdateGenreRequest) error
	DeleteGenre(ctx context.Context, id int) error
}

type Handler struct {
	service GenreService
	logger  *logger.Logger
}

func NewHandler(service GenreService, logger *logger.Logger) *Handler {
	return &Handler{
		service: service,
		logger:  logger,
	}
}

func (h *Handler) CreateGenre(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateGenreRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.logger.Error("invalid request body", "error", err)
		utilites.RenderError(w, r, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := req.Validate(); err != nil {
		utilites.RenderError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	newGenre, err := domain.NewGenre(req.Title)
	if err != nil {
		h.logger.Error("invalid input genre", "error", err)
		utilites.RenderError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.service.CreateGenre(r.Context(), newGenre); err != nil {
		h.logger.Error("failed to create album", "error", err)
		utilites.RenderError(w, r, http.StatusInternalServerError, "failed to create genre")
		return
	}

	utilites.RenderJSON(w, r, http.StatusCreated, dto.ToResponse(*newGenre))
}

func (h *Handler) GetAllGenre(w http.ResponseWriter, r *http.Request) {
	genres, err := h.service.GetAllGenre(r.Context())
	if err != nil {
		h.logger.Error("failed to get genres", "error", err)
		utilites.RenderError(w, r, http.StatusInternalServerError, "failed to get genres")
		return
	}

	genresList := make([]dto.Response, 0, len(genres))
	for _, genre := range genres {
		genresList = append(genresList, dto.ToResponse(*genre))
	}

	utilites.RenderJSON(w, r, http.StatusOK, genresList)
}

func (h *Handler) GetGenreByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		h.logger.Error("invalid genre id", "error", err)
		utilites.RenderError(w, r, http.StatusBadRequest, "invalid genre id")
		return
	}

	g, err := h.service.GetGenreByID(r.Context(), id)
	if err != nil {
		//TODO: err not found
		h.logger.Error("failed to get genre by id", "error", err)
		utilites.RenderError(w, r, http.StatusInternalServerError, "failed to get genre by id")
		return
	}

	utilites.RenderJSON(w, r, http.StatusOK, dto.ToResponse(*g))
}

func (h *Handler) UpdateGenre(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		h.logger.Error("invalid genre id", "error", err)
		utilites.RenderError(w, r, http.StatusBadRequest, "ivnalid genre id")
		return
	}

	var req dto.UpdateGenreRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.logger.Error("invalid request body", "error", err)
		utilites.RenderError(w, r, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := req.Validate(); err != nil {
		utilites.RenderError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.service.UpdateGenre(r.Context(), id, req); err != nil {
		//TODO: err not found
		h.logger.Error("failed to update genre", "error", err)
		utilites.RenderError(w, r, http.StatusInternalServerError, "failed to update genre")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) DeleteGenre(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		h.logger.Error("invalid genre id", "error", err)
		utilites.RenderError(w, r, http.StatusBadRequest, "invalid genre id")
		return
	}

	if err := h.service.DeleteGenre(r.Context(), id); err != nil {
		//TODO: err not found
		h.logger.Error("failed to delete genre", "error", err)
		utilites.RenderError(w, r, http.StatusInternalServerError, "failed to delete genre")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
