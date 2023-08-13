package sentry

import validation "github.com/go-ozzo/ozzo-validation/v4"

type Config struct {
	DSN         string
	Environment string
	Enabled     bool `default:"true"`
}

func (c *Config) Validate() error {
	if !c.Enabled {
		return nil
	}

	return validation.ValidateStruct(
		c,
		validation.Field(&c.DSN, validation.Required),
		validation.Field(&c.Environment, validation.Required),
	)
}
