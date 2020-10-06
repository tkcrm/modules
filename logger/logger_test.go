package logger_test

import (
	"testing"

	"github.com/tkcrm/modules/logger"
)

func TestLogger(t *testing.T) {

	l := logger.DefaultLogger("debug", "tracker")

	l.Info("Hello world")
}
