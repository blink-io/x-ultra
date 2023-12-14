package slog

import (
	"context"
	"log/slog"
)

const (
	LevelFatal slog.Level = 12
)

type Logger struct {
	*slog.Logger
}

func NewExtLogger(l *slog.Logger) *Logger {
	return &Logger{l}
}

func (l *Logger) Fatal(msg string, args ...any) {
	l.FatalContext(context.Background(), msg, args...)
}

func (l *Logger) FatalContext(ctx context.Context, msg string, args ...any) {
	l.Log(ctx, LevelFatal, msg, args...)
}
