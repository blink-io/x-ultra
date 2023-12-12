package slog

import (
	"fmt"
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
	msg := fmt.Sprint(v...)
	l.sl.Info(msg)
}
