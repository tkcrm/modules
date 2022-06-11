package cfg

type Option func(*Options)

type Options struct {
	EnvPath string
}

func WithEnvPath(v string) Option {
	return func(o *Options) {
		o.EnvPath = v
	}
}
