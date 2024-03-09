package sentry

import (
	"fmt"
	"time"

	"github.com/getsentry/sentry-go"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// InitSentryForZap init sentry for zap logger
func InitSentryForZap(cfg Config, opts ...Option) (zap.Option, error) {
	for _, opt := range opts {
		opt(&cfg)
	}

	if !cfg.Enabled {
		return nil, fmt.Errorf("sentry is not enabled")
	}

	if err := cfg.Validate(); err != nil {
		return nil, fmt.Errorf("validate sentry config error: %w", err)
	}

	sConf := cfg.sentryConfig
	sConf.Dsn = cfg.DSN
	sConf.Release = cfg.appVersion
	sConf.Environment = cfg.Environment
	sConf.TracesSampleRate = cfg.TracesSampleRate
	sConf.AttachStacktrace = cfg.AttachStacktrace

	if err := sentry.Init(sConf); err != nil {
		return nil, fmt.Errorf("init sentry error: %w", err)
	}

	return zap.Hooks(func(entry zapcore.Entry) error {
		if entry.Level >= zapcore.ErrorLevel {
			sentry.CaptureMessage(entry.Message)
		}

		return nil
	}), nil
}

// Flush execute sentry.Flush
//
// By default, the timeout is 2 seconds
func Flush(timeout time.Duration) bool {
	if timeout == 0 {
		timeout = 2 * time.Second
	}

	return sentry.Flush(timeout)
}
