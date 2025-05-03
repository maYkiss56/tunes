package postgresql

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PgConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	Database string
	SSLMode  string
}

func NewPgConfig(username, password, host, port, database, sslmode string) *PgConfig {
	return &PgConfig{
		Username: username,
		Password: password,
		Host:     host,
		Port:     port,
		Database: database,
		SSLMode:  sslmode,
	}
}

type PgClient struct {
	pool *pgxpool.Pool
}

func NewClient(ctx context.Context, cfg *PgConfig) (*PgClient, error) {
	dsn := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.Username, cfg.Password,
		cfg.Host, cfg.Port, cfg.Database, cfg.SSLMode,
	)

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	pgxConfig, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to parse config: %v", err)
	}

	pool, err := pgxpool.NewWithConfig(ctx, pgxConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create pool: %v", err)
	}

	if err = pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("dont connect to db: %v", err)
	}

	log.Println("connect to DB successfully")
	return &PgClient{pool: pool}, nil
}

func (c *PgClient) GetPool() *pgxpool.Pool {
	return c.pool
}

func (c *PgClient) Close() {
	c.pool.Close()
}

