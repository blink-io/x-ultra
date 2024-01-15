package webhook

import (
	"log/slog"

	"github.com/samber/slog-webhook"
)

type Option = slogwebhook.Option

type Handler = slogwebhook.WebhookHandler

func NewHandler(o Option) slog.Handler {
	return o.NewWebhookHandler()
}
