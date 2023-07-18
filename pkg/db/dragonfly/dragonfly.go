package dragonfly

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/tkcrm/modules/pkg/logger"
)

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
		return nil, fmt.Errorf("failed to connect to dragonfly: %w", err)
	}

	logger.Info("successfully connected to dragonfly")

	return &Dragonfly{conn}, nil
}
