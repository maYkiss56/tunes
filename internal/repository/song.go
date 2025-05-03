package repository

import (
	"context"
	"time"

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
			songCreatedAt   time.Duration
			songUpdatedAt   time.Duration
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
	return nil, nil
}

func (r *SongRepository) UpdateSong(
	ctx context.Context,
	id int,
	update domain.UpdateSongRequest,
) error {
	return nil
}

func (r *SongRepository) DeleteSong(ctx context.Context, id int) error {
	return nil
}
