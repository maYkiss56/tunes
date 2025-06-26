package review

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	domain "github.com/maYkiss56/tunes/internal/domain/review"
	"github.com/maYkiss56/tunes/internal/domain/review/dto"
	"github.com/maYkiss56/tunes/internal/logger"
	"github.com/maYkiss56/tunes/internal/session"
	"github.com/maYkiss56/tunes/internal/utilites"
)

type ReviewService interface {
	CreateReview(ctx context.Context, review *domain.Review) error
	GetAllReviews(ctx context.Context) ([]dto.Response, error)
	GetAllReviewsByUserID(ctx context.Context, id int) ([]dto.Response, error)
	GetReviewByID(ctx context.Context, id int) (*dto.Response, error)
	UpdateReview(ctx context.Context, id int, update dto.UpdateReviewRequest) error
	DeleteReview(ctx context.Context, id int) error
}

type Handler struct {
	service ReviewService
	logger  *logger.Logger
}

func NewHandler(service ReviewService, logger *logger.Logger) *Handler {
	return &Handler{
		service: service,
		logger:  logger,
	}
}

func (h *Handler) CreateReview(w http.ResponseWriter, r *http.Request) {
	s := session.FromContext(r.Context())

	var req dto.CreateReviewRequest
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.logger.Error("invalid request body", "error", err)
		utilites.RenderError(w, r, http.StatusBadRequest, "invalid request body")
		return
	}

	req.UserID = s.UserID

	if err := req.Validate(); err != nil {
		utilites.RenderError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	newReview, err := domain.NewReview(req.UserID, req.SongID, req.Body, req.IsLike, req.IsValid)
	if err != nil {
		h.logger.Error("invalid input review", "error", err)
		utilites.RenderError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.service.CreateReview(r.Context(), newReview); err != nil {
		h.logger.Error("failed to create review", "error", err)
		utilites.RenderError(w, r, http.StatusInternalServerError, "failed to create review")
		return
	}

	utilites.RenderJSON(w, r, http.StatusCreated, *newReview)
}

func (h *Handler) GetAllReviews(w http.ResponseWriter, r *http.Request) {
	reviews, err := h.service.GetAllReviews(r.Context())
	if err != nil {
		h.logger.Error("failed to get reviews", "error", err)
		utilites.RenderError(w, r, http.StatusInternalServerError, "failed to get reviews")
		return
	}

	reviewsList := make([]dto.Response, 0, len(reviews))
	for _, review := range reviews {
		reviewsList = append(reviewsList, review)
	}

	utilites.RenderJSON(w, r, http.StatusOK, reviewsList)
}

func (h *Handler) GetAllReviewsByUserID(w http.ResponseWriter, r *http.Request) {
	s := session.FromContext(r.Context())

	reviews, err := h.service.GetAllReviewsByUserID(r.Context(), s.UserID)
	if err != nil {
		h.logger.Error("failed to get reviews", "error", err)
		utilites.RenderError(w, r, http.StatusInternalServerError, "failed to get revires")
		return
	}

	reviewsList := make([]dto.Response, 0, len(reviews))
	for _, review := range reviews {
		reviewsList = append(reviewsList, review)
	}

	utilites.RenderJSON(w, r, http.StatusOK, reviewsList)
}

func (h *Handler) GetReviewByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		h.logger.Error("invalid review id", "error", err)
		utilites.RenderError(w, r, http.StatusBadRequest, "invalid review id")
		return
	}

	review, err := h.service.GetReviewByID(r.Context(), id)
	if err != nil {
		h.logger.Error("failed to get review by id", "error", err)
		utilites.RenderError(w, r, http.StatusInternalServerError, "failed to get review by id")
		return
	}

	utilites.RenderJSON(w, r, http.StatusOK, *review)
}

func (h *Handler) UpdateReview(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		h.logger.Error("invalid review id", "error", err)
		utilites.RenderError(w, r, http.StatusBadRequest, "invalid review id")
		return
	}

	var req dto.UpdateReviewRequest
	defer r.Body.Close()
	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.logger.Error("invalid request body", "error", err)
		utilites.RenderError(w, r, http.StatusBadRequest, "invalid request body")
		return
	}

	if err = h.service.UpdateReview(r.Context(), id, req); err != nil {
		h.logger.Error("failed to update review", "error", err)
		utilites.RenderError(w, r, http.StatusInternalServerError, "failed to update review")
		return
	}

	updatedReview, err := h.service.GetReviewByID(r.Context(), id)
	if err != nil {
		h.logger.Error("failed to get review", "error", err)
		utilites.RenderError(w, r, http.StatusInternalServerError, "failed to get review")
		return
	}

	utilites.RenderJSON(w, r, http.StatusOK, *updatedReview)
}

func (h *Handler) DeleteReview(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		h.logger.Error("invalid review id", "error", err)
		utilites.RenderError(w, r, http.StatusBadRequest, "invalid review id")
		return
	}

	if err := h.service.DeleteReview(r.Context(), id); err != nil {
		h.logger.Error("failed to delete review", "error", err)
		utilites.RenderError(w, r, http.StatusInternalServerError, "failed to delete review")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
