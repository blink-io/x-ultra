package slog

import (
	"context"
	"fmt"
	"log/slog"
	"strings"

	"github.com/blink-io/x/gocron"
)

type logger struct {
	log *slog.Logger
}

func NewLogger(log *slog.Logger) gocron.Logger {
	return &logger{log: log}
}

func (l *logger) Debug(msg string, args ...any) {
	l.doLog(slog.LevelDebug, msg, args...)
}

func (l *logger) Error(msg string, args ...any) {
	l.doLog(slog.LevelError, msg, args...)
}

func (l *logger) Info(msg string, args ...any) {
	l.doLog(slog.LevelInfo, msg, args...)
}

func (l *logger) Warn(msg string, args ...any) {
	l.doLog(slog.LevelWarn, msg, args...)
}

func (l *logger) doLog(level slog.Level, msg string, args ...any) {
	l.log.LogAttrs(context.Background(), level, msg, slog.String("args", logFormatArgs(args...)))
}

func logFormatArgs(args ...any) string {
	if len(args) == 0 {
		return ""
	}
	if len(args)%2 != 0 {
		return ", " + fmt.Sprint(args...)
	}
	var pairs []string
	for i := 0; i < len(args); i += 2 {
		pairs = append(pairs, fmt.Sprintf("%s=%v", args[i], args[i+1]))
	}
	return ", " + strings.Join(pairs, ", ")
}
