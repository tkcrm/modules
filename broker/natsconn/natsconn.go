package natsconn

import (
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/nats-io/nats.go"
	"github.com/tkcrm/modules/logger"
)

type Config struct {
	DSN string
}

func (c *Config) Validate() error {
	return validation.ValidateStruct(
		c,
		validation.Field(&c.DSN, validation.Required),
	)
}

func New(logger logger.Logger, config Config, appName string, opts ...nats.Option) (*nats.Conn, error) {

	if opts == nil {
		opts = make([]nats.Option, 0)
	}

	opts = append(opts, []nats.Option{
		nats.Name(appName),
		nats.DisconnectErrHandler(func(nc *nats.Conn, err error) {
			logger.Error("nats was disconnected")
		}),
		nats.ReconnectHandler(func(nc *nats.Conn) {
			logger.Warn("nats was reconnected")
		}),
		nats.ClosedHandler(func(nc *nats.Conn) {
			logger.Errorf("nats connection closed. Reason: %q", nc.LastError())
		}),
		nats.MaxReconnects(-1),
	}...)

	nc, err := nats.Connect(
		config.DSN,
		opts...,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to nats: %v", err)
	}

	logger.Info("successfully connected to nats")

	return nc, nil
}
