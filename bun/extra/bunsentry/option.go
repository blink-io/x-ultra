package bunsentry

import (
	"github.com/getsentry/sentry-go"
)

type Option func(h *QueryHook)

func Hub(hub *sentry.Hub) Option {
	return func(h *QueryHook) {
		if hub == nil {
			h.Hub = sentry.CurrentHub()
		} else {
			h.Hub = hub
		}
	}
}
