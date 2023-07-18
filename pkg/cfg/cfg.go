package cfg

import (
	"fmt"
	"os"
	"path"
	"reflect"

	"github.com/cristalhq/aconfig"
	"github.com/cristalhq/aconfig/aconfigdotenv"
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
//
//	var config internalConfig.Config
//	if err := cfg.LoadConfig(&config); err != nil {
//		log.Fatalf("could not load configuration: %v", err)
//	}
func LoadConfig(cfg IConfig, opts ...Option) error {
	if reflect.ValueOf(cfg).Kind() != reflect.Ptr {
		return fmt.Errorf("config variable must be a pointer")
	}

	options := options{}
	for _, opt := range opts {
		opt(&options)
	}

	if options.EnvPath == "" {
		pwdDir, err := os.Getwd()
		if err != nil {
			return err
		}
		options.EnvPath = pwdDir
	}

	aconf := aconfig.Config{
		AllowUnknownFields: true,
		SkipFlags:          true,
		Files:              []string{path.Join(options.EnvPath, ".env")},
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
