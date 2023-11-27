package httpsign

import (
	"sync"
	"time"

	"github.com/blink-io/x/cache"
	"github.com/blink-io/x/cache/ttlcache"
)

type nonceCache struct {
	sync.Mutex

	cache    cache.TTLCache[string]
	cacheTTL time.Duration
}

// Return a new nonceCache. Allows you to control cache capacity, ttl, as well as the TimeProvider.
func newNonceCache(cache cache.TTLCache[string], cacheTTL time.Duration) (*nonceCache, error) {
	if cache == nil {
		cache = ttlcache.New[string](cacheTTL)
	}
	return &nonceCache{
		cache:    cache,
		cacheTTL: cacheTTL,
	}, nil
}

// inCache checks if a nonce is in the cache. If not, it adds it to the
// cache and returns false. Otherwise it returns true.
func (n *nonceCache) inCache(nonce string) (exists bool, err error) {
	n.Lock()
	defer n.Unlock()

	// check if the nonce is already in the cache
	_, exists = n.cache.Get(nonce)
	if exists {
		return
	}

	n.cache.SetWithTTL(nonce, "", n.cacheTTL)
	return
}
