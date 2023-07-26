package sentry

import (
	"fmt"
	"time"

	"github.com/getsentry/sentry-go"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// InitSentryForZap init sentry for zap logger
func InitSentryForZap(cfg Config, appVersion string) (zap.Option, error) {
	if err := sentry.Init(sentry.ClientOptions{
		Dsn:              cfg.DSN,
		Release:          appVersion,
		AttachStacktrace: true,
		Environment:      cfg.Environment,
	}); err != nil {
		return nil, fmt.Errorf("init sentry error: %w", err)
	}

	return zap.Hooks(func(entry zapcore.Entry) error {
		if entry.Level >= zapcore.ErrorLevel {
			defer sentry.Flush(2 * time.Second)
			sentry.CaptureMessage(
				fmt.Sprintf(
					"%s, Line No: %d :: %s\n\nstack:\n%s",
					entry.Caller.File,
					entry.Caller.Line,
					entry.Message,
					entry.Stack,
				),
			)
		}

		return nil
	}), nil
}
