package repository

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	domain "github.com/maYkiss56/tunes/internal/domain/artist"
	"github.com/maYkiss56/tunes/internal/domain/artist/dto"
	"github.com/maYkiss56/tunes/internal/logger"
)

type ArtistRepository struct {
	db     *pgxpool.Pool
	logger *logger.Logger
}

func NewArtistRepository(db *pgxpool.Pool, logger *logger.Logger) *ArtistRepository {
	return &ArtistRepository{
		db:     db,
		logger: logger,
	}
}

func (r *ArtistRepository) CreateArtist(ctx context.Context, artist *domain.Artist) error {
	query := `insert into artist 
		(nickname, bio, country) 
		values ($1, $2, $3) returning id`

	err := r.db.QueryRow(
		ctx,
		query,
		artist.Nickname,
		artist.BIO,
		artist.Country,
	).Scan(&artist.ID)

	if err != nil {
		r.logger.Error("failed to create artist", "error", err)
		return err
	}
	return nil
}

func (r *ArtistRepository) GetAllArtists(ctx context.Context) ([]*domain.Artist, error) {
	query := `select id, nickname, bio, country from artist`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		r.logger.Error("failed to get all artist", "error", err)
		return nil, err
	}
	defer rows.Close()

	artists := make([]*domain.Artist, 0)

	for rows.Next() {
		var (
			artistID       int
			artistNickname string
			artistBIO      string
			artistCountry  string
		)
		err = rows.Scan(
			&artistID, &artistNickname, &artistBIO, &artistCountry,
		)
		if err != nil {
			r.logger.Error("failed to scan rows", "error", err)
			return nil, err
		}

		artistRow := domain.Artist{
			ID:       artistID,
			Nickname: artistNickname,
			BIO:      artistBIO,
			Country:  artistCountry,
		}

		artists = append(artists, &artistRow)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return artists, nil
}

func (r *ArtistRepository) GetArtistByID(ctx context.Context, id int) (*domain.Artist, error) {
	query := `select id, nickname, bio, country from artist where id=$1`

	var (
		artistID       int
		artistNickname string
		artistBIO      string
		artistCountry  string
	)

	err := r.db.QueryRow(ctx, query, id).
		Scan(&artistID, &artistNickname, &artistBIO, &artistCountry)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			r.logger.Error("artist not found", "id", id)
			return nil, err
		}
		r.logger.Error("failed to search artist", "error", err)
		return nil, err
	}

	return &domain.Artist{
		ID:       artistID,
		Nickname: artistNickname,
		BIO:      artistBIO,
		Country:  artistCountry,
	}, nil
}

func (r *ArtistRepository) UpdateArtist(
	ctx context.Context,
	id int,
	update dto.UpdateArtistRequest,
) error {
	var fields []string
	var args []interface{}
	argPos := 1

	if update.Nickname != nil {
		fields = append(fields, fmt.Sprintf("nickname=$%d", argPos))
		args = append(args, *update.Nickname)
		argPos++
	}
	if update.BIO != nil {
		fields = append(fields, fmt.Sprintf("bio=$%d", argPos))
		args = append(args, *&update.BIO)
		argPos++
	}
	if update.Country != nil {
		fields = append(fields, fmt.Sprintf("country=$%d", argPos))
		args = append(args, *update.Country)
		argPos++
	}

	if len(fields) == 0 {
		return nil
	}
	//TODO: check need field update

	args = append(args, id)
	whereClause := fmt.Sprintf("where id=$%d", argPos)

	query := fmt.Sprintf("update artist set %s %s", strings.Join(fields, ", "), whereClause)

	res, err := r.db.Exec(
		ctx,
		query,
		args...,
	)
	if err != nil {
		r.logger.Error("failed to update artist", "id", id, "error", err)
		return err
	}

	rowsAffect := res.RowsAffected()
	if rowsAffect == 0 {
		r.logger.Info("no updated")
		// nil
		return err
	}
	return nil
}

func (r *ArtistRepository) DeleteArtist(ctx context.Context, id int) error {
	query := `delete from artist where id=$1`

	if _, err := r.db.Exec(ctx, query, id); err != nil {
		return err
	}

	return nil
}
