package grpc

import (
	"github.com/blink-io/x/session"
)

type Option func(*SessionHandler)

func WithHeader(header string) Option {
	return func(sh *SessionHandler) {
		sh.header = header
	}
}

func WithExposeExpiry() Option {
	return func(sh *SessionHandler) {
		sh.exposeExpiry = true
	}
}

func WithSessionManager(sm *session.Manager) Option {
	return func(sh *SessionHandler) {
		sh.sm = sm
	}
}
