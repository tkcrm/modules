package cfg

import (
	"fmt"
	"os"
	"path"
	"reflect"

	"github.com/cristalhq/aconfig"
	"github.com/cristalhq/aconfig/aconfigdotenv"
	"github.com/nats-io/nats.go"
	"github.com/tkcrm/modules/logger"
)

type IConfig interface {
	Validate() error
}

// LoadConfig - load environment variables from `os env`, `.env` file and pass it to struct.
//
// For local development use `.env` file from root project.
//
// LoadConfig also call a `Validate` method.
//
// Example:
//	var config internalConfig.Config
//	if err := cfg.LoadConfig(&config); err != nil {
//		log.Fatalf("could not load configuration: %v", err)
//	}
func LoadConfig(cfg IConfig) error {

	if reflect.ValueOf(cfg).Kind() != reflect.Ptr {
		return fmt.Errorf("config variable must be a pointer")
	}

	pwdDir, err := os.Getwd()
	if err != nil {
		return err
	}

	aconf := aconfig.Config{
		AllowUnknownFields: true,
		SkipFlags:          true,
		Files:              []string{path.Join(pwdDir, ".env")},
		FileDecoders: map[string]aconfig.FileDecoder{
			".env": aconfigdotenv.New(),
		},
	}

	loader := aconfig.LoaderFor(cfg, aconf)
	if err := loader.Load(); err != nil {
		return err
	}

	return cfg.Validate()
}

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

// GetPostgreSqlURL - return formated postgres url
func GetPostgreSqlURL(user, pass, host, port, name string) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		user,
		pass,
		host,
		port,
		name,
	)
}
