package loglevel

import (
	"log/slog"

	"go.uber.org/zap/zapcore"
)

var slogToZapLevels = map[zapcore.Level]slog.Level{
	zapcore.DebugLevel: slog.LevelDebug,
	zapcore.InfoLevel:  slog.LevelInfo,
	zapcore.WarnLevel:  slog.LevelWarn,
	zapcore.ErrorLevel: slog.LevelError,
}

var zapToSlogLevels = map[slog.Level]zapcore.Level{
	slog.LevelDebug: zapcore.DebugLevel,
	slog.LevelInfo:  zapcore.InfoLevel,
	slog.LevelWarn:  zapcore.WarnLevel,
	slog.LevelError: zapcore.ErrorLevel,
}

func ZapToSlogLevel(level zapcore.Level) slog.Level {
	slevel, ok := slogToZapLevels[level]
	if ok {
		return slevel
	}
	return slog.LevelInfo
}

func SlogToZapLevel(level slog.Level) zapcore.Level {
	zlevel, ok := zapToSlogLevels[level]
	if ok {
		return zlevel
	}
	return zapcore.InfoLevel
}
