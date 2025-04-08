package bunconn

import (
	"database/sql"

	"github.com/tkcrm/modules/pkg/db/dbutils"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
)

type BunConn struct {
	DB *sql.DB
	bun.IDB
}

func New(logger logger, cfg Config) (*BunConn, error) {
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(
		dbutils.PostgresDSN(cfg.Addr, cfg.User, cfg.Password, cfg.DbName),
	)))
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
