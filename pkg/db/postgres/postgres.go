package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/tkcrm/modules/pkg/db/dbutils"
	"github.com/tkcrm/modules/pkg/retry"
)

type PostgreSQL struct {
	DB  *pgxpool.Pool
	cfg Config
}

func New(ctx context.Context, cfg Config) (*PostgreSQL, error) {
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

	instance.DB = pool

	return instance, nil
}

// CheckConnection checks if the PostgreSQL connection pool is initialized and attempts to ping the database.
// It retries the ping operation up to 5 times with a 2-second delay between attempts.
// If the connection pool is not initialized, it returns an error.
func (p *PostgreSQL) CheckConnection(ctx context.Context, logger logger) error {
	if p.DB == nil {
		return fmt.Errorf("PostgreSQL connection pool is not initialized")
	}

	return retry.New(
		retry.WithDelay(time.Second*3),
		retry.WithMaxAttempts(10),
		retry.WithPolicy(retry.PolicyLinear),
		retry.WithLogger(logger),
		retry.WithContext(ctx),
	).Do(func() error {
		return p.DB.Ping(ctx)
	})
}
