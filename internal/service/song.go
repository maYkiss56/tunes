package service

import (
	"context"

	domain "github.com/maYkiss56/tunes/internal/domain/song"
	"github.com/maYkiss56/tunes/internal/domain/song/dto"
	"github.com/maYkiss56/tunes/internal/logger"
)

type SongRepository interface {
	CreateSong(ctx context.Context, song *domain.Song) error
	GetSongRating(ctx context.Context, songID int) (int, int, int, error)
	GetAllSongsSortedByRating(ctx context.Context) ([]dto.Response, error)
	GetTopSongs(ctx context.Context, timeRange string, limit int) ([]dto.Response, error)
	GetAllSongs(ctx context.Context) ([]dto.Response, error)
	GetSongByID(ctx context.Context, id int) (*dto.Response, error)
	UpdateSongRating(ctx context.Context, songID int) error
	UpdateSong(ctx context.Context, id int, update dto.UpdateSongRequest) error
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

func (s *SongService) GetSongRating(ctx context.Context, songID int) (int, int, int, error) {
	return s.repo.GetSongRating(ctx, songID)
}

func (s *SongService) GetAllSongsSortedByRating(ctx context.Context) ([]dto.Response, error) {
	return s.repo.GetAllSongsSortedByRating(ctx)
}

func (s *SongService) GetTopSongs(ctx context.Context, timeRange string, limit int) ([]dto.Response, error) {
	return s.repo.GetTopSongs(ctx, timeRange, limit)
}

func (s *SongService) GetAllSongs(ctx context.Context) ([]dto.Response, error) {
	songs, err := s.repo.GetAllSongs(ctx)
	if err != nil {
		return nil, err
	}
	return songs, nil
}

func (s *SongService) GetSongByID(ctx context.Context, id int) (*dto.Response, error) {
	return s.repo.GetSongByID(ctx, id)
}

func (s *SongService) UpdateSongRating(ctx context.Context, songID int) error {
	return s.repo.UpdateSongRating(ctx, songID)
}

func (s *SongService) UpdateSong(
	ctx context.Context,
	id int,
	update dto.UpdateSongRequest,
) error {
	return s.repo.UpdateSong(ctx, id, update)
}

func (s *SongService) DeleteSong(ctx context.Context, id int) error {
	return s.repo.DeleteSong(ctx, id)
}
