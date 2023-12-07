package http

import (
	"net/http"

	"github.com/blink-io/x/session"
	"github.com/blink-io/x/session/http/resolver"
)

type Option func(*SessionHandler)

func WithResolver(rv resolver.Resolver) Option {
	return func(sh *SessionHandler) {
		sh.resolver = rv
	}
}

func WithSessionManager(manager session.Manager) Option {
	return func(sh *SessionHandler) {
		if manager != nil {
			sh.manager = manager
		}
	}
}

func WithErrorFunc(errFunc func(http.ResponseWriter, *http.Request, error)) Option {
	return func(sh *SessionHandler) {
		sh.errFunc = errFunc
	}
}
