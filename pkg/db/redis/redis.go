package redis

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/tkcrm/modules/pkg/logger"
)

type Redis struct {
	Conn *redis.Client
}

func New(ctx context.Context, cfg Config, logger logger.Logger) (*Redis, error) {
	conn := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Username: cfg.User,
		Password: cfg.Password,
		DB:       cfg.DbIndex,
	})

	if err := conn.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("failed to connect to redis: %w", err)
	}

	logger.Info("successfully connected to redis")

	return &Redis{conn}, nil
}
