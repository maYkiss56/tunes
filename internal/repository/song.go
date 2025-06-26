package repository

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/maYkiss56/tunes/internal/domain/album"
	"github.com/maYkiss56/tunes/internal/domain/artist"
	"github.com/maYkiss56/tunes/internal/domain/genre"
	domain "github.com/maYkiss56/tunes/internal/domain/song"
	"github.com/maYkiss56/tunes/internal/domain/song/dto"
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
		(title, full_title, image_url, release_date, genre_id, artist_id, album_id, created_at, updated_at)
		values ($1, $2, $3, $4, $5, $6, $7, $8, $9) returning id`

	err := r.db.QueryRow(
		ctx,
		query,
		song.Title,
		song.FullTitle,
		song.ImageURL,
		song.ReleaseDate,
		song.GenreID,
		song.ArtistID,
		song.AlbumID,
		time.Now(),
		time.Now(),
	).Scan(&song.ID)
	if err != nil {
		r.logger.Error("failed to create song", "error", err)
		return err
	}

	return nil
}

func (r *SongRepository) GetSongRating(ctx context.Context, songID int) (int, int, int, error) {
	query := `
		SELECT
      COALESCE(SUM(CASE WHEN is_like = true THEN 1 ELSE 0 END), 0) as like_count,
      COALESCE(SUM(CASE WHEN is_like = false THEN 1 ELSE 0 END), 0) as dislike_count
     FROM review
     WHERE song_id = $1 AND is_valid = true
