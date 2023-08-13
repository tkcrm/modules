package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/tkcrm/modules/pkg/db/dbutils"
	"github.com/tkcrm/modules/pkg/logger"
)

type PostgreSQL struct {
	DB *pgxpool.Pool
}

func New(ctx context.Context, cfg Config, logger logger.Logger) (*PostgreSQL, error) {
	config, err := pgxpool.ParseConfig(
		dbutils.PostgresDSN(cfg.Addr, cfg.User, cfg.Password, cfg.DBName),
	)
	if err != nil {
		return nil, err
	}

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, err
	}

	if err := pool.Ping(ctx); err != nil {
		return nil, err
	}

	logger.Info("successfully connected to postgres")

	return &PostgreSQL{pool}, nil
}
