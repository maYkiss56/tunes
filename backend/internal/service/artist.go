package service

import (
	"context"

	domain "github.com/maYkiss56/tunes/internal/domain/artist"
	"github.com/maYkiss56/tunes/internal/domain/artist/dto"
	"github.com/maYkiss56/tunes/internal/logger"
)

type ArtistRepository interface {
	CreateArtist(ctx context.Context, artist *domain.Artist) error
	GetAllArtists(ctx context.Context) ([]*domain.Artist, error)
	GetArtistByID(ctx context.Context, id int) (*domain.Artist, error)
	UpdateArtist(ctx context.Context, id int, update dto.UpdateArtistRequest) error
	DeleteArtist(ctx context.Context, id int) error
}

type ArtistService struct {
	repo   ArtistRepository
	logger *logger.Logger
}

func NewArtistService(repo ArtistRepository, logger *logger.Logger) *ArtistService {
	return &ArtistService{
		repo:   repo,
		logger: logger,
	}
}

func (s *ArtistService) CreateArtist(ctx context.Context, artist *domain.Artist) error {
	return s.repo.CreateArtist(ctx, artist)
}

func (s *ArtistService) GetAllArtists(ctx context.Context) ([]*domain.Artist, error) {
	artists, err := s.repo.GetAllArtists(ctx)
	if err != nil {
		return nil, err
	}

	return artists, nil
}

func (s *ArtistService) GetArtistByID(ctx context.Context, id int) (*domain.Artist, error) {
	return s.repo.GetArtistByID(ctx, id)
}

func (s *ArtistService) UpdateArtist(
	ctx context.Context,
	id int,
	update dto.UpdateArtistRequest,
) error {
	return s.repo.UpdateArtist(ctx, id, update)
}

func (s *ArtistService) DeleteArtist(ctx context.Context, id int) error {
	return s.repo.DeleteArtist(ctx, id)
}
