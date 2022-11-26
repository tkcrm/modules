package logger_test

import (
	"testing"

	"github.com/tkcrm/modules/pkg/logger"
)

func TestLogger(t *testing.T) {
	l := logger.New(
		logger.WithAppName("test"),
		logger.WithLogLevel(logger.DEBUG),
	)

	l.Info("Hello world")
}
