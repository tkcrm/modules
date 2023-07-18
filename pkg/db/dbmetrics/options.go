package dbmetrics

type Option func(*metric)

func WithName(v string) Option {
	return func(m *metric) {
		m.funcName = v
	}
}

func WithPrefix(v string) Option {
	return func(m *metric) {
		m.metricPrefix = v
	}
}

func WithSuffix(v string) Option {
	return func(m *metric) {
		m.metricSuffix = v
	}
}
