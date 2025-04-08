package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/tkcrm/modules/pkg/db/dbutils"
)

type PostgreSQL struct {
	DB  *pgxpool.Pool
	cfg Config
}

func New(ctx context.Context, cfg Config, logger logger) (*PostgreSQL, error) {
	instance := &PostgreSQL{
		cfg: cfg,
	}

	if !cfg.Enabled {
		return instance, nil
	}

	config, err := pgxpool.ParseConfig(
		dbutils.PostgresDSN(cfg.Addr, cfg.User, cfg.Password, cfg.DbName),
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

	instance.DB = pool

	return instance, nil
}
