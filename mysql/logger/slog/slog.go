package slog

import (
	"log/slog"

	"github.com/blink-io/x/mysql/logger"
)

type logz struct {
	sl *slog.Logger
}

func New(sl *slog.Logger) logger.Logger {
	return &logz{sl: sl}
}

func (l *logz) Print(v ...any) {
	l.sl.Info("[mysql]", slog.Any("args", v))
}
