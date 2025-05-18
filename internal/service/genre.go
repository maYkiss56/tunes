package service

import (
	"context"

	domain "github.com/maYkiss56/tunes/internal/domain/genre"
	"github.com/maYkiss56/tunes/internal/domain/genre/dto"
	"github.com/maYkiss56/tunes/internal/logger"
)

type GenreRepository interface {
	CreateGenre(ctx context.Context, genre *domain.Genre) error
	GetAllGenre(ctx context.Context) ([]*domain.Genre, error)
	GetGenreByID(ctx context.Context, id int) (*domain.Genre, error)
	UpdateGenre(ctx context.Context, id int, update dto.UpdateGenreRequest) error
	DeleteGenre(ctx context.Context, id int) error
}

type GenreService struct {
	repo   GenreRepository
	logger *logger.Logger
}

func NewGenreService(repo GenreRepository, logger *logger.Logger) *GenreService {
	return &GenreService{
		repo:   repo,
		logger: logger,
	}
}

func (s *GenreService) CreateGenre(ctx context.Context, genre *domain.Genre) error {
	return s.repo.CreateGenre(ctx, genre)
}

func (s *GenreService) GetAllGenre(ctx context.Context) ([]*domain.Genre, error) {
	genres, err := s.repo.GetAllGenre(ctx)
	if err != nil {
		return nil, err
	}

	return genres, nil
}

func (s *GenreService) GetGenreByID(ctx context.Context, id int) (*domain.Genre, error) {
	return s.repo.GetGenreByID(ctx, id)
}

func (s *GenreService) UpdateGenre(
	ctx context.Context,
	id int,
	update dto.UpdateGenreRequest,
) error {
	return s.repo.UpdateGenre(ctx, id, update)
}

func (s *GenreService) DeleteGenre(ctx context.Context, id int) error {
	return s.repo.DeleteGenre(ctx, id)
}
