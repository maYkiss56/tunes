package repository

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	domain "github.com/maYkiss56/tunes/internal/domain/users"
	"github.com/maYkiss56/tunes/internal/logger"
)

type UserRepository struct {
	db     *pgxpool.Pool
	logger *logger.Logger
}

func NewUserRepository(db *pgxpool.Pool, logger *logger.Logger) *UserRepository {
	return &UserRepository{
		db:     db,
		logger: logger,
	}
}

func (r *UserRepository) CreateUser(ctx context.Context, user *domain.User) error {
	query := `insert into users
		(email, username, password_hash, role_id, created_at, updated_at)
		values ($1, $2, $3, $4, $5, $6) returning id`

	err := r.db.QueryRow(
		ctx,
		query,
		user.Email,
		user.Username,
		user.PasswordHash,
		user.RoleID,
		user.CreatedAt,
		user.UpdatedAt,
	).Scan(&user.ID)

	if err != nil {
		r.logger.Error("failed to create user", "error", err)
		return err
	}
	return nil
}

func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	query := `select id, email, username, password_hash, role_id from users where email=$1`

	var (
		userID       int
		userEmail    string
		userUsername string
		userPassword string
		userRoleID   int
	)

	err := r.db.QueryRow(ctx, query, email).
		Scan(&userID, &userEmail, &userUsername, &userPassword, &userRoleID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			r.logger.Error("user not found", "email", email, "error", err)
			return nil, err
		}
		r.logger.Error("failed to get user", "error", err)
		return nil, err
	}

	return &domain.User{
		ID:           userID,
		Email:        userEmail,
		Username:     userUsername,
		PasswordHash: userPassword,
		RoleID:       userRoleID,
	}, nil
}

func (r *UserRepository) GetUserByID(ctx context.Context, id int) (*domain.User, error) {
	query := `select id, email, username, password_hash, role_id from users where id=$1`

	var (
		userID       int
		userEmail    string
		userUsername string
		userPassword string
		userRoleID   int
	)

	err := r.db.QueryRow(ctx, query, id).
		Scan(&userID, &userEmail, &userUsername, &userPassword, &userRoleID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			r.logger.Error("user not found", "email", id, "error", err)
			return nil, err
		}
		r.logger.Error("failed to get user", "error", err)
		return nil, err
	}

	return &domain.User{
		ID:           userID,
		Email:        userEmail,
		Username:     userUsername,
		PasswordHash: userPassword,
		RoleID:       userRoleID,
	}, nil
}
