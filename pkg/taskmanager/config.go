package taskmanager

type Config struct {
	//ShutdownTime time.Duration `env:"SHUTDOWN_TIME" example:""`
}

func (c *Config) Validate() error {
	return nil
}
