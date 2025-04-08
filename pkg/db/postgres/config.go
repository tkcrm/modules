package postgres

import "errors"

type Config struct {
	Enabled      bool   `yaml:"enabled" default:"true"`
	Addr         string `example:"localhost:5432"`
	User         string `secret:"true"`
	Password     string `secret:"true"`
	DbName       string `yaml:"db_name"`
	PingInterval int    `yaml:"ping_interval" validate:"gt=0" default:"10"`
	MinConns     int32  `yaml:"min_conns" validate:"gt=0" default:"3"`
	MaxConns     int32  `yaml:"max_conns" validate:"gt=0" default:"6"`
}

// Validate validates the config
func (c *Config) Validate() error {
	if !c.Enabled {
		return nil
	}

	if c.Addr == "" {
		return errors.New("addr is required")
	}

	if c.User == "" {
		return errors.New("user is required")
	}

	if c.Password == "" {
		return errors.New("password is required")
	}

	if c.DbName == "" {
		return errors.New("db_name is required")
	}

	if c.PingInterval < 0 {
		return errors.New("ping_interval must be greater than 0")
	}

	if c.MinConns <= 0 {
		return errors.New("min_conns must be greater than 0")
	}

	if c.MaxConns <= 0 {
		return errors.New("max_conns must be greater than 0")
	}

	return nil
}
