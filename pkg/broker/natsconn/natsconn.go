package natsconn

import (
	"fmt"

	"github.com/nats-io/nats.go"
	"github.com/tkcrm/modules/pkg/logger"
)

type Nats struct {
	Conn *nats.Conn
	cfg  Config
}

func New(logger logger.Logger, config Config, appName string, opts ...nats.Option) (*Nats, error) {
	instance := &Nats{
		cfg: config,
	}

	if !config.Enabled {
		return instance, nil
	}

	if opts == nil {
		opts = make([]nats.Option, 0)
	}

	opts = append(opts, []nats.Option{
		nats.Name(appName),
		nats.DisconnectErrHandler(func(nc *nats.Conn, err error) {
			if err != nil {
				logger.Errorf("nats was disconnected with err: %v", err)
			}
		}),
		nats.ReconnectHandler(func(nc *nats.Conn) {
			logger.Warn("nats was reconnected")
		}),
		nats.ClosedHandler(func(nc *nats.Conn) {
			logger.Infof("nats connection closed")
		}),
		nats.MaxReconnects(-1),
	}...)

	if config.User != "" && config.Password != "" {
		opts = append(opts, nats.UserInfo(config.User, config.Password))
	}

	nc, err := nats.Connect(
		config.Addr,
		opts...,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to nats: %w", err)
	}

	instance.Conn = nc

	return instance, nil
}
