package hash

import "github.com/twmb/murmur3"

func Murmur3(data []byte) []byte {
	mm := murmur3.New128()
	_, _ = mm.Write(data) //nolint:errcheck
	hdata := mm.Sum(nil)
	return hdata
}
