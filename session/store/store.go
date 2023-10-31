package store

import (
	"context"
	"time"
)

const DefaultPrefix = "default:session:"

// Store is an interface for session stores which take a context.Context
// parameter.
type Store interface {

	// Delete is the same as Store.Delete, except it takes a context.Context.
	Delete(ctx context.Context, token string) (err error)

	// Find is the same as Store.Find, except it takes a context.Context.
	Find(ctx context.Context, token string) (b []byte, found bool, err error)

	// Commit is the same as Store.Commit, except it takes a context.Context.
	Commit(ctx context.Context, token string, b []byte, expiry time.Time) (err error)
	// All is the same as Store.All, expect it takes a
	// context.Context.
	All(ctx context.Context) (map[string][]byte, error)
}
