package slog

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/blink-io/x/log"
)

var _ log.Logger = (*Logger)(nil)

type Logger struct {
	log *slog.Logger
}

func NewLogger(slog *slog.Logger) *Logger {
	return &Logger{slog}
}

func (l *Logger) Log(level log.Level, keyvals ...any) error {
	keylen := len(keyvals)
	if keylen == 0 || keylen%2 != 0 {
		l.log.Warn(fmt.Sprint("keyvals must appear in pairs: ", keyvals))
		return nil
	}

	args := make([]slog.Attr, 0, (keylen/2)+1)
	for i := 0; i < keylen; i += 2 {
		args = append(args, slog.Any(fmt.Sprint(keyvals[i]), keyvals[i+1]))
	}

	lv := slog.LevelInfo
	switch level {
	case log.LevelDebug:
		lv = slog.LevelDebug
	case log.LevelInfo:
		lv = slog.LevelInfo
	case log.LevelWarn:
		lv = slog.LevelWarn
	case log.LevelError:
		lv = slog.LevelError
	case log.LevelFatal:
		lv = slog.LevelError
	}

	l.log.LogAttrs(context.Background(), lv, "", args...)

	return nil
}
