package logger

import (
	"github.com/aws/smithy-go/logging"
	"go.uber.org/zap"
)

type LogLevel string

func (l LogLevel) String() string {
	return string(l)
}

const (
	LogLevelDebug   LogLevel = "debug"
	LogLevelInfo    LogLevel = "info"
	LogLevelWarning LogLevel = "warning"
	LogLevelError   LogLevel = "error"
	LogLevelFatal   LogLevel = "fatal"
	LogLevelPanic   LogLevel = "panic"
)

// GetAllLevels return all log levels. Used in validation.
func GetAllLevels() []any {
	return []any{
		LogLevelDebug.String(),
		LogLevelInfo.String(),
		LogLevelWarning.String(),
		LogLevelError.String(),
		LogLevelFatal.String(),
		LogLevelPanic.String(),
	}
}

type LogFormat string

const (
	LoggerFormatConsole LogFormat = "console"
	LoggerFormatJSON    LogFormat = "json"
)

type SugaredLogger = zap.SugaredLogger

// Logger common interface
type Logger interface {
	Debug(...any)
	Debugf(template string, args ...any)
	Debugw(msg string, keysAndValues ...any)

	Info(...any)
	Infof(template string, args ...any)
	Infow(msg string, keysAndValues ...any)

	Warn(...any)
	Warnf(template string, args ...any)
	Warnw(msg string, keysAndValues ...any)

	Error(...any)
	Errorf(template string, args ...any)
	Errorw(msg string, keysAndValues ...any)

	Fatal(...any)
	Fatalf(template string, args ...any)
	Fatalw(msg string, keysAndValues ...any)

	Panic(...any)
	Panicf(template string, args ...any)
	Panicw(msg string, keysAndValues ...any)

	With(...any) Logger

	Sugar() *SugaredLogger

	Sync() error
}

type AWSLogger interface {
	Logger
	Logf(classification logging.Classification, format string, v ...any)
}
