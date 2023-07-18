package bunconn

import "context"

func (s *BunConn) Name() string { return "bunconn" }

func (s *BunConn) Start(_ context.Context) error { return nil }

func (s *BunConn) Stop(_ context.Context) error {
	return s.DB.Close()
}

func (s *BunConn) Ping(_ context.Context) error {
	return s.DB.Ping()
}
