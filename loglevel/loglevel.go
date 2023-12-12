package loglevel

import (
	"github.com/blink-io/x/log"
	"go.uber.org/zap/zapcore"
)

var logToZapLevels = map[zapcore.Level]log.Level{
	zapcore.DebugLevel: log.LevelDebug,
	zapcore.InfoLevel:  log.LevelInfo,
	zapcore.WarnLevel:  log.LevelWarn,
	zapcore.ErrorLevel: log.LevelError,
}

var zapToLogLevels = map[log.Level]zapcore.Level{
	log.LevelDebug: zapcore.DebugLevel,
	log.LevelInfo:  zapcore.InfoLevel,
	log.LevelWarn:  zapcore.WarnLevel,
	log.LevelError: zapcore.ErrorLevel,
}

func ZapToLogLevel(level zapcore.Level) log.Level {
	slevel, ok := logToZapLevels[level]
	if ok {
		return slevel
	}
	return log.LevelInfo
}

func LogToZapLevel(level log.Level) zapcore.Level {
	zlevel, ok := zapToLogLevels[level]
	if ok {
		return zlevel
	}
	return zapcore.InfoLevel
}
