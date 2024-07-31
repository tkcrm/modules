package tdengine

import (
	"database/sql"
	"fmt"

	_ "github.com/taosdata/driver-go/v3/taosSql"
)

type TDEngine struct {
	config Config
	DB     *sql.DB
}

func New(logger logger, cfg Config) (*TDEngine, error) {
	db, err := sql.Open("taosSql",
		fmt.Sprintf(
			"%s:%s@tcp(%s)/%s",
			cfg.User, cfg.Password, cfg.Addr, cfg.DBName,
		),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to connect tdengine: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	logger.Info("successfully connected to tdengine")

	return &TDEngine{cfg, db}, nil
}
