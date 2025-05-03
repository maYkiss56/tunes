package service

import (
	"context"

	domain "github.com/maYkiss56/tunes/internal/domain/song"
	"github.com/maYkiss56/tunes/internal/logger"
)

type SongRepository interface {
	CreateSong(ctx context.Context, song *domain.Song) error
	GetAllSongs(ctx context.Context) ([]*domain.Song, error)
	GetSongByID(ctx context.Context, id int) (*domain.Song, error)
	UpdateSong(ctx context.Context, id int, update domain.UpdateSongRequest) error
	DeleteSong(ctx context.Context, id int) error
}

type SongService struct {
	repo   SongRepository
	logger *logger.Logger
}

func NewSongService(repo SongRepository, logger *logger.Logger) *SongService {
	return &SongService{
		repo:   repo,
		logger: logger,
	}
}

func (s *SongService) CreateSong(ctx context.Context, song *domain.Song) error {
	return s.repo.CreateSong(ctx, song)
}

func (s *SongService) GetAllSongs(ctx context.Context) ([]*domain.Song, error) {
	songs, err := s.repo.GetAllSongs(ctx)
	if err != nil {
		return nil, err
	}
	return songs, nil
}

func (s *SongService) GetSongByID(ctx context.Context, id int) (*domain.Song, error) {
	return s.repo.GetSongByID(ctx, id)
}

func (s *SongService) UpdateSong(
	ctx context.Context,
	id int,
	update domain.UpdateSongRequest,
) error {
	return s.repo.UpdateSong(ctx, id, update)
}

func (s *SongService) DeleteSong(ctx context.Context, id int) error {
	return s.repo.DeleteSong(ctx, id)
}
