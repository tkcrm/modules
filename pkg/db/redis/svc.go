package redis

import "context"

func (s *Redis) Name() string { return "redis" }

func (s *Redis) Start(_ context.Context) error { return nil }

func (s *Redis) Stop(_ context.Context) error {
	return s.Conn.Close()
}

func (s *Redis) Ping(ctx context.Context) error {
	return s.Conn.Ping(ctx).Err()
}

// Enabled returns true if the database is enabled
func (s *Redis) Enabled() bool {
	return s.cfg.Enabled
}
