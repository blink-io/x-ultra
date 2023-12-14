package otel

import (
	"github.com/go-slog/otelslog"
	"github.com/remychantenay/slog-otel"
)

type Handler = slogotel.OtelHandler

type Handler2 = otelslog.Handler
