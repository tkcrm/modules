package postgres

import validation "github.com/go-ozzo/ozzo-validation/v4"

type Config struct {
	DSN          string `json:"POSTGRES_DSN"`
	PingInterval int    `json:"POSTGRES_PING_INTERVAL" default:"10"`
	MinConns     int32  `json:"POSTGRES_MIN_CONNS" default:"3"`
	MaxConns     int32  `json:"POSTGRES_MAX_CONNS" default:"6"`
}

func (c *Config) Validate() error {
	return validation.ValidateStruct(
		c,
		validation.Field(&c.DSN, validation.Required),
		validation.Field(&c.PingInterval, validation.Required),
		validation.Field(&c.MinConns, validation.Required),
		validation.Field(&c.MaxConns, validation.Required),
	)
}
