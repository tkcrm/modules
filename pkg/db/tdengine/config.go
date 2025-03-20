package tdengine

import "errors"

type Config struct {
	Enabled  bool   `yaml:"enabled" default:"true"`
	Addr     string `example:"localhost:6030"`
	User     string `secret:"true"`
	Password string `secret:"true"`
	DBName   string `yaml:"db_name" env:"DB_NAME"`
	// In seconds. Default 10 seconds
	PingInterval int `yaml:"ping_interval" validate:"gt=0" default:"10"`
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

	if c.DBName == "" {
		return errors.New("db_name is required")
	}

	if c.PingInterval < 0 {
		return errors.New("ping_interval must be greater than 0")
	}

	return nil
}
