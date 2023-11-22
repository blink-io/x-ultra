package hash

import (
	"github.com/dchest/siphash"
)

func SIP(key []byte, data []byte) []byte {
	hh := siphash.New(key)
	hh.Write(data)
	sum := hh.Sum(nil)
	return sum
}

func SIP128(key []byte, data []byte) []byte {
	hh := siphash.New128(key)
	hh.Write(data)
	sum := hh.Sum(nil)
	return sum
}
