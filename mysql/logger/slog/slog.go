package slog

import (
	"log/slog"

	"github.com/blink-io/x/mysql/logger"
)

type log struct {
	sl *slog.Logger
}

func New(sl *slog.Logger) logger.Logger {
	return &log{sl: sl}
}

func (l *log) Print(v ...any) {
	l.sl.Info("[mysql]", slog.Any("values", v))
}
