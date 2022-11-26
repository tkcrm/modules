package postgres

import (
	"context"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/tkcrm/modules/pkg/logger"
)

type Config struct {
	DSN          string `json:"POSTGRES_DSN"`
	PingInterval int    `json:"POSTGRES_PING_INTERVAL" default:"10"`
	MinConns     int32  `json:"POSTGRES_MIN_CONNS" default:"3"`
	MaxConns     int32  `json:"POSTGRES_MAX_CONNS" default:"6"`
}

func (c *Config) Validate() error {
	return validation.ValidateStruct(
		c,
		validation.Field(&c.DSN, validation.Required),
	)
}

type PostgreSQL struct {
	DB *pgxpool.Pool
}

func New(ctx context.Context, cfg Config, logger logger.Logger) (*PostgreSQL, error) {
	config, err := pgxpool.ParseConfig(cfg.DSN)
	if err != nil {
		return nil, err
	}

	pool, err := pgxpool.ConnectConfig(ctx, config)
	if err != nil {
		return nil, err
	}

	if err := pool.Ping(ctx); err != nil {
		return nil, err
	}

	logger.Info("successfully connected to postgres")

	return &PostgreSQL{pool}, nil
}

func (p *PostgreSQL) Ping(ctx context.Context) error {
	return p.DB.Ping(ctx)
}

func (p *PostgreSQL) Close() {
	p.DB.Close()
}
