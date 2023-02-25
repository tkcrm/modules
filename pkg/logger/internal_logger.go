package logger

import (
	"github.com/aws/smithy-go/logging"
	"go.uber.org/zap"
)

type internalLogger struct {
	*zap.SugaredLogger
}

func newInternalLogger(opts ...Option) *internalLogger {
	options := Options{}

	for _, opt := range opts {
		opt(&options)
	}

	if options.LogLevel == "" {
		options.LogLevel = LogLevelDebug
	}

	logger := initZapLogger(
		options.LogLevel,
		options.LogFormat,
		options.ConsoleColored,
		options.TimeKey,
	)

	if options.AppName != "" {
		logger = logger.With(
			zap.String("app", options.AppName),
		)
	}

	return &internalLogger{
		SugaredLogger: logger.Sugar(),
	}
}

func (l internalLogger) Logf(classification logging.Classification, format string, v ...interface{}) {
	switch string(classification) {
	case LogLevelDebug.String():
		l.Debugf(format, v)
	case LogLevelInfo.String():
		l.Infof(format, v)
	case LogLevelWarning.String():
		l.Warnf(format, v)
	case LogLevelError.String():
		l.Errorf(format, v)
	case LogLevelFatal.String():
		l.Fatalf(format, v)
	case LogLevelPanic.String():
		l.Panicf(format, v)
	default:
		l.Infof(format, v)
	}
}

func (l internalLogger) With(params ...any) Logger {
	return internalLogger{l.SugaredLogger.With(params...)}
}

func (l *internalLogger) AWSLogger() AWSLogger {
	return l
}
