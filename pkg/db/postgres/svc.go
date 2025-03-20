package postgres

import "context"

func (s *PostgreSQL) Name() string { return "postgres" }

func (s *PostgreSQL) Start(_ context.Context) error { return nil }

func (s *PostgreSQL) Stop(_ context.Context) error {
	s.DB.Close()
	return nil
}

func (s *PostgreSQL) Ping(ctx context.Context) error {
	return s.DB.Ping(ctx)
}

// Enabled returns true if the database is enabled
func (s *PostgreSQL) Enabled() bool {
	return s.cfg.Enabled
}
