package sentry

import (
	"github.com/getsentry/sentry-go"
)

type Option func(h *hook)

func Hub(hub *sentry.Hub) Option {
	return func(h *hook) {
		if hub == nil {
			h.hub = sentry.CurrentHub()
		} else {
			h.hub = hub
		}
	}
}
