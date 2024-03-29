package natsconn

import (
	"fmt"

	"github.com/nats-io/nats.go"
	"github.com/tkcrm/modules/pkg/logger"
)

type Nats struct {
	ConnType    ConnType
	Conn        *nats.Conn
	EncodedConn *nats.EncodedConn
}

func New(logger logger.Logger, config Config, appName string, opts ...nats.Option) (*Nats, error) {
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

	return &Nats{
		ConnType: ConnTypeDefault,
		Conn:     nc,
	}, nil
}

func NewEncoded(logger logger.Logger, config Config, appName string, encType NatsEncodeType, opts ...nats.Option) (*Nats, error) {
	if encType == "" {
		return nil, fmt.Errorf("empty encode type")
	}

	nc, err := New(logger, config, appName, opts...)
	if err != nil {
		return nil, err
	}

	nc.ConnType = ConnTypeEncoded
	nc.EncodedConn, err = nats.NewEncodedConn(nc.Conn, string(encType))
	if err != nil {
		return nil, err
	}

	return nc, nil
}
