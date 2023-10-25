package slog

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/blink-io/x/redis/logging"
)

var _ logging.Logging = (*logz)(nil)

type logz struct {
	ll *slog.Logger
}

func New(ll *slog.Logger) logging.Logging {
	return &logz{ll}
}

func (s *logz) Printf(ctx context.Context, format string, v ...any) {
	msg := fmt.Sprintf(format, v...)
	s.ll.InfoContext(ctx, msg)
}
