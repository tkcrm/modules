package taskmanagerpg

import "time"

type Config struct {
	ShutdownTime time.Duration `env:"SHUTDOWN_TIME" yaml:"shutdown_time"`
}

func (c *Config) Validate() error {
	return nil
}
