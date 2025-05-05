package repository

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	domain "github.com/maYkiss56/tunes/internal/domain/album"
	"github.com/maYkiss56/tunes/internal/domain/album/dto"
	"github.com/maYkiss56/tunes/internal/logger"
)

type AlbumRepository struct {
	db     *pgxpool.Pool
	logger *logger.Logger
}

func NewAlbumRepository(db *pgxpool.Pool, logger *logger.Logger) *AlbumRepository {
	return &AlbumRepository{
		db:     db,
		logger: logger,
	}
}

func (r *AlbumRepository) CreateAlbum(ctx context.Context, album *domain.Album) error {
	query := `insert into album
		(title, image_url, artist_id)
		values ($1, $2, $3) returning id`

	err := r.db.QueryRow(
		ctx,
		query,
		album.Title,
		album.ImageURL,
		album.ArtistID,
	).Scan(&album.ID)

	if err != nil {
		r.logger.Error("failed to create song", "error", err)
		return err
	}

	return nil
}

func (r *AlbumRepository) GetAllAlbums(ctx context.Context) ([]*domain.Album, error) {
	query := `select id, title, image_url, artist_id from album`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		r.logger.Error("failed to get all albums", "error", err)
		return nil, err
	}
	defer rows.Close()

	albums := make([]*domain.Album, 0)

	for rows.Next() {
		var (
			albumID       int
			albumTitle    string
			albumImageURL string
			albumArtistID int
		)
		err = rows.Scan(
			&albumID,
			&albumTitle,
			&albumImageURL,
			&albumArtistID,
		)
		if err != nil {
			r.logger.Error("failed to scan rows", "error", err)
			return nil, err
		}

		albumRow := domain.Album{
			ID:       albumID,
			Title:    albumTitle,
			ImageURL: albumImageURL,
			ArtistID: albumArtistID,
		}

		albums = append(albums, &albumRow)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return albums, nil
}

func (r *AlbumRepository) GetAlbumByID(ctx context.Context, id int) (*domain.Album, error) {
	query := `select id, title, image_url, artist_id from album where id=$1`

	var (
		albumID       int
		albumTitle    string
		albumImageURL string
		albumArtistID int
	)

	err := r.db.QueryRow(ctx, query, id).Scan(&albumID, &albumTitle, &albumImageURL, &albumArtistID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			r.logger.Error("album not found", "id", id)
			return nil, err
		}
		r.logger.Error("failed to search album", "error", err)
		return nil, err
	}

	return &domain.Album{
		ID:       albumID,
		Title:    albumTitle,
		ImageURL: albumImageURL,
		ArtistID: albumArtistID,
	}, nil
}

func (r *AlbumRepository) UpdateAlbum(
	ctx context.Context,
	id int,
	update dto.UpdateAlbumRequest,
) error {
	var fields []string
	var args []interface{}
	argPos := 1

	if update.Title != nil {
		fields = append(fields, fmt.Sprintf("title=$%d", argPos))
		args = append(args, *update.Title)
		argPos++
	}

	if update.ImageURL != nil {
		fields = append(fields, fmt.Sprintf("image_url=$%d", argPos))
		args = append(args, *update.ImageURL)
		argPos++
	}

	if update.ArtistID != nil {
		fields = append(fields, fmt.Sprintf("artist_id=$%d", argPos))
		args = append(args, *update.ArtistID)
		argPos++
	}

	if len(fields) == 0 {
		return nil
	}

	args = append(args, id)
	whereClause := fmt.Sprintf("where id=$%d", argPos)

	query := fmt.Sprintf("update album set %s %s", strings.Join(fields, ", "), whereClause)

	res, err := r.db.Exec(
		ctx,
		query,
		args...,
	)
	if err != nil {
		r.logger.Error("failed to update album", "id", id, "error", err)
		return err
	}

	rowsAffect := res.RowsAffected()
	if rowsAffect == 0 {
		r.logger.Info("no updated")
		return err
	}

	return nil

}

func (r *AlbumRepository) DeleteAlbum(ctx context.Context, id int) error {
	query := `delete from album where id=$1`

	if _, err := r.db.Exec(ctx, query, id); err != nil {
		return err
	}

	return nil
}
