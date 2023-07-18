package tdengine

import validation "github.com/go-ozzo/ozzo-validation/v4"

type Config struct {
	DSN string `json:"TDENGINE_DSN"`
	// In seconds. Default 10 seconds
	PingInterval int  `json:"TDENGINE_PING_INTERVAL" default:"10"`
	Enabled      bool `json:"TDENGINE_ENABLED" default:"true"`
}

func (c *Config) Validate() error {
	if !c.Enabled {
		return nil
	}

	return validation.ValidateStruct(
		c,
		validation.Field(&c.DSN, validation.Required),
		validation.Field(&c.PingInterval, validation.Required),
	)
}
