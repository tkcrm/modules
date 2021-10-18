package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger common interface
type Logger interface {
	Debug(...interface{})
	Debugf(string, ...interface{})
	Info(...interface{})
	Infof(string, ...interface{})
	Warn(...interface{})
	Warnf(string, ...interface{})
	Error(...interface{})
	Errorf(string, ...interface{})
	Fatal(...interface{})
	Fatalf(string, ...interface{})
	Panic(...interface{})
	Panicf(string, ...interface{})
	With(...interface{}) *zap.SugaredLogger
}

// New ...
func New(l string) *zap.Logger {

	atom := zap.NewAtomicLevel()

	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = "timestamp"
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	logger := zap.New(zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderCfg),
		zapcore.Lock(os.Stdout),
		atom,
	))
	defer logger.Sync()

	switch l {
	case "debug":
		atom.SetLevel(zap.DebugLevel)
	case "info":
		atom.SetLevel(zap.InfoLevel)
	case "warning":
		atom.SetLevel(zap.WarnLevel)
	case "error":
		atom.SetLevel(zap.ErrorLevel)
	case "fatal":
		atom.SetLevel(zap.FatalLevel)
	case "panic":
		atom.SetLevel(zap.PanicLevel)
	default:
		atom.SetLevel(zap.InfoLevel)
	}

	return logger
}

// DefaultLogger ...
func DefaultLogger(level, microservice string) Logger {
	if level == "" {
		level = "info"
	}
	return New(level).With(
		zap.String("ms", microservice),
	).Sugar()
}
