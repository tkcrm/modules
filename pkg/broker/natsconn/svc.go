package natsconn

import "context"

func (s *Nats) Name() string { return "nats" }

func (s *Nats) Start(_ context.Context) error { return nil }

func (s *Nats) Stop(_ context.Context) error {
	s.Conn.Close()
	return nil
}

func (s *Nats) Ping(_ context.Context) error {
	return nil
}

// Enabled returns true if the broker is enabled
func (s *Nats) Enabled() bool {
	return s.cfg.Enabled
}
