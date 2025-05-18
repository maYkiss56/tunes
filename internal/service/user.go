package service

import (
	"context"

	domain "github.com/maYkiss56/tunes/internal/domain/users"
	"github.com/maYkiss56/tunes/internal/domain/users/dto"
	"github.com/maYkiss56/tunes/internal/logger"
	"github.com/maYkiss56/tunes/internal/utilites"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *domain.User) error
	GetUserByEmail(ctx context.Context, email string) (*domain.User, error)
	GetUserByID(ctx context.Context, id int) (*domain.User, error)
	UpdateUserAvatar(ctx context.Context, id int, req dto.UpdateAvatarRequest) error
	UpdateUserPassword(ctx context.Context, id int, req dto.UpdatePasswordRequest) error
	UpdateUserRequest(ctx context.Context, id int, req dto.UpdateUsersRequest) error
}

type UserService struct {
	repo   UserRepository
	logger *logger.Logger
}

func NewUserService(repo UserRepository, logger *logger.Logger) *UserService {
	return &UserService{
		repo:   repo,
		logger: logger,
	}
}

func (s *UserService) CreateUser(ctx context.Context, user *domain.User) error {
	enc, err := utilites.EncryptString(user.PasswordHash)
	if err != nil {
		s.logger.Error("failed to hash password", "error", err)
		return err
	}
	user.SetPasswordHash(enc)

	if err := s.repo.CreateUser(ctx, user); err != nil {
		s.logger.Error("failed to save user to DB", "error", err)
		return err
	}

	return nil
}

func (s *UserService) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	user, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		s.logger.Error("failed to find user", "error", err)
		return nil, err
	}
	return user, nil
}

func (s *UserService) GetUserByID(ctx context.Context, id int) (*domain.User, error) {
	user, err := s.repo.GetUserByID(ctx, id)
	if err != nil {
		s.logger.Error("failed to find user", "error", err)
		return nil, err
	}
	return user, nil
}

func (s *UserService) UpdateUserAvatar(
	ctx context.Context,
	id int,
	req dto.UpdateAvatarRequest,
) error {
	return s.UpdateUserAvatar(ctx, id, req)
}

func (s *UserService) UpdateUserPassword(
	ctx context.Context,
	id int,
	req dto.UpdatePasswordRequest,
) error {
	return s.UpdateUserPassword(ctx, id, req)
}

func (s *UserService) UpdateUserRequest(
	ctx context.Context,
	id int,
	req dto.UpdateUsersRequest,
) error {
	return s.UpdateUserRequest(ctx, id, req)
}
