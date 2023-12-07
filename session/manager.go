package session

import (
	"context"
	"time"

	"github.com/blink-io/x/session/encoding"
	"github.com/blink-io/x/session/encoding/msgpack"
	"github.com/blink-io/x/session/store"
	"github.com/blink-io/x/session/store/mem"
)

const (
	DefaultLifetime = 24 * time.Hour

	DefaultIdleTimeout = 0
)

type Manager interface {
	Load(ctx context.Context, token string) (context.Context, error)
	Commit(ctx context.Context) (string, time.Time, error)
	Destroy(ctx context.Context) error
	Put(ctx context.Context, key string, val any)
	Get(ctx context.Context, key string) any
	Pop(ctx context.Context, key string) any
	Remove(ctx context.Context, key string)
	Clear(ctx context.Context) error
	Exists(ctx context.Context, key string) bool
	Keys(ctx context.Context) []string
	RenewToken(ctx context.Context) error
	MergeSession(ctx context.Context, token string) error
	Status(ctx context.Context) Status
	GetString(ctx context.Context, key string) string
	GetBool(ctx context.Context, key string) bool
	GetInt(ctx context.Context, key string) int
	GetInt64(ctx context.Context, key string) int64
	GetInt32(ctx context.Context, key string) int32
	GetFloat(ctx context.Context, key string) float64
	GetBytes(ctx context.Context, key string) []byte
	GetTime(ctx context.Context, key string) time.Time
	PopString(ctx context.Context, key string) string
	PopBool(ctx context.Context, key string) bool
	PopInt(ctx context.Context, key string) int
	PopFloat(ctx context.Context, key string) float64
	PopBytes(ctx context.Context, key string) []byte
	PopTime(ctx context.Context, key string) time.Time
	SetRememberMe(ctx context.Context, key string, val bool)
	IsRememberMe(ctx context.Context, key string) bool
	Iterate(ctx context.Context, fn func(context.Context) error) error
	Deadline(ctx context.Context) time.Time
	Token(ctx context.Context) string
}

type TokenGenFunc func() (string, error)

// manager holds the configuration settings for your sessions.
type manager struct {
	// idleTimeout controls the maximum length of time a session can be inactive
	// before it expires. For example, some applications may wish to set this so
	// there is a timeout after 20 minutes of inactivity. By default, idleTimeout
	// is not set and there is no inactivity timeout.
	idleTimeout time.Duration

	// lifetime controls the maximum length of time that a session is valid for
	// before it expires. The lifetime is an 'absolute expiry' which is set when
	// the session is first created and does not change. The default value is 24
	// hours.
	lifetime time.Duration

	// codec controls the encoder/decoder used to transform session data to a
	// byte slice for use by the session store.
	// Session data is encoded/decoded using msgpack by default.
	codec encoding.Codec

	// store controls the session store where the session data is persisted.
	store store.Store

	// contextKey is the key used to set and retrieve the session data from a
	// context.Context. It's automatically generated to ensure uniqueness.
	contextKey contextKey

	tokenGenerator TokenGenFunc
}

// NewManager returns a new session manager with the default options. It is safe for
// concurrent use.
func NewManager(ops ...Option) Manager {
	return newManager(ops...)
}

func newManager(ops ...Option) *manager {
	m := &manager{
		idleTimeout:    DefaultIdleTimeout,
		lifetime:       DefaultLifetime,
		store:          mem.New(),
		codec:          msgpack.New(),
		contextKey:     generateContextKey(),
		tokenGenerator: generateToken,
	}

	for _, o := range ops {
		o(m)
	}

	m.setDefaults()

	return m
}

func (m *manager) setDefaults() {
	if m == nil {
		return
	}
	if m.tokenGenerator == nil {
		m.tokenGenerator = generateToken
	}
	if m.store == nil {
		m.store = mem.New()
	}
	if m.codec == nil {
		m.codec = msgpack.New()
	}
}
