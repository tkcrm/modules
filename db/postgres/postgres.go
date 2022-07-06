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

	psql := &PostgreSQL{pool}

	if err := psql.Ping(); err != nil {
		return nil, err
	}

	logger.Info("successfully connected to postgres")

	return psql, nil
}

func (p *PostgreSQL) Ping() error {
	return p.DB.Ping(context.Background())
}

func (p *PostgreSQL) Close() {
	p.DB.Close()
}
