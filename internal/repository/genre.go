package repository

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	domain "github.com/maYkiss56/tunes/internal/domain/genre"
	"github.com/maYkiss56/tunes/internal/domain/genre/dto"
	"github.com/maYkiss56/tunes/internal/logger"
)

type GenreRepository struct {
	db     *pgxpool.Pool
	logger *logger.Logger
}

func NewGenreRepository(db *pgxpool.Pool, logger *logger.Logger) *GenreRepository {
	return &GenreRepository{
		db:     db,
		logger: logger,
	}
}

func (r *GenreRepository) CreateGenre(ctx context.Context, genre *domain.Genre) error {
	query := `insert ingo genre (title) values ($1) returning id`

	err := r.db.QueryRow(ctx, query, genre.Title).Scan(*&genre.ID)
	if err != nil {
		r.logger.Error("failed to create genre", "error", err)
		return err
	}

	return nil
}

func (r GenreRepository) GetAllGenre(ctx context.Context) ([]*domain.Genre, error) {
	query := `select id, title from genre`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		r.logger.Error("failed to get all genres", "error", err)
		return nil, err
	}
	defer rows.Close()

	genres := make([]*domain.Genre, 0)

	for rows.Next() {
		var (
			genreID    int
			genreTitle string
		)
		err = rows.Scan(&genreID, &genreTitle)
		if err != nil {
			r.logger.Error("failed to scan rows", "error", err)
			return nil, err
		}

		genreRow := domain.Genre{
			ID:    genreID,
			Title: genreTitle,
		}

		genres = append(genres, &genreRow)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return genres, nil
}

func (r *GenreRepository) GetGenreByID(ctx context.Context, id int) (*domain.Genre, error) {
	query := `select id, title from genre where id=$1`

	var (
		genreID    int
		genreTitle string
	)

	err := r.db.QueryRow(ctx, query, id).Scan(&genreID, genreTitle)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			r.logger.Error("genre not found", "id", id)
			return nil, err
		}
		r.logger.Error("failed to search genre", "error", err)
		return nil, err
	}

	return &domain.Genre{
		ID:    genreID,
		Title: genreTitle,
	}, nil
}

func (r *GenreRepository) UpdateGenre(
	ctx context.Context,
	id int,
	update dto.UpdateGenreRequest,
) error {
	var fields []string
	var args []interface{}
	argPos := 1

	if update.Title != nil {
		fields = append(fields, fmt.Sprintf("title=$%d", argPos))
		args = append(args, *update.Title)
		argPos++
	}

	if len(fields) == 0 {
		return nil
	}

	args = append(args, id)
	whereClause := fmt.Sprintf("where id=$%d", argPos)

	query := fmt.Sprintf("update genre set %s %s", strings.Join(fields, ", "), whereClause)

	res, err := r.db.Exec(
		ctx,
		query,
		args...,
	)

	if err != nil {
		r.logger.Error("failed to update genre", "id", id, "error", err)
		return err
	}

	rowsAffect := res.RowsAffected()
	if rowsAffect == 0 {
		r.logger.Info("no updated")
		return nil
	}

	return nil
}

func (r *GenreRepository) DeleteGenre(ctx context.Context, id int) error {
	query := `delete from genre where id=$1`

	if _, err := r.db.Exec(ctx, query, id); err != nil {
		return err
	}

	return nil
}
