package session

import (
	"context"
	"time"

	"github.com/blink-io/x/session/encoding"
	"github.com/blink-io/x/session/encoding/msgpack"
	"github.com/blink-io/x/session/store"
	"github.com/blink-io/x/session/store/mem"
)

type Manager interface {
	Load(ctx context.Context, token string) (context.Context, error)
	Commit(ctx context.Context) (string, time.Time, error)
	Destroy(ctx context.Context) error
	Put(ctx context.Context, key string, val interface{})
	Get(ctx context.Context, key string) interface{}
	Pop(ctx context.Context, key string) interface{}
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

// manager holds the configuration settings for your sessions.
type manager struct {
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
func NewManager(ops ...Option) Manager {
	return newManager(ops...)
}

func newManager(ops ...Option) *manager {
	m := &manager{
		IdleTimeout: 0,
		Lifetime:    24 * time.Hour,
		Store:       mem.New(),
		Codec:       msgpack.New(),
		contextKey:  generateContextKey(),
	}

	for _, o := range ops {
		o(m)
	}

	return m
}
