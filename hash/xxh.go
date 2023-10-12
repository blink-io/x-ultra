package hash

import (
	"github.com/cespare/xxhash/v2"
)

func XXH(data, key []byte) []byte {
	h := xxhash.New()
	_, _ = h.Write(data)
	return h.Sum(nil)
}
