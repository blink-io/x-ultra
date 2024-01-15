package zap

import (
	"log/slog"

	slogzap "github.com/samber/slog-zap/v2"
)

type Option = slogzap.Option

type Handler = slogzap.ZapHandler

func NewHandler(o Option) slog.Handler {
	return o.NewZapHandler()
}
