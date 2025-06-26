package repository

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	domain "github.com/maYkiss56/tunes/internal/domain/review"
	"github.com/maYkiss56/tunes/internal/domain/review/dto"
	"github.com/maYkiss56/tunes/internal/domain/song"
	"github.com/maYkiss56/tunes/internal/domain/users"
	"github.com/maYkiss56/tunes/internal/logger"
)

type ReviewRepository struct {
	db       *pgxpool.Pool
	logger   *logger.Logger
	userRepo *UserRepository
	songRepo *SongRepository
}

func NewReviewRepository(db *pgxpool.Pool, logger *logger.Logger, userRepo *UserRepository, songRepo *SongRepository) *ReviewRepository {
	return &ReviewRepository{
		db:       db,
		logger:   logger,
		userRepo: userRepo,
		songRepo: songRepo,
	}
}

func (r *ReviewRepository) CreateReview(ctx context.Context, review *domain.Review) error {
	query := `
	insert into review
	(user_id, song_id, body, is_like, is_valid, created_at, updated_at)
	values ($1, $2, $3, $4, $5, $6, $7) returning id`

	err := r.db.QueryRow(
		ctx,
		query,
		review.UserID,
		review.SongID,
		review.Body,
		review.IsLike,
		review.IsValid,
		time.Now(),
		time.Now(),
	).Scan(&review.ID)
	if err != nil {
		r.logger.Error("failed to create review", "error", err)
		return err
	}

	return nil
}

func (r *ReviewRepository) GetAllReviews(ctx context.Context) ([]dto.Response, error) {
	query := `
		select r.id, r.user_id, r.song_id,
		r.body, r.is_like, r.is_valid,
		r.created_at, r.updated_at,
		u.id, u.email, u.username, u.avatar_url,
		s.id, s.title, s.full_title, s.image_url, s.release_date
		from review r
		join users u on r.user_id = u.id
		join song s on r.song_id = s.id`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		r.logger.Error("failed to get all reviews", "error", err)
		return nil, err
	}
	defer rows.Close()

	reviews := make([]dto.Response, 0)

	for rows.Next() {
		var (
			review domain.Review
			user   users.User
			song   song.Song
		)
		err = rows.Scan(
			&review.ID,
			&review.UserID,
			&review.SongID,
			&review.Body,
			&review.IsLike,
			&review.IsValid,
			&review.CreatedAt,
			&review.UpdatedAt,
			&user.ID,
			&user.Email,
			&user.Username,
			&user.AvatarURL,
			&song.ID,
			&song.Title,
			&song.FullTitle,
			&song.ImageURL,
			&song.ReleaseDate,
		)
		if err != nil {
			r.logger.Error("failed to scan rows", "error", err)
			return nil, err
		}

		reviews = append(reviews, dto.ToResponse(review, user, song))
	}

	return reviews, nil
}

func (r *ReviewRepository) GetAllReviewsByUserID(ctx context.Context, id int) ([]dto.Response, error) {
	query := `
        SELECT r.id, r.user_id, r.song_id,
        r.body, r.is_like, r.is_valid,
        r.created_at, r.updated_at,
        u.id, u.email, u.username, u.avatar_url,
        s.id, s.title, s.full_title, s.image_url, s.release_date
        FROM review r
        JOIN users u ON r.user_id = u.id
        JOIN song s ON r.song_id = s.id
        WHERE r.user_id = $1
        ORDER BY r.created_at DESC`

	rows, err := r.db.Query(ctx, query, id)
	if err != nil {
		r.logger.Error("failed to get all reviews for user", "error", err)
		return nil, err
	}
	defer rows.Close()

	reviews := make([]dto.Response, 0)

	for rows.Next() {
		var (
			review domain.Review
			user   users.User
			song   song.Song
		)
		err = rows.Scan(
			&review.ID,
			&review.UserID,
			&review.SongID,
			&review.Body,
			&review.IsLike,
			&review.IsValid,
			&review.CreatedAt,
			&review.UpdatedAt,
			&user.ID,
			&user.Email,
			&user.Username,
			&user.AvatarURL,
			&song.ID,
			&song.Title,
			&song.FullTitle,
			&song.ImageURL,
			&song.ReleaseDate,
		)
		if err != nil {
			r.logger.Error("failed to scan review row", "error", err)
			return nil, err
		}

		reviews = append(reviews, dto.ToResponse(review, user, song))
	}

	return reviews, nil
}

func (r *ReviewRepository) GetReviewByID(ctx context.Context, id int) (*dto.Response, error) {
	query := `
		select r.id, r.user_id, r.song_id,
		r.body, r.is_like, r.is_valid,
		r.created_at, r.updated_at,
		u.id, u.email, u.username, u.avatar_url,
		s.id, s.title, s.full_title, s.image_url, s.release_date
		from review r
		join users u on r.user_id = u.id
		join song s on r.song_id = s.id
		where r.id = $1`

	var (
		review domain.Review
		user   users.User
		song   song.Song
	)

	err := r.db.QueryRow(ctx, query, id).
		Scan(&review.ID, &review.UserID, &review.SongID,
			&review.Body, &review.IsLike, &review.IsValid,
			&review.CreatedAt, &review.UpdatedAt,
			&user.ID, &user.Email, &user.Username, &user.AvatarURL,
			&song.ID, &song.Title, &song.FullTitle, &song.ImageURL, &song.ReleaseDate)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			r.logger.Error("review not found", "id", id)
			return nil, err
		}
		r.logger.Error("failed to search review", "error", err)
		return nil, err
	}

	res := dto.ToResponse(review, user, song)

	return &res, nil
}

func (r *ReviewRepository) UpdateReview(ctx context.Context, id int, update dto.UpdateReviewRequest) error {
	var fields []string
	var args []interface{}
	argPos := 1

	if update.Body != nil {
		fields = append(fields, fmt.Sprintf("body=$%d", argPos))
		args = append(args, *update.Body)
		argPos++
	}
	if update.IsLike != nil {
		fields = append(fields, fmt.Sprintf("is_like=$%d", argPos))
		args = append(args, *update.IsLike)
		argPos++
	}

	if len(fields) == 0 {
		return nil
	}

	fields = append(fields, fmt.Sprintf("updated_at=$%d", argPos))
	args = append(args, time.Now())
	argPos++

	args = append(args, id)
	whereClause := fmt.Sprintf("where id=$%d", argPos)

	query := fmt.Sprintf("update review set %s %s", strings.Join(fields, ", "), whereClause)

	res, err := r.db.Exec(
		ctx,
		query,
		args...,
	)
	if err != nil {
		r.logger.Error("failed to update review", "id", id, "error", err)
		return err
	}

	rowsAffect := res.RowsAffected()
	if rowsAffect == 0 {
		return err
	}

	return nil
}

func (r *ReviewRepository) DeleteReview(ctx context.Context, id int) error {
	query := `delete from review where id=$1`

	if _, err := r.db.Exec(ctx, query, id); err != nil {
		return err
	}

	return nil
}
