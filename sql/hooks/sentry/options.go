package sentry

import "github.com/getsentry/sentry-go"

type Option func(*hook)

func Hub(hub *sentry.Hub) Option {
	return func(h *hook) {
		if h == nil {
			h.hub = sentry.CurrentHub()
		} else {
			h.hub = hub
		}
	}
}
