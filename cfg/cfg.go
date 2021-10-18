package cfg

import (
	"fmt"

	"github.com/nats-io/nats.go"
	"github.com/tkcrm/modules/logger"
)

// GetNATSOpts return default NATS options for all microservices
func GetNATSOpts(l logger.Logger, appName, user, pass, token string) []nats.Option {
	opts := []nats.Option{
		nats.Name(appName),
		nats.DisconnectErrHandler(func(nc *nats.Conn, err error) {
			l.Error("Nats was disconnected")
		}),
		nats.ReconnectHandler(func(nc *nats.Conn) {
			l.Warn("Nats was reconnected")
		}),
		nats.ClosedHandler(func(nc *nats.Conn) {
			l.Errorf("Nats connection closed. Reason: %q", nc.LastError())
		}),
		nats.MaxReconnects(-1),
	}

	if user != "" && pass != "" {
		opts = append(opts, nats.UserInfo(user, pass))
	}

	if token != "" {
		opts = append(opts, nats.Token(token))
	}

	return opts
}

// GetNATSURL return formated NATS url
func GetNATSURL(host, port string) string {
	return fmt.Sprintf("%s:%s", host, port)
}

// GetDatabaseURL - return formated postgres url
func GetDatabaseURL(user, pass, host, port, name string) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		user,
		pass,
		host,
		port,
		name,
	)
}
