package dragonfly

import (
	"context"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
	"github.com/tkcrm/modules/pkg/logger"
)

type Config struct {
	Addr         string `json:"DRAGONFLY_ADDR"`
	User         string `json:"DRAGONFLY_USER" secret:"true"`
	Pass         string `json:"DRAGONFLY_PASS" secret:"true"`
	DbIndex      int    `json:"DRAGONFLY_DB_INDEX"`
	PingInterval int    `json:"TDENGINE_PING_INTERVAL" default:"10"`
}

func (c *Config) Validate() error {
	return validation.ValidateStruct(
		c,
		validation.Field(&c.Addr, validation.Required),
	)
}

type Dragonfly struct {
	Conn *redis.Client
}

func New(ctx context.Context, cfg Config, logger logger.Logger) (*Dragonfly, error) {
	conn := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Username: cfg.User,
		Password: cfg.Pass,
		DB:       cfg.DbIndex,
	})

	if err := conn.Ping(ctx).Err(); err != nil {
		return nil, errors.Wrap(err, "failed to connect to dragonfly")
	}

	logger.Info("successfully connected to dragonfly")

	return &Dragonfly{conn}, nil
}

func (p *Dragonfly) Ping(ctx context.Context) error {
	return p.Conn.Ping(ctx).Err()
}

func (p *Dragonfly) Close() error {
	return p.Conn.Close()
}
