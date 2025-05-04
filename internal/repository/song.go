package repository

import (
	"context"
	"errors"
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
		(id, title, full_title, image_url, release_date, created_at, updated_at)
		values ($1, $2, $3, $4, $5, $6, $7)`

	_, err := r.db.Exec(
		ctx,
		query,
		song.ID,
		song.Title,
		song.FullTitle,
		song.ImageURL,
		song.ReleaseDate,
		song.CreatedAt,
		song.UpdatedAt,
	)
	if err != nil {
		r.logger.Error("failed to create song")
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
		r.logger.Error("failed to get all songs")
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
			r.logger.Error("failed to scan rows")
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

	err := r.db.QueryRow(ctx, query).
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
	query := `update song set title=$1, full_title=$2, image_url=$3, release_date=$4 where id=$5`

	res, err := r.db.Exec(
		ctx,
		query,
		update.Title,
		update.FullTitle,
		update.ImageURL,
		update.ReleaseDate,
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
