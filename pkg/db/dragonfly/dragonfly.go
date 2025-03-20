package dragonfly

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type Dragonfly struct {
	Conn *redis.Client
	cfg  Config
}

func New(ctx context.Context, cfg Config, logger logger) (*Dragonfly, error) {
	conn := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Username: cfg.User,
		Password: cfg.Password,
		DB:       cfg.DbIndex,
	})

	if err := conn.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("failed to connect to dragonfly: %w", err)
	}

	logger.Info("successfully connected to dragonfly")

	return &Dragonfly{conn, cfg}, nil
}
