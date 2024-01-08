package natsconn

import "context"

func (s *Nats) Name() string { return "nats" }

func (s *Nats) Start(_ context.Context) error { return nil }

func (s *Nats) Stop(_ context.Context) error {
	if s.ConnType == ConnTypeDefault {
		s.Conn.Close()
	}

	if s.ConnType == ConnTypeEncoded {
		s.EncodedConn.Close()
	}

	return nil
}

func (s *Nats) Ping(_ context.Context) error {
	return nil
}
