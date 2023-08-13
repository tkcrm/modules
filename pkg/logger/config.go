package logger

import validation "github.com/go-ozzo/ozzo-validation/v4"

type Config struct {
	EncodingConsole bool   `usage:"allows to set user-friendly formatting"`
	Level           string `default:"info" usage:"allows to set custom logger level" example:"debug, info, warn, error"`
	Trace           string `default:"fatal" usage:"allows to set custom trace level"`
	WithCaller      bool   `usage:"allows to show caller"`
	WithStackTrace  bool   `usage:"allows to show stack trace"`
}

func (c *Config) Validate() error {
	if err := validation.ValidateStruct(c,
		validation.Field(&c.Level, validation.Required, validation.In(allLevels...)),
		validation.Field(&c.Trace, validation.Required, validation.In(allLevels...)),
	); err != nil {
		return err
	}

	return nil
}