`

	var likeCount, dislikeCount int
	err := r.db.QueryRow(ctx, query, songID).Scan(&likeCount, &dislikeCount)
	if err != nil {
		return 0, 0, 0, err
	}

	rating := likeCount - dislikeCount
	return likeCount, dislikeCount, rating, nil
}

func (r *SongRepository) GetAllSongsSortedByRating(ctx context.Context) ([]dto.Response, error) {
	query := `
        SELECT s.id, s.title, s.full_title,
        s.image_url, s.release_date,
		    s.like_count, s.dislike_count, s.rating,
		    s.genre_id, s.artist_id, s.album_id,
        s.created_at, s.updated_at,
        g.id, g.title, g.image_url,
        ar.id, ar.nickname, ar.bio, ar.country,
        al.id, al.title, al.image_url, al.artist_id,
        al_ar.id, al_ar.nickname, al_ar.bio, al_ar.country
        FROM song s
        JOIN genre g ON s.genre_id = g.id
        JOIN artist ar ON s.artist_id = ar.id
        JOIN album al ON s.album_id = al.id
        JOIN artist al_ar ON al.artist_id = al_ar.id
        ORDER BY s.rating DESC, s.created_at DESC`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		r.logger.Error("failed to get all songs", "error", err)
		return nil, err
	}
	defer rows.Close()

	songs := make([]dto.Response, 0)

	for rows.Next() {
		var (
			song        domain.Song
			genre       genre.Genre
			songArtist  artist.Artist
			album       album.Album
			albumArtist artist.Artist
		)
		err = rows.Scan(
			&song.ID,
			&song.Title,
			&song.FullTitle,
			&song.ImageURL,
			&song.ReleaseDate,
			&song.LikeCount,
			&song.DislikeCount,
			&song.Rating,
			&song.GenreID,
			&song.ArtistID,
			&song.AlbumID,
			&song.CreatedAt,
			&song.UpdatedAt,
			&genre.ID,
			&genre.Title,
			&genre.ImageURL,
			&songArtist.ID,
			&songArtist.Nickname,
			&songArtist.BIO,
			&songArtist.Country,
			&album.ID,
			&album.Title,
			&album.ImageURL,
			&album.ArtistID,
			&albumArtist.ID,
			&albumArtist.Nickname,
			&albumArtist.BIO,
			&albumArtist.Country,
		)
		if err != nil {
			r.logger.Error("failed to scan rows", "error", err)
			return nil, err
		}

		songs = append(songs, dto.ToResponse(song, genre, songArtist, album, albumArtist))
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return songs, nil
}

// repository/song_repository.go

func (r *SongRepository) GetTopSongs(ctx context.Context, timeRange string, limit int) ([]dto.Response, error) {
	var timeCondition string

	switch timeRange {
	case "week":
		timeCondition = "AND r.created_at >= NOW() - INTERVAL '7 days'"
	case "month":
		timeCondition = "AND r.created_at >= NOW() - INTERVAL '30 days'"
	default:
		timeCondition = ""
	}

	query := `
        SELECT 
            s.id, s.title, s.full_title,
            s.image_url, s.release_date,
            COALESCE(SUM(CASE WHEN r.is_like = true THEN 1 ELSE 0 END), 0) as like_count,
            COALESCE(SUM(CASE WHEN r.is_like = false THEN 1 ELSE 0 END), 0) as dislike_count,
            COALESCE(SUM(CASE WHEN r.is_like = true THEN 1 ELSE 0 END), 0) - 
            COALESCE(SUM(CASE WHEN r.is_like = false THEN 1 ELSE 0 END), 0) as rating,
            s.genre_id, s.artist_id, s.album_id,
            s.created_at, s.updated_at,
            g.id, g.title, g.image_url,
            ar.id, ar.nickname, ar.bio, ar.country,
            al.id, al.title, al.image_url, al.artist_id,
            al_ar.id, al_ar.nickname, al_ar.bio, al_ar.country
        FROM song s
        LEFT JOIN review r ON s.id = r.song_id AND r.is_valid = true ` + timeCondition + `
        JOIN genre g ON s.genre_id = g.id
        JOIN artist ar ON s.artist_id = ar.id
        JOIN album al ON s.album_id = al.id
        JOIN artist al_ar ON al.artist_id = al_ar.id
        GROUP BY s.id, g.id, ar.id, al.id, al_ar.id
        ORDER BY rating DESC, like_count DESC
        LIMIT $1`

	rows, err := r.db.Query(ctx, query, limit)
	if err != nil {
		r.logger.Error("failed to get top songs", "error", err, "timeRange", timeRange)
		return nil, err
	}
	defer rows.Close()

	songs := make([]dto.Response, 0)

	for rows.Next() {
		var (
			song        domain.Song
			genre       genre.Genre
			songArtist  artist.Artist
			album       album.Album
			albumArtist artist.Artist
		)
		err = rows.Scan(
			&song.ID,
			&song.Title,
			&song.FullTitle,
			&song.ImageURL,
			&song.ReleaseDate,
			&song.LikeCount,
			&song.DislikeCount,
			&song.Rating,
			&song.GenreID,
			&song.ArtistID,
			&song.AlbumID,
			&song.CreatedAt,
			&song.UpdatedAt,
			&genre.ID,
			&genre.Title,
			&genre.ImageURL,
			&songArtist.ID,
			&songArtist.Nickname,
			&songArtist.BIO,
			&songArtist.Country,
			&album.ID,
			&album.Title,
			&album.ImageURL,
			&album.ArtistID,
			&albumArtist.ID,
			&albumArtist.Nickname,
			&albumArtist.BIO,
			&albumArtist.Country,
		)
		if err != nil {
			r.logger.Error("failed to scan rows", "error", err)
			return nil, err
		}

		songs = append(songs, dto.ToResponse(song, genre, songArtist, album, albumArtist))
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return songs, nil
}

func (r *SongRepository) GetAllSongs(ctx context.Context) ([]dto.Response, error) {
	query := `
		select s.id, s.title, s.full_title,
		s.image_url, s.release_date, s.like_count,
		s.dislike_count, s.rating, s.genre_id,
		s.artist_id, s.album_id, s.created_at, s.updated_at,
		g.id, g.title, g.image_url,
		ar.id, ar.nickname, ar.bio, ar.country,
		al.id, al.title, al.image_url, al.artist_id,
		al_ar.id, al_ar.nickname, al_ar.bio, al_ar.country
		from song s
		join genre g on s.genre_id = g.id
		join artist ar on s.artist_id = ar.id
		join album al on s.album_id = al.id
		join artist al_ar on al.artist_id = al_ar.id`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		r.logger.Error("failed to get all songs", "error", err)
		return nil, err
	}
	defer rows.Close()

	songs := make([]dto.Response, 0)

	for rows.Next() {
		var (
			song        domain.Song
			genre       genre.Genre
			songArtist  artist.Artist
			album       album.Album
			albumArtist artist.Artist
		)
		err = rows.Scan(
			&song.ID,
			&song.Title,
			&song.FullTitle,
			&song.ImageURL,
			&song.ReleaseDate,
			&song.LikeCount,
			&song.DislikeCount,
			&song.Rating,
			&song.GenreID,
			&song.ArtistID,
			&song.AlbumID,
			&song.CreatedAt,
			&song.UpdatedAt,
			&genre.ID,
			&genre.Title,
			&genre.ImageURL,
			&songArtist.ID,
			&songArtist.Nickname,
			&songArtist.BIO,
			&songArtist.Country,
			&album.ID,
			&album.Title,
			&album.ImageURL,
			&album.ArtistID,
			&albumArtist.ID,
			&albumArtist.Nickname,
			&albumArtist.BIO,
			&albumArtist.Country,
		)
		if err != nil {
			r.logger.Error("failed to scan rows", "error", err)
			return nil, err
		}

		songs = append(songs, dto.ToResponse(song, genre, songArtist, album, albumArtist))
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return songs, nil
}

func (r *SongRepository) GetSongByID(ctx context.Context, id int) (*dto.Response, error) {
	query := `
		select s.id, s.title, s.full_title,
		s.image_url, s.release_date, s.like_count,
		s.dislike_count, s.rating, s.genre_id, s.artist_id,
		s.album_id, s.created_at, s.updated_at,
		g.id, g.title, g.image_url,
		ar.id, ar.nickname, ar.bio, ar.country,
		al.id, al.title, al.image_url, al.artist_id,
		al_ar.id, al_ar.nickname, al_ar.bio, al_ar.country
		from song s
		join genre g on s.genre_id = g.id
		join artist ar on s.artist_id = ar.id
		join album al on s.album_id = al.id
		join artist al_ar on al.artist_id = al_ar.id
		where s.id = $1`

	var (
		song        domain.Song
		genre       genre.Genre
		songArtist  artist.Artist
		album       album.Album
		albumArtist artist.Artist
	)

	err := r.db.QueryRow(ctx, query, id).
		Scan(&song.ID, &song.Title, &song.FullTitle,
			&song.ImageURL, &song.ReleaseDate, &song.LikeCount,
			&song.DislikeCount, &song.Rating, &song.GenreID, &song.ArtistID,
			&song.AlbumID, &song.CreatedAt, &song.UpdatedAt,
			&genre.ID, &genre.Title, &genre.ImageURL,
			&songArtist.ID, &songArtist.Nickname, &songArtist.BIO, &songArtist.Country,
			&album.ID, &album.Title, &album.ImageURL, &album.ArtistID,
			&albumArtist.ID, &albumArtist.Nickname, &albumArtist.BIO, &albumArtist.Country)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			r.logger.Error("song not found", "id", id)
			return nil, err
		}
		r.logger.Error("failed to search song", "error", err)
		return nil, err
	}

	res := dto.ToResponse(song, genre, songArtist, album, albumArtist)

	return &res, nil
}

func (r *SongRepository) UpdateSongRating(ctx context.Context, songID int) error {
	// Сначала проверяем существование песни
	var exists bool
	err := r.db.QueryRow(ctx, "SELECT EXISTS(SELECT 1 FROM song WHERE id = $1)", songID).Scan(&exists)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("song with id %d does not exist", songID)
	}

	likeCount, dislikeCount, rating, err := r.GetSongRating(ctx, songID)
	if err != nil {
		return err
	}

	query := `
		UPDATE song SET
      like_count = $1,
      dislike_count = $2,
      rating = $3,
      updated_at = NOW()
    WHERE id = $4`
	_, err = r.db.Exec(ctx, query, likeCount, dislikeCount, rating, songID)

	return err
}

func (r *SongRepository) UpdateSong(
	ctx context.Context,
	id int,
	update dto.UpdateSongRequest,
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

	if update.GenreID != nil {
		fields = append(fields, fmt.Sprintf("genre_id=$%d", argPos))
		args = append(args, *update.GenreID)
		argPos++
	}

	if update.ArtistID != nil {
		fields = append(fields, fmt.Sprintf("artist_id=$%d", argPos))
		args = append(args, *update.ArtistID)
		argPos++
	}

	if update.AlbumID != nil {
		fields = append(fields, fmt.Sprintf("album_id=$%d", argPos))
		args = append(args, *update.AlbumID)
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
