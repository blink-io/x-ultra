package sentry

import (
	"log/slog"

	"github.com/samber/slog-sentry/v2"
)

type Option = slogsentry.Option

type Handler = slogsentry.SentryHandler

func NewHandler(o Option) slog.Handler {
	return o.NewSentryHandler()
}
