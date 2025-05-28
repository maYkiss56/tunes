package album

import (
	"context"
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
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		h.logger.Error("failed to parse mulipart form", "error", err)
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
	artistIDStr := r.FormValue("artist_id")
	artistID, err := strconv.Atoi(artistIDStr)
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
		imagePath, err = utilites.SaveImage(file, header, "static/uploads/albums")
		if err != nil {
			h.logger.Error("failed to save image", "error", err)
			utilites.RenderError(w, r, http.StatusInternalServerError, "failed to save image")
			return
		}
	}

	newAlbum, err := domain.NewAlbum(title, imagePath, artistID)
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

	res := dto.ToResponse(*newAlbum)
	res.ImageURL = utilites.GetImageURL(imagePath)

	utilites.RenderJSON(w, r, http.StatusCreated, res)
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

	// max 10 MB
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		h.logger.Error("failed to parse form", "error", err)
		utilites.RenderError(w, r, http.StatusBadRequest, "failed to parse form")
		return
	}

	req := dto.UpdateAlbumRequest{}

	if title := r.FormValue("title"); title != "" {
		req.Title = &title
	}

	file, header, err := r.FormFile("image")
	if err == nil {
		defer file.Close()

		imagePath, err := utilites.SaveImage(file, header, "static/uploads/albums")
		if err != nil {
			h.logger.Error("failed to save image", "error", err)
			utilites.RenderError(w, r, http.StatusInternalServerError, "failed to save image")
			return
		}
		req.ImageURL = &imagePath
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

	if err := h.service.UpdateAlbum(r.Context(), id, req); err != nil {
		h.logger.Error("failed to update album", "error", err)
		utilites.RenderError(w, r, http.StatusInternalServerError, "failed to update album")
		return
	}

	updatedUlbum, err := h.service.GetAlbumByID(r.Context(), id)
	if err != nil {
		h.logger.Error("failed to get updated album", "error", err)
		utilites.RenderError(w, r, http.StatusInternalServerError, "failed to get updated song")
		return
	}

	utilites.RenderJSON(w, r, http.StatusOK, dto.ToResponse(*updatedUlbum))
}

func (h *Handler) DeleteAlbum(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		h.logger.Error("invalid album id", "error", err)
		utilites.RenderError(w, r, http.StatusBadRequest, "invalid album id")
		return
	}

	if err := h.service.DeleteAlbum(r.Context(), id); err != nil {
		h.logger.Error("failed to delete album", "error", err)
		utilites.RenderError(w, r, http.StatusInternalServerError, "failed to delete album")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
