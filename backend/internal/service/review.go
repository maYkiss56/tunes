package service

import (
	"context"

	domain "github.com/maYkiss56/tunes/internal/domain/review"
	"github.com/maYkiss56/tunes/internal/domain/review/dto"
	"github.com/maYkiss56/tunes/internal/logger"
)

type ReviewRepository interface {
	CreateReview(ctx context.Context, review *domain.Review) error
	GetAllReviews(ctx context.Context) ([]dto.Response, error)
	GetAllReviewsByUserID(ctx context.Context, id int) ([]dto.Response, error)
	GetReviewByID(ctx context.Context, id int) (*dto.Response, error)
	UpdateReview(ctx context.Context, id int, update dto.UpdateReviewRequest) error
	DeleteReview(ctx context.Context, id int) error
}

type ReviewService struct {
	repo     ReviewRepository
	songRepo SongRepository
	logger   *logger.Logger
}

func NewReviewService(repo ReviewRepository, songRepo SongRepository, logger *logger.Logger) *ReviewService {
	return &ReviewService{
		repo:     repo,
		songRepo: songRepo,
		logger:   logger,
	}
}

func (s *ReviewService) CreateReview(ctx context.Context, review *domain.Review) error {
	err := s.repo.CreateReview(ctx, review)
	if err != nil {
		return err
	}

	if err := s.songRepo.UpdateSongRating(ctx, review.SongID); err != nil {
		s.logger.Error("Failed to update song rating", "error", err)
		return err
	}

	return nil
}

func (s *ReviewService) GetAllReviews(ctx context.Context) ([]dto.Response, error) {
	reviews, err := s.repo.GetAllReviews(ctx)
	if err != nil {
		return nil, err
	}

	return reviews, nil
}

func (s *ReviewService) GetAllReviewsByUserID(ctx context.Context, id int) ([]dto.Response, error) {
	reviews, err := s.repo.GetAllReviewsByUserID(ctx, id)
	if err != nil {
		return nil, err
	}

	return reviews, nil
}

func (s *ReviewService) GetReviewByID(ctx context.Context, id int) (*dto.Response, error) {
	return s.repo.GetReviewByID(ctx, id)
}

func (s *ReviewService) UpdateReview(
	ctx context.Context,
	id int,
	update dto.UpdateReviewRequest,
) error {
	// Сначала получаем текущую рецензию
	currentReview, err := s.repo.GetReviewByID(ctx, id)
	if err != nil {
		return err
	}

	// Сохраняем предыдущее значение isLike
	oldIsLike := currentReview.IsLike

	// Обновляем рецензию
	if err := s.repo.UpdateReview(ctx, id, update); err != nil {
		return err
	}

	// Если изменился isLike, обновляем рейтинг песни
	if update.IsLike != nil && *update.IsLike != oldIsLike {
		if err := s.songRepo.UpdateSongRating(ctx, currentReview.Song.ID); err != nil {
			s.logger.Error("Failed to update song rating after review update", "error", err)
			return err
		}
	}

	return nil
}

func (s *ReviewService) DeleteReview(ctx context.Context, id int) error {
	// Сначала получаем рецензию, чтобы знать songID
	review, err := s.repo.GetReviewByID(ctx, id)
	if err != nil {
		return err
	}

	// Удаляем рецензию
	if err := s.repo.DeleteReview(ctx, id); err != nil {
		return err
	}

	// Обновляем рейтинг песни
	if err := s.songRepo.UpdateSongRating(ctx, review.Song.ID); err != nil {
		s.logger.Error("Failed to update song rating after review deletion", "error", err)
		return err
	}

	return nil
}
