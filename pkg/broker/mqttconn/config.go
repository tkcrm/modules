package mqttconn

import "errors"

type Config struct {
	Enabled  bool   `yaml:"enabled" default:"true"`
	Addr     string `validate:"required,hostname_port" example:"localhost:1883"`
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
