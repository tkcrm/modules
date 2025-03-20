package redis

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type Redis struct {
	Conn *redis.Client
	cfg  Config
}

func New(ctx context.Context, cfg Config, logger logger) (*Redis, error) {
	instance := &Redis{
		cfg: cfg,
	}

	if !cfg.Enabled {
		return instance, nil
	}

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

	instance.Conn = conn

	return instance, nil
}
