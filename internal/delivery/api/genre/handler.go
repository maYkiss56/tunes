package genre

import (
	"context"
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
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		h.logger.Error("failed to parse multipart form", "error", err)
		utilites.RenderError(w, r, http.StatusBadRequest, "failed to parse form")
		return
	}

	title := r.FormValue("title")
	file, header, err := r.FormFile("image")
	if err != nil {
		h.logger.Error("failed to get image file", "error", err)
		utilites.RenderError(w, r, http.StatusBadRequest, "failed to get image")
		return
	}

	defer func() {
		if file != nil {
			file.Close()
		}
	}()

	var imagePath string
	if file != nil {
		imagePath, err = utilites.SaveImage(file, header, "static/uploads/genres")
		if err != nil {
			h.logger.Error("failed to save image", "error", err)
			utilites.RenderError(w, r, http.StatusInternalServerError, "failed to save image")
			return
		}
	}

	newGenre, err := domain.NewGenre(title, imagePath)
	if err != nil {
		h.logger.Error("invalid input genre", "error", err)
		utilites.RenderError(w, r, http.StatusBadRequest, "invalid input genre")
		return
	}

	if err := h.service.CreateGenre(r.Context(), newGenre); err != nil {
		h.logger.Error("failed to create genre", "error", err)
		utilites.RenderError(w, r, http.StatusInternalServerError, "failed to create genre")
		return
	}

	res := dto.ToResponse(*newGenre)
	res.ImageURl = utilites.GetImageURL(imagePath)

	utilites.RenderJSON(w, r, http.StatusCreated, res)
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
		utilites.RenderError(w, r, http.StatusBadRequest, "invalid genre id")
		return
	}

	// max 10 MB
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		h.logger.Error("failed to parse form", "error", err)
		utilites.RenderError(w, r, http.StatusBadRequest, "failed to parse form")
		return
	}

	req := dto.UpdateGenreRequest{}

	if title := r.FormValue("title"); title != "" {
		req.Title = &title
	}

	file, header, err := r.FormFile("image")
	if err == nil {
		defer file.Close()

		imagePath, err := utilites.SaveImage(file, header, "static/uploads/genres")
		if err != nil {
			h.logger.Error("failed to save image", "error", err)
			utilites.RenderError(w, r, http.StatusInternalServerError, "failed to save image")
			return
		}
		req.ImageURl = &imagePath
	}

	if err := h.service.UpdateGenre(r.Context(), id, req); err != nil {
		h.logger.Error("failed to update genre", "error", err)
		utilites.RenderError(w, r, http.StatusInternalServerError, "failed to update genre")
		return
	}

	updatedGenre, err := h.service.GetGenreByID(r.Context(), id)
	if err != nil {
		h.logger.Error("failed to get updated genre", "error", err)
		utilites.RenderError(w, r, http.StatusInternalServerError, "failed to get updated genre")
		return
	}

	utilites.RenderJSON(w, r, http.StatusOK, dto.ToResponse(*updatedGenre))
}

func (h *Handler) DeleteGenre(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		h.logger.Error("invalid genre id", "error", err)
		utilites.RenderError(w, r, http.StatusBadRequest, "invalid genre id")
		return
	}

	if err := h.service.DeleteGenre(r.Context(), id); err != nil {
		h.logger.Error("failed to delete genre", "error", err)
		utilites.RenderError(w, r, http.StatusInternalServerError, "failed to delete genre")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
