package tdengine

import (
	"database/sql"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/pkg/errors"
	"github.com/tkcrm/modules/pkg/logger"

	_ "github.com/taosdata/driver-go/v3/taosSql"
)

type Config struct {
	DSN string `json:"TDENGINE_DSN"`
	// In seconds. Default 10 seconds
	PingInterval int `json:"TDENGINE_PING_INTERVAL" default:"10"`
}

func (c *Config) Validate() error {
	return validation.ValidateStruct(
		c,
		validation.Field(&c.DSN, validation.Required),
	)
}

type TDEngine struct {
	DB *sql.DB
}

func New(cfg Config, logger logger.Logger) (*TDEngine, error) {
	db, err := sql.Open("taosSql", cfg.DSN)
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect tdengine")
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	logger.Info("successfully connected to tdengine")

	return &TDEngine{db}, nil
}

func (p *TDEngine) Ping() error {
	return p.DB.Ping()
}

func (p *TDEngine) Close() error {
	return p.DB.Close()
}
