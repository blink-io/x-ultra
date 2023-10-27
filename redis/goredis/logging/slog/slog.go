package slog

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/blink-io/x/redis/goredis/logging"
)

var _ logging.Logging = (*logger)(nil)

type logger struct {
	ll *slog.Logger
}

func New(ll *slog.Logger) logging.Logging {
	return &logger{ll}
}

func (s *logger) Printf(ctx context.Context, format string, v ...any) {
	s.ll.InfoContext(ctx, fmt.Sprintf(format, v...))
}
