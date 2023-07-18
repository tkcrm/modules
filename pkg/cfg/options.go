package cfg

type Option func(*options)

type options struct {
	EnvPath string
}

func WithEnvPath(v string) Option {
	return func(o *options) {
		o.EnvPath = v
	}
}
