package hash

import "github.com/twmb/murmur3"

// MM3 is short for murmur3
func MM3(data []byte) []byte {
	mm := murmur3.New128()
	_, _ = mm.Write(data) //nolint:errcheck
	hdata := mm.Sum(nil)
	return hdata
}
