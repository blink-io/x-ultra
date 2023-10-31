package session

import (
	"time"

	"github.com/blink-io/x/session/encoding"
	"github.com/blink-io/x/session/store"
)

type Option func(*Manager)

func Store(s store.Store) Option {
	return func(m *Manager) {
		m.Store = s
	}
}

func Codec(c encoding.Codec) Option {
	return func(m *Manager) {
		m.Codec = c
	}
}

func IdleTimeout(t time.Duration) Option {
	return func(m *Manager) {
		m.IdleTimeout = t
	}
}

func Lifetime(t time.Duration) Option {
	return func(m *Manager) {
		m.Lifetime = t
	}
}

func ContextKey(k contextKey) Option {
	return func(m *Manager) {
		m.contextKey = k
	}
}
