package slog

import (
	"context"
	"log/slog"

	"github.com/jackc/pgx/v5/tracelog"
)

var _ tracelog.Logger = (*log)(nil)

type log struct {
	sl *slog.Logger
}

func New(sl *slog.Logger) tracelog.Logger {
	return &log{sl: sl}
}

func (l *log) Log(ctx context.Context, level tracelog.LogLevel, msg string, data map[string]any) {
	attrs := make([]slog.Attr, len(data))
	i := 0
	for k, v := range data {
		attrs[i] = slog.Any(k, v)
		i++
	}

	var slvl slog.Level

	switch level {
	case tracelog.LogLevelTrace:
		attrs = append(attrs, slog.String("PGX_LOG_LEVEL", level.String()))
		slvl = slog.LevelDebug
	case tracelog.LogLevelDebug:
		slvl = slog.LevelDebug
	case tracelog.LogLevelInfo:
		slvl = slog.LevelInfo
	case tracelog.LogLevelWarn:
		slvl = slog.LevelWarn
	case tracelog.LogLevelError:
		slvl = slog.LevelError
	default:
		attrs = append(attrs, slog.String("PGX_LOG_LEVEL", level.String()))
		slvl = slog.LevelError
	}
	l.sl.LogAttrs(ctx, slvl, msg, attrs...)
}
