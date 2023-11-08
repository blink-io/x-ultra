package session

import (
	"time"

	"github.com/blink-io/x/session/encoding"
	"github.com/blink-io/x/session/store"
)

type Option func(*manager)

func Store(s store.Store) Option {
	return func(m *manager) {
		m.store = s
	}
}

func Codec(c encoding.Codec) Option {
	return func(m *manager) {
		m.codec = c
	}
}

func IdleTimeout(t time.Duration) Option {
	return func(m *manager) {
		m.idleTimeout = t
	}
}

func Lifetime(t time.Duration) Option {
	return func(m *manager) {
		m.lifetime = t
	}
}

func UTC() Option {
	return func(m *manager) {
		m.isUTC = true
	}
}
