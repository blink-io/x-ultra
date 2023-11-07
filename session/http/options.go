package http

import (
	"net/http"

	"github.com/blink-io/x/session"
	"github.com/blink-io/x/session/http/resolver"
)

type Option func(*SessionHandler)

func WithResolver(rv resolver.Resolver) Option {
	return func(sh *SessionHandler) {
		sh.rv = rv
	}
}

func WithSessionManager(sm *session.Manager) Option {
	return func(sh *SessionHandler) {
		sh.sm = sm
	}
}

func WithErrorFunc(ef func(http.ResponseWriter, *http.Request, error)) Option {
	return func(sh *SessionHandler) {
		sh.ef = ef
	}
}
