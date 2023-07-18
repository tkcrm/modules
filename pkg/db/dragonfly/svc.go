package dragonfly

import "context"

func (s *Dragonfly) Name() string { return "dragonfly" }

func (s *Dragonfly) Start(_ context.Context) error { return nil }

func (s *Dragonfly) Stop(_ context.Context) error {
	return s.Conn.Close()
}

func (s *Dragonfly) Ping(ctx context.Context) error {
	return s.Conn.Ping(ctx).Err()
}
