package taskmanagerpg

import "time"

type Config struct {
	ShutdownTime time.Duration `yaml:"shutdown_time"`
}

func (c *Config) Validate() error {
	return nil
}
