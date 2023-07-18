package natsconn

import "context"

func (s *Nats) Name() string { return "nats" }

func (s *Nats) Start(_ context.Context) error { return nil }

func (s *Nats) Stop(_ context.Context) error {
	if s.connType == ConnTypeDefault {
		s.conn.Close()
	}

	if s.connType == ConnTypeEncoded {
		s.encodedConn.Close()
	}

	return nil
}

func (s *Nats) Ping(_ context.Context) error {
	return nil
}
