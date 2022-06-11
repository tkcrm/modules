package postgres

import (
	"context"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/tkcrm/modules/logger"
)

type Config struct {
	DSN string
	// In seconds. Default 10 seconds
	PingInterval int   `default:"10"`
	MaxConns     int32 `default:"6"`
	MinConns     int32 `default:"3"`
}

func (c *Config) Validate() error {
	return validation.ValidateStruct(
		c,
		validation.Field(&c.DSN, validation.Required),
	)
}

type PostgreSQL struct {
	*pgxpool.Pool
}

func New(ctx context.Context, cfg Config, logger logger.Logger) (*PostgreSQL, error) {

	config, err := pgxpool.ParseConfig(cfg.DSN)
	if err != nil {
		return nil, err
	}

	pool, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		return nil, err
	}

	psql := &PostgreSQL{pool}

	if err := psql.PingDB(); err != nil {
		return nil, err
	}

	logger.Info("successfully connected to postgres")

	return psql, nil
}

func (p *PostgreSQL) PingDB() error {
	return p.Ping(context.Background())
}
