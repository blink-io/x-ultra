package mem

import (
	"context"
	"sync"
	"time"

	"github.com/blink-io/x/session/store"
)

type item struct {
	object     []byte
	expiration int64
}

// istore represents the session store.
type istore struct {
	items       map[string]item
	mu          sync.RWMutex
	stopCleanup chan bool
}

// New returns a new store instance, with a background cleanup goroutine that
// runs every minute to remove expired session data.
func New() store.Store {
	return NewWithCleanupInterval(time.Minute)
}

// NewWithCleanupInterval returns a new store instance. The cleanupInterval
// parameter controls how frequently expired session data is removed by the
// background cleanup goroutine. Setting it to 0 prevents the cleanup goroutine
// from running (i.e. expired sessions will not be removed).
func NewWithCleanupInterval(cleanupInterval time.Duration) store.Store {
	return newRawWithCleanupInterval(cleanupInterval)
}

func newRawWithCleanupInterval(cleanupInterval time.Duration) *istore {
	s := &istore{
		items: make(map[string]item),
	}
	if cleanupInterval > 0 {
		go s.startCleanup(cleanupInterval)
	}
	return s
}

// Find returns the data for a given session token from the store instance.
// If the session token is not found or is expired, the returned exists flag will
// be set to false.
func (s *istore) Find(ctx context.Context, token string) ([]byte, bool, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	item, found := s.items[token]
	if !found {
		return nil, false, nil
	}

	if time.Now().UnixNano() > item.expiration {
		return nil, false, nil
	}

	return item.object, true, nil
}

// Commit adds a session token and data to the store instance with the given
// expiry time. If the session token already exists, then the data and expiry
// time are updated.
func (s *istore) Commit(ctx context.Context, token string, b []byte, expiry time.Time) error {
	s.mu.Lock()
	s.items[token] = item{
		object:     b,
		expiration: expiry.UnixNano(),
	}
	s.mu.Unlock()

	return nil
}

// Delete removes a session token and corresponding data from the store
// instance.
func (s *istore) Delete(ctx context.Context, token string) error {
	s.mu.Lock()
	delete(s.items, token)
	s.mu.Unlock()

	return nil
}

// All returns a map containing the token and data for all active (i.e.
// not expired) sessions.
func (s *istore) All(ctx context.Context) (map[string][]byte, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var mm = make(map[string][]byte)

	for token, item := range s.items {
		if item.expiration > time.Now().UnixNano() {
			mm[token] = item.object
		}
	}

	return mm, nil
}

func (s *istore) startCleanup(interval time.Duration) {
	s.stopCleanup = make(chan bool)
	ticker := time.NewTicker(interval)
	for {
		select {
		case <-ticker.C:
			s.deleteExpired()
		case <-s.stopCleanup:
			ticker.Stop()
			return
		}
	}
}

// StopCleanup terminates the background cleanup goroutine for the store
// instance. It's rare to terminate this; generally store instances and
// their cleanup goroutines are intended to be long-lived and run for the lifetime
// of your application.
//
// There may be occasions though when your use of the store is transient.
// An example is creating a new store instance in a test function. In this
// scenario, the cleanup goroutine (which will run forever) will prevent the
// store object from being garbage collected even after the test function
// has finished. You can prevent this by manually calling StopCleanup.
func (s *istore) StopCleanup() {
	if s.stopCleanup != nil {
		s.stopCleanup <- true
	}
}

func (s *istore) deleteExpired() {
	now := time.Now().UnixNano()
	s.mu.Lock()
	for token, item := range s.items {
		if now > item.expiration {
			delete(s.items, token)
		}
	}
	s.mu.Unlock()
}
