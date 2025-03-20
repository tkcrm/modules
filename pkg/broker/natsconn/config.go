package natsconn

import "errors"

type Config struct {
	Enabled  bool   `yaml:"enabled" default:"true"`
	Addr     string `example:"localhost:4222"`
	User     string `secret:"true"`
	Password string `secret:"true"`
}

// Validate validates the config
func (c *Config) Validate() error {
	if !c.Enabled {
		return nil
	}

	if c.Addr == "" {
		return errors.New("addr is required")
	}

	return nil
}
