package redis

import "errors"

type Config struct {
	Enabled      bool   `yaml:"enabled" default:"true"`
	Addr         string `example:"localhost:6379"`
	User         string `secret:"true"`
	Password     string `secret:"true"`
	DbIndex      int    `validate:"gte=0" default:"0" yaml:"db_index"`
	PingInterval int    `validate:"gt=0" default:"10" yaml:"ping_interval"`
}

// Validate validates the config
func (c *Config) Validate() error {
	if !c.Enabled {
		return nil
	}

	if c.Addr == "" {
		return errors.New("addr is required")
	}

	if c.DbIndex < 0 {
		return errors.New("db_index must be greater than or equal to 0")
	}

	if c.PingInterval < 0 {
		return errors.New("ping_interval must be greater than 0")
	}

	return nil
}
