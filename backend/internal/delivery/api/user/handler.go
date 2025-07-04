package user

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"golang.org/x/crypto/bcrypt"

	domain "github.com/maYkiss56/tunes/internal/domain/users"
	"github.com/maYkiss56/tunes/internal/domain/users/dto"
	"github.com/maYkiss56/tunes/internal/logger"
	"github.com/maYkiss56/tunes/internal/session"
	"github.com/maYkiss56/tunes/internal/utilites"
)

type UserService interface {
	CreateUser(ctx context.Context, user *domain.User) error
	GetUserByEmail(ctx context.Context, email string) (*domain.User, error)
	GetUserByID(ctx context.Context, id int) (*domain.User, error)
	GetTopReviewers(ctx context.Context) ([]dto.TopResponse, error)
	UpdateUserAvatar(ctx context.Context, id int, req dto.UpdateAvatarRequest) error
	UpdateUserPassword(ctx context.Context, id int, req dto.UpdatePasswordRequest) error
	UpdateUserRequest(ctx context.Context, id int, req dto.UpdateUsersRequest) error
}

type Handler struct {
	service UserService
	logger  *logger.Logger
}

func NewHandler(service UserService, logger *logger.Logger) *Handler {
	return &Handler{
		service: service,
		logger:  logger,
	}
}

func (h *Handler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateUsersRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.logger.Error("invalid request body", "error", err)
		utilites.RenderError(w, r, http.StatusBadRequest, "invalid request body")
		return
	}
	defer r.Body.Close()

	if err := req.Validate(); err != nil {
		utilites.RenderError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	newUser, err := domain.NewUser(
		req.Email,
		req.Username,
		req.Password,
	)
	if err != nil {
		h.logger.Error("invalid input request body", "error", err)
		utilites.RenderError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.service.CreateUser(r.Context(), newUser); err != nil {
		h.logger.Error("failed to create user", "error", err)
		utilites.RenderError(w, r, http.StatusInternalServerError, "failed to register user")
		return
	}

	utilites.RenderJSON(w, r, http.StatusCreated, dto.ToResponse(*newUser))
}

func (h *Handler) LoginUser(w http.ResponseWriter, r *http.Request) {
	var req dto.LoginUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.logger.Error("invalid request body", "error", err)
		utilites.RenderError(w, r, http.StatusBadRequest, "invalid request body")
		return
	}
	defer r.Body.Close()

	if err := req.Validate(); err != nil {
		utilites.RenderError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.service.GetUserByEmail(r.Context(), req.Email)
	if err != nil {
		utilites.RenderError(w, r, http.StatusUnauthorized, err.Error())
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password))
	if err != nil {
		utilites.RenderError(w, r, http.StatusUnauthorized, err.Error())
		return
	}

	s, err := session.GenerateSession(r, user, req.RemeberMe)
	if err != nil {
		h.logger.Error("failed to create session", "error", err.Error())
		utilites.RenderError(w, r, http.StatusInternalServerError, err.Error())
		return
	}
	session.SaveSession(s)

	utilites.SetCookie(w, s)

	utilites.RenderJSON(w, r, http.StatusOK, dto.ToResponse(*user))
}

func (h *Handler) ProfileUser(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_id")
	if err != nil {
		h.logger.Error("cookie session not found", "error", err)
		utilites.RenderError(w, r, http.StatusUnauthorized, err.Error())
		return
	}

	s, ok := session.GetSession(cookie.Value)
	if !ok {
		h.logger.Error("invalid or expired session")
		utilites.RenderError(w, r, http.StatusUnauthorized, "invalid or expired session")
		return
	}

	user, err := h.service.GetUserByID(r.Context(), s.UserID)
	if err != nil {
		h.logger.Error("failed to get user", "error", err)
		utilites.RenderError(w, r, http.StatusInternalServerError, "failed to get user")
		return
	}

	utilites.RenderJSON(w, r, http.StatusOK, dto.ToResponse(*user))
}

func (h *Handler) LogoutUser(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_id")
	if err != nil {
		h.logger.Error("cookie session not found", "error", err)
		utilites.RenderError(w, r, http.StatusUnauthorized, err.Error())
		return
	}

	session.DeleteSession(cookie.Value)

	utilites.CleanCookie(w)

	utilites.RenderJSON(
		w,
		r,
		http.StatusOK,
		map[string]string{"message": "successfully logged out"},
	)
}

func (h *Handler) UpdateUserAvatar(w http.ResponseWriter, r *http.Request) {
	s := session.FromContext(r.Context())
	id := s.UserID

	if err := r.ParseMultipartForm(10 << 20); err != nil {
		h.logger.Error("failed to parse form", "error", err)
		utilites.RenderError(w, r, http.StatusInternalServerError, "failed to parse form")
		return
	}

	req := dto.UpdateAvatarRequest{}

	file, header, err := r.FormFile("avatar")
	if err == nil {
		defer file.Close()
		imagePath, err := utilites.SaveImage(file, header, "static/uploads/avatars")
		if err != nil {
			h.logger.Error("failed to save avatar", "error", err)
			utilites.RenderError(w, r, http.StatusInternalServerError, "failed to save avatar")
			return
		}
		req.AvatarURL = imagePath
	}

	if err := h.service.UpdateUserAvatar(r.Context(), id, req); err != nil {
		h.logger.Error("failed to update avatar", "error", err)
		utilites.RenderError(w, r, http.StatusInternalServerError, "failed to update avatar")
		return
	}

	updatedUser, err := h.service.GetUserByID(r.Context(), id)
	if err != nil {
		h.logger.Error("failed to get updated user", "error", err)
		utilites.RenderError(w, r, http.StatusInternalServerError, "failed to get updated user")
		return
	}

	utilites.RenderJSON(w, r, http.StatusOK, dto.ToResponse(*updatedUser))
}

func (h *Handler) UpdateUserPassword(w http.ResponseWriter, r *http.Request) {
	s := session.FromContext(r.Context())
	id := s.UserID

	var req dto.UpdatePasswordRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.logger.Error("invalid request body", "error", err)
		utilites.RenderError(w, r, http.StatusBadRequest, "invalid request body")
		return
	}
	defer r.Body.Close()

	if err := h.service.UpdateUserPassword(r.Context(), id, req); err != nil {
		h.logger.Error("failed to update user password", "error", err)
		utilites.RenderError(w, r, http.StatusInternalServerError, "failed to update password")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) UpdateUserInfo(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		h.logger.Error("invalid user id", "error", err)
		utilites.RenderError(w, r, http.StatusBadRequest, "invalid user id")
		return
	}

	var req dto.UpdateUsersRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.logger.Error("invalid request body", "error", err)
		utilites.RenderError(w, r, http.StatusBadRequest, "invalid request body")
		return
	}
	defer r.Body.Close()

	if err := h.service.UpdateUserRequest(r.Context(), id, req); err != nil {
		h.logger.Error("failed to update user password", "error", err)
		utilites.RenderError(w, r, http.StatusInternalServerError, "failed to update password")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) GetTopReviewers(w http.ResponseWriter, r *http.Request) {
	reviewers, err := h.service.GetTopReviewers(r.Context())
	if err != nil {
		h.logger.Error("failed to get reviewers", "error", err)
		utilites.RenderError(w, r, http.StatusInternalServerError, "failed to get reviewers")
		return
	}

	reviewersList := make([]dto.TopResponse, 0, len(reviewers))
	for _, reviewer := range reviewers {
		reviewersList = append(reviewersList, reviewer)
	}
	utilites.RenderJSON(w, r, http.StatusOK, reviewersList)
}
