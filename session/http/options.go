package http

import (
	"net/http"
	"time"

	"github.com/blink-io/x/session/encoding"
	"github.com/blink-io/x/session/http/resolver"
	"github.com/blink-io/x/session/store"
)

type Option func(*Manager)

func WithResolver(rv resolver.Resolver) Option {
	return func(m *Manager) {
		m.rv = rv
	}
}

func WithStore(s store.Store) Option {
	return func(m *Manager) {
		m.Store = s
	}
}

func WithCodec(c encoding.Codec) Option {
	return func(m *Manager) {
		m.Codec = c
	}
}

func WithErrorFunc(ef func(http.ResponseWriter, *http.Request, error)) Option {
	return func(m *Manager) {
		m.errorFunc = ef
	}
}

func WithLifetime(t time.Duration) Option {
	return func(m *Manager) {
		m.Lifetime = t
	}
}

func WithIdleTimeout(t time.Duration) Option {
	return func(m *Manager) {
		m.IdleTimeout = t
	}
}
