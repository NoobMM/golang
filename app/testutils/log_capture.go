package testutils

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest/observer"
)

// SetupLogsCapture is util for a test that want to assert a zap log
func SetupLogsCapture() (*zap.Logger, *observer.ObservedLogs) {
	core, logs := observer.New(zap.InfoLevel)
	logger := zap.New(core)
	zap.ReplaceGlobals(logger)
	return logger, logs
}
