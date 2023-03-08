package logger_test

import (
	"testing"

	"github.com/aws/smithy-go/logging"
	"github.com/tkcrm/modules/pkg/logger"
)

func TestLogger(t *testing.T) {
	l := logger.New(
		logger.WithAppName("test"),
		logger.WithLogLevel(logger.LogLevelDebug),
	)

	l.Info("Hello world")
}

func Test_LoggerWith(t *testing.T) {
	l := logger.New(
		logger.WithAppName("test"),
		logger.WithLogLevel(logger.LogLevelDebug),
		logger.WithLogFormat(logger.LoggerFormatConsole),
	).With("key", "value").With("key2", "value2")

	l = l.With("key3", "value3")

	l.Infof("some test value: %d", 1234)
}

func Test_AWSLogger(t *testing.T) {
	l := logger.NewAWSLogger(
		logger.WithAppName("test"),
		logger.WithLogLevel(logger.LogLevelDebug),
	)

	l.Logf(logging.Debug, "test: %s", "1234")
}
