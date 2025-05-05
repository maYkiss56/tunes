package service

import (
	"context"

	domain "github.com/maYkiss56/tunes/internal/domain/album"
	"github.com/maYkiss56/tunes/internal/domain/album/dto"
	"github.com/maYkiss56/tunes/internal/logger"
)

type AlbumRepository interface {
	CreateAlbum(ctx context.Context, album *domain.Album) error
	GetAllAlbums(ctx context.Context) ([]*domain.Album, error)
	GetAlbumByID(ctx context.Context, id int) (*domain.Album, error)
	UpdateAlbum(ctx context.Context, id int, update dto.UpdateAlbumRequest) error
	DeleteAlbum(ctx context.Context, id int) error
}

type AlbumService struct {
	repo   AlbumRepository
	logger *logger.Logger
}

func NewAlbumService(repo AlbumRepository, logger *logger.Logger) *AlbumService {
	return &AlbumService{
		repo:   repo,
		logger: logger,
	}
}

func (s *AlbumService) CreateAlbum(ctx context.Context, album *domain.Album) error {
	return s.repo.CreateAlbum(ctx, album)
}

func (s *AlbumService) GetAllAlbums(ctx context.Context) ([]*domain.Album, error) {
	albums, err := s.repo.GetAllAlbums(ctx)
	if err != nil {
		return nil, err
	}

	return albums, nil
}

func (s *AlbumService) GetAlbumByID(ctx context.Context, id int) (*domain.Album, error) {
	return s.repo.GetAlbumByID(ctx, id)
}

func (s *AlbumService) UpdateAlbum(
	ctx context.Context,
	id int,
	update dto.UpdateAlbumRequest,
) error {
	return s.repo.UpdateAlbum(ctx, id, update)
}

func (s *AlbumService) DeleteAlbum(ctx context.Context, id int) error {
	return s.repo.DeleteAlbum(ctx, id)
}
