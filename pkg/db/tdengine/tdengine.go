package tdengine

import (
	"database/sql"
	"fmt"

	_ "github.com/taosdata/driver-go/v3/taosSql"
)

type TDEngine struct {
	DB  *sql.DB
	cfg Config
}

func New(logger logger, cfg Config) (*TDEngine, error) {
	instance := &TDEngine{
		cfg: cfg,
	}

	if !cfg.Enabled {
		return instance, nil
	}

	db, err := sql.Open("taosSql",
		fmt.Sprintf(
			"%s:%s@tcp(%s)/%s",
			cfg.User, cfg.Password, cfg.Addr, cfg.DbName,
		),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to connect tdengine: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	logger.Info("successfully connected to tdengine")

	instance.DB = db

	return instance, nil
}
