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
	"github.com/maYkiss56/tunes/internal/domain/artist"
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

func (r *AlbumRepository) GetAllAlbums(ctx context.Context) ([]dto.Response, error) {
	query := `
		select a.id, a.title,
		a.image_url, a.artist_id,
		ar.id, ar.nickname, ar.bio, ar.country
		from album a
		join artist ar on a.artist_id = ar.id`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		r.logger.Error("failed to get all albums", "error", err)
		return nil, err
	}
	defer rows.Close()

	albums := make([]dto.Response, 0)

	for rows.Next() {
		var (
			album  domain.Album
			artist artist.Artist
		)
		err = rows.Scan(
			&album.ID,
			&album.Title,
			&album.ImageURL,
			&album.ArtistID,
			&artist.ID,
			&artist.Nickname,
			&artist.BIO,
			&artist.Country,
		)
		if err != nil {
			r.logger.Error("failed to scan rows", "error", err)
			return nil, err
		}

		albums = append(albums, dto.ToResponse(album, artist))
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return albums, nil
}

func (r *AlbumRepository) GetAlbumByID(ctx context.Context, id int) (*dto.Response, error) {
	query := `
		select a.id, a.title,
		a.image_url, a.artist_id,
		ar.id, ar.nickname,
		ar.bio, ar.country
		from album a
		join artist ar on a.artist_id = ar.id
		where a.id=$1`

	var (
		album  domain.Album
		artist artist.Artist
	)

	err := r.db.QueryRow(ctx, query, id).
		Scan(&album.ID, &album.Title,
			&album.ImageURL, &album.ArtistID,
			&artist.ID, &artist.Nickname, &artist.BIO,
			&artist.Country)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			r.logger.Error("album not found", "id", id)
			return nil, err
		}
		r.logger.Error("failed to search album", "error", err)
		return nil, err
	}

	res := dto.ToResponse(album, artist)

	return &res, nil
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
