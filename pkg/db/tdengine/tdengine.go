package tdengine

import (
	"database/sql"
	"fmt"

	"github.com/tkcrm/modules/pkg/logger"

	_ "github.com/taosdata/driver-go/v3/taosSql"
)

type TDEngine struct {
	config Config
	DB     *sql.DB
}

func New(logger logger.Logger, cfg Config) (*TDEngine, error) {
	db, err := sql.Open("taosSql", cfg.DSN)
	if err != nil {
		return nil, fmt.Errorf("failed to connect tdengine: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	logger.Info("successfully connected to tdengine")

	return &TDEngine{cfg, db}, nil
}
