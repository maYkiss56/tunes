package repository

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	domain "github.com/maYkiss56/tunes/internal/domain/song"
	"github.com/maYkiss56/tunes/internal/logger"
)

type SongRepository struct {
	db     *pgxpool.Pool
	logger *logger.Logger
}

func NewSongRepository(db *pgxpool.Pool, logger *logger.Logger) *SongRepository {
	return &SongRepository{
		db:     db,
		logger: logger,
	}
}

func (r *SongRepository) CreateSong(ctx context.Context, song *domain.Song) error {
	query := `insert into song 
		(title, full_title, image_url, release_date, created_at, updated_at)
		values ($1, $2, $3, $4, $5, $6) returning id`

	err := r.db.QueryRow(
		ctx,
		query,
		song.Title,
		song.FullTitle,
		song.ImageURL,
		song.ReleaseDate,
		time.Now(),
		time.Now(),
	).Scan(&song.ID)

	if err != nil {
		r.logger.Error("failed to create song", "error", err)
		return err
	}

	return nil
}

func (r *SongRepository) GetAllSongs(ctx context.Context) ([]*domain.Song, error) {
	query := `select id, title, full_title,
		image_url, release_date, 
		created_at, updated_at
		from song`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		r.logger.Error("failed to get all songs", "error", err)
		return nil, err
	}
	defer rows.Close()

	songs := make([]*domain.Song, 0)

	for rows.Next() {
		var (
			songID          int
			songTitle       string
			songFullTitle   string
			songImageURL    string
			songReleaseDate time.Time
			songCreatedAt   time.Time
			songUpdatedAt   time.Time
		)
		err = rows.Scan(
			&songID,
			&songTitle,
			&songFullTitle,
			&songImageURL,
			&songReleaseDate,
			&songCreatedAt,
			&songUpdatedAt,
		)
		if err != nil {
			r.logger.Error("failed to scan rows", "error", err)
			return nil, err
		}

		songRow := domain.Song{
			ID:          songID,
			Title:       songTitle,
			FullTitle:   songFullTitle,
			ImageURL:    songImageURL,
			ReleaseDate: songReleaseDate,
			CreatedAt:   songCreatedAt,
			UpdatedAt:   songUpdatedAt,
		}

		songs = append(songs, &songRow)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return songs, nil
}

func (r *SongRepository) GetSongByID(ctx context.Context, id int) (*domain.Song, error) {
	query := `select id, title, full_title, image_url, release_date, created_at, updated_at from song where id=$1`

	var (
		songID          int
		songTitle       string
		songFullTitle   string
		songImageURL    string
		songReleaseDate time.Time
		songCreatedAt   time.Time
		songUpdatedAt   time.Time
	)

	err := r.db.QueryRow(ctx, query, id).
		Scan(&songID, &songTitle, &songFullTitle, &songImageURL, &songReleaseDate, &songCreatedAt, &songUpdatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			r.logger.Error("song not found", "id", id)
			return nil, err
		}
		r.logger.Error("failed to search song", "error", err)
		return nil, err
	}

	return &domain.Song{
		ID:          songID,
		Title:       songTitle,
		FullTitle:   songFullTitle,
		ImageURL:    songImageURL,
		ReleaseDate: songReleaseDate,
		CreatedAt:   songCreatedAt,
		UpdatedAt:   songUpdatedAt,
	}, nil
}

func (r *SongRepository) UpdateSong(
	ctx context.Context,
	id int,
	update domain.UpdateSongRequest,
) error {

	var fields []string
	var args []interface{}
	argPos := 1

	if update.Title != nil {
		fields = append(fields, fmt.Sprintf("title=$%d", argPos))
		args = append(args, *update.Title)
		argPos++
	}

	if update.FullTitle != nil {
		fields = append(fields, fmt.Sprintf("full_title=$%d", argPos))
		args = append(args, *update.FullTitle)
		argPos++
	}
	if update.ImageURL != nil {
		fields = append(fields, fmt.Sprintf("image_url=$%d", argPos))
		args = append(args, *update.ImageURL)
		argPos++
	}
	if update.ReleaseDate != nil {
		fields = append(fields, fmt.Sprintf("release_date=$%d", argPos))
		args = append(args, *update.ReleaseDate)
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

	query := fmt.Sprintf("update song set %s %s", strings.Join(fields, ", "), whereClause)

	res, err := r.db.Exec(
		ctx,
		query,
		args...,
	)
	if err != nil {
		r.logger.Error("failed to update song", "id", id, "error", err)
		return err
	}

	rowsAffect := res.RowsAffected()
	if rowsAffect == 0 {
		r.logger.Info("no updated")
		return err
	}

	return nil
}

func (r *SongRepository) DeleteSong(ctx context.Context, id int) error {
	query := `DELETE FROM song WHERE id=$1`

	if _, err := r.db.Exec(ctx, query, id); err != nil {
		return err
	}

	return nil
}
