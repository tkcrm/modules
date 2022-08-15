package bunconn

import (
	"database/sql"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/tkcrm/modules/logger"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
)

type Config struct {
	DSN      string
	BUNDEBUG bool
	// In seconds. Default 10 seconds
	PingInterval int `default:"10"`
}

func (c *Config) Validate() error {
	return validation.ValidateStruct(
		c,
		validation.Field(&c.DSN, validation.Required),
	)
}

type BunConn struct {
	DB *sql.DB
	bun.IDB
}

func New(cfg Config, logger logger.Logger) (*BunConn, error) {
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(cfg.DSN)))
	if err := sqldb.Ping(); err != nil {
		return nil, err
	}

	db := bun.NewDB(sqldb, pgdialect.New())
	if cfg.BUNDEBUG {
		db.AddQueryHook(bundebug.NewQueryHook(
			bundebug.WithVerbose(true),
		))
	}

	bunconn := &BunConn{sqldb, db}

	logger.Info("successfully connected to postgres")

	return bunconn, nil
}

func (s *BunConn) Ping() error {
	return s.DB.Ping()
}

func (s *BunConn) Close() error {
	return s.DB.Close()
}
