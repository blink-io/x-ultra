package session

import (
	"time"

	"github.com/blink-io/x/session/encoding"
	"github.com/blink-io/x/session/encoding/json"
	"github.com/blink-io/x/session/store"
	"github.com/blink-io/x/session/store/mem"
)

// Manager holds the configuration settings for your sessions.
type Manager struct {
	// IdleTimeout controls the maximum length of time a session can be inactive
	// before it expires. For example, some applications may wish to set this so
	// there is a timeout after 20 minutes of inactivity. By default, IdleTimeout
	// is not set and there is no inactivity timeout.
	IdleTimeout time.Duration

	// Lifetime controls the maximum length of time that a session is valid for
	// before it expires. The lifetime is an 'absolute expiry' which is set when
	// the session is first created and does not change. The default value is 24
	// hours.
	Lifetime time.Duration

	// Codec controls the encoder/decoder used to transform session data to a
	// byte slice for use by the session store. By default, session data is
	// encoded/decoded using encoding/gob.
	Codec encoding.Codec

	// Store controls the session store where the session data is persisted.
	Store store.Store

	// contextKey is the key used to set and retrieve the session data from a
	// context.Context. It's automatically generated to ensure uniqueness.
	contextKey contextKey
}

// NewManager returns a new session manager with the default options. It is safe for
// concurrent use.
func NewManager(ops ...Option) *Manager {
	m := &Manager{
		IdleTimeout: 0,
		Lifetime:    24 * time.Hour,
		Store:       mem.New(),
		Codec:       json.New(),
		contextKey:  generateContextKey(),
	}

	for _, o := range ops {
		o(m)
	}

	return m
}

//
//func (s *Manager) GetContextKey() contextKey {
//	return s.contextKey
//}
