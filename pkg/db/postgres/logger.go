package postgres

type logger interface {
	Infof(format string, args ...any)
}
