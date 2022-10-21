package mqttconn

import validation "github.com/go-ozzo/ozzo-validation/v4"

type Config struct {
	Host     string `json:"MQTT_HOST" default:"localhost"`
	Port     string `json:"MQTT_PORT" default:"1883"`
	User     string `json:"MQTT_USER"`
	Password string `json:"MQTT_PASSWORD"`
}

func (c *Config) Validate() error {
	return validation.ValidateStruct(
		c,
		validation.Field(&c.Host, validation.Required),
		validation.Field(&c.Port, validation.Required),
	)
}
