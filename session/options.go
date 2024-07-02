package session

import (
	"time"

	"github.com/blink-io/x/session/encoding"
	"github.com/blink-io/x/session/store"
)

type Option func(*manager)

func applyOptions(m *manager, ops ...Option) {
	for _, o := range ops {
		o(m)
	}
}

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

func TokenGenerator(fn TokenGenFunc) Option {
	return func(m *manager) {
		m.tokenGen = fn
	}
}

func HashTokenInStore(b bool) Option {
	return func(m *manager) {
		m.hashTokenInStore = b
	}
}
