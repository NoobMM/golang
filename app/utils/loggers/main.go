package loggers

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var config zap.Config

// BuildLogger ...
func BuildLogger(logLevel zapcore.Level) *zap.Logger {
	config = zap.NewProductionConfig()
	config.Level.SetLevel(logLevel)

	logger, err := config.Build()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()
	zap.ReplaceGlobals(logger)

	return logger
}

// GetLoggerLevel ...
func GetLoggerLevel() string {
	return config.Level.Level().CapitalString()
}

// SetLoggerLevel ...
func SetLoggerLevel(level zapcore.Level) {
	config.Level.SetLevel(level)
}
