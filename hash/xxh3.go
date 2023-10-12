package hash

import (
	"github.com/zeebo/xxh3"
)

func XXH3(data []byte) []byte {
	h := xxh3.New()
	_, _ = h.Write(data)
	sum := h.Sum(nil) // returns 8-byte []byte
	return sum
}
