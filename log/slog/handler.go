package slog

import (
	"context"
	"log/slog"

	"github.com/blink-io/x/log"
)

var logToSlogLevels = map[log.Level]slog.Level{
	log.LevelDebug: slog.LevelDebug,
	log.LevelInfo:  slog.LevelInfo,
	log.LevelWarn:  slog.LevelWarn,
	log.LevelError: slog.LevelError,
}

type handler struct {
	logger log.Logger
}

func NewHandler(logger log.Logger) slog.Handler {
	h := &handler{
		logger: logger,
	}
	return h
}

func (h *handler) Enabled(ctx context.Context, level slog.Level) bool {
	return false
}

func (h *handler) Handle(ctx context.Context, record slog.Record) error {
	return nil
}

func (h *handler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return h
}

func (h *handler) WithGroup(name string) slog.Handler {
	return h
}
