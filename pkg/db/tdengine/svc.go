package tdengine

import "context"

func (s *TDEngine) Name() string { return "tdengine" }

func (s *TDEngine) Start(_ context.Context) error { return nil }

func (s *TDEngine) Stop(_ context.Context) error {
	return s.DB.Close()
}

func (s *TDEngine) Ping(_ context.Context) error {
	return s.DB.Ping()
}
