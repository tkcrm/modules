package logger

import "github.com/aws/smithy-go/logging"

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

// Logger common interface
type Logger interface {
	Debug(...any)
	Debugf(string, ...any)
	Info(...any)
	Infof(string, ...any)
	Warn(...any)
	Warnf(string, ...any)
	Error(...any)
	Errorf(string, ...any)
	Fatal(...any)
	Fatalf(string, ...any)
	Panic(...any)
	Panicf(string, ...any)
	With(...any) Logger
	Sync() error
}

type AWSLogger interface {
	Logger
	Logf(classification logging.Classification, format string, v ...any)
}
