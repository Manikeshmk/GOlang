package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger wraps zap logger
type Logger struct {
	*zap.SugaredLogger
}

// NewLogger creates a new logger instance
func NewLogger(environment string) *Logger {
	var zapConfig zap.Config

	if environment == "production" {
		zapConfig = zap.NewProductionConfig()
		zapConfig.OutputPaths = []string{"stdout"}
		zapConfig.ErrorOutputPaths = []string{"stderr"}
	} else {
		zapConfig = zap.NewDevelopmentConfig()
		zapConfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	logger, _ := zapConfig.Build(zap.AddCallerSkip(1))
	return &Logger{logger.Sugar()}
}

// NewTestLogger creates a logger for testing
func NewTestLogger() *Logger {
	zapConfig := zap.NewDevelopmentConfig()
	zapConfig.OutputPaths = []string{os.Stderr}
	logger, _ := zapConfig.Build()
	return &Logger{logger.Sugar()}
}

// Info logs an info level message
func (l *Logger) Info(msg string, fields ...interface{}) {
	l.SugaredLogger.Infow(msg, fields...)
}

// Error logs an error level message
func (l *Logger) Error(msg string, err error, fields ...interface{}) {
	if err != nil {
		l.SugaredLogger.Errorw(msg, append([]interface{}{"error", err}, fields...)...)
	} else {
		l.SugaredLogger.Errorw(msg, fields...)
	}
}

// Debug logs a debug level message
func (l *Logger) Debug(msg string, fields ...interface{}) {
	l.SugaredLogger.Debugw(msg, fields...)
}

// Warn logs a warn level message
func (l *Logger) Warn(msg string, fields ...interface{}) {
	l.SugaredLogger.Warnw(msg, fields...)
}
