package mqttconn

import validation "github.com/go-ozzo/ozzo-validation/v4"

type Config struct {
	Host string `default:"localhost"`
	Port string `default:"1883"`
	User string
	Pass string
}

func (c *Config) Validate() error {
	return validation.ValidateStruct(
		c,
		validation.Field(&c.Host, validation.Required),
		validation.Field(&c.Port, validation.Required),
	)
}
