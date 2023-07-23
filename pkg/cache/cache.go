package cache

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/goccy/go-json"
	"github.com/redis/go-redis/v9"
)

type Cache interface {
	Get(ctx context.Context, key string) ([]byte, error)
	GetFromJSON(ctx context.Context, key string, dst any) error
	Set(ctx context.Context, key string, value any, expiration time.Duration) error
	SetJSON(ctx context.Context, key string, value any, expiration time.Duration) error
	Delete(ctx context.Context, key string) error
	Exists(ctx context.Context, key string) (bool, error)
}

type cache struct {
	conn *redis.Client
}

func New(redisClient *redis.Client) Cache {
	return &cache{redisClient}
}

func (c *cache) Get(ctx context.Context, key string) ([]byte, error) {
	cmd := c.conn.Get(ctx, key)
	if cmd.Err() != nil {
		if errors.Is(cmd.Err(), redis.Nil) {
			return nil, ErrNotFound
		}
		return nil, cmd.Err()
	}
	return cmd.Bytes()
}

func (c *cache) GetFromJSON(ctx context.Context, key string, dst any) error {
	cmd := c.conn.Get(ctx, key)
	if cmd.Err() != nil {
		if errors.Is(cmd.Err(), redis.Nil) {
			return ErrNotFound
		}
		return cmd.Err()
	}

	res, err := cmd.Bytes()
	if err != nil {
		return fmt.Errorf("failed to get result from cmd: %w", err)
	}

	if err := json.Unmarshal(res, dst); err != nil {
		return fmt.Errorf("failed to unmarshal to dst: %w", err)
	}

	return nil
}

func (c *cache) Set(ctx context.Context, key string, value any, expiration time.Duration) error {
	return c.conn.Set(ctx, key, value, expiration).Err()
}

func (c *cache) SetJSON(ctx context.Context, key string, value any, expiration time.Duration) error {
	marshalledData, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return c.conn.Set(ctx, key, marshalledData, expiration).Err()
}

func (c *cache) Delete(ctx context.Context, key string) error {
	return c.conn.Del(ctx, key).Err()
}

func (c *cache) Exists(ctx context.Context, key string) (bool, error) {
	_, err := c.Get(ctx, key)
	if err == nil {
		return true, nil
	}

	if errors.Is(err, ErrNotFound) {
		return false, nil
	}

	return false, err
}
