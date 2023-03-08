package logger

import (
	"github.com/aws/smithy-go/logging"
	"go.uber.org/zap"
)

type logger struct {
	*SugaredLogger
}

func newInternalLogger(opts ...Option) *logger {
	options := Options{}

	for _, opt := range opts {
		opt(&options)
	}

	if options.LogLevel == "" {
		options.LogLevel = LogLevelDebug
	}

	l := initZapLogger(
		options.LogLevel,
		options.LogFormat,
		options.ConsoleColored,
		options.TimeKey,
	)

	if options.AppName != "" {
		l = l.With(
			zap.String("app", options.AppName),
		)
	}

	return &logger{
		SugaredLogger: l.Sugar(),
	}
}

func (l logger) Logf(classification logging.Classification, format string, v ...interface{}) {
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

func (l logger) With(args ...any) Logger {
	return &logger{l.SugaredLogger.With(args...)}
}

func (l *logger) AWSLogger() AWSLogger {
	return l
}

func (l *logger) Sugar() *SugaredLogger {
	return l.SugaredLogger
}
