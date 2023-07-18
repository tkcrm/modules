package natsconn

import (
	"fmt"

	"github.com/nats-io/nats.go"
	"github.com/tkcrm/modules/pkg/logger"
)

type Nats struct {
	connType    ConnType
	conn        *nats.Conn
	encodedConn *nats.EncodedConn
}

func New(logger logger.Logger, config Config, appName string, opts ...nats.Option) (*Nats, error) {
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

	return &Nats{
		connType: ConnTypeDefault,
		conn:     nc,
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

	nc.connType = ConnTypeEncoded
	nc.encodedConn, err = nats.NewEncodedConn(nc.conn, string(encType))
	if err != nil {
		return nil, err
	}

	return nc, nil
}
