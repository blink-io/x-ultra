package nats

import (
	"log/slog"

	"github.com/samber/slog-nats"
)

type Option = slognats.Option

type Handler = slognats.NATSHandler

func NewHandler(o Option) slog.Handler {
	return o.NewNATSHandler()
}
