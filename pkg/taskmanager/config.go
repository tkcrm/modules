package taskmanager

import "time"

type Config struct {
	ShutdownTime time.Duration `env:"SHUTDOWN_TIME"`
}

func (c *Config) Validate() error {
	return nil
}
