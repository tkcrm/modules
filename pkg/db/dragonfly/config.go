package dragonfly

import validation "github.com/go-ozzo/ozzo-validation/v4"

type Config struct {
	Addr         string `json:"DRAGONFLY_ADDR"`
	User         string `json:"DRAGONFLY_USER" secret:"true"`
	Pass         string `json:"DRAGONFLY_PASS" secret:"true"`
	DbIndex      int    `json:"DRAGONFLY_DB_INDEX"`
	PingInterval int    `json:"TDENGINE_PING_INTERVAL" default:"10"`
}

func (c *Config) Validate() error {
	return validation.ValidateStruct(
		c,
		validation.Field(&c.Addr, validation.Required),
		validation.Field(&c.PingInterval, validation.Required),
	)
}
