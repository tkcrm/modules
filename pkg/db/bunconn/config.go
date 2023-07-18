package bunconn

import validation "github.com/go-ozzo/ozzo-validation/v4"

type Config struct {
	DSN      string `json:"POSTGRES_DSN"`
	BUNDEBUG bool   `json:"BUNDEBUG" env:"BUNDEBUG"`
	// In seconds. Default 10 seconds
	PingInterval int `json:"POSTGRES_PING_INTERVAL" default:"10"`
}

func (c *Config) Validate() error {
	return validation.ValidateStruct(
		c,
		validation.Field(&c.DSN, validation.Required),
	)
}
