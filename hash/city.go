package hash

import (
	"github.com/go-faster/city"
)

func City128(data []byte) city.U128 {
	return city.Hash128(data)
}

func City64(data []byte) uint64 {
	return city.Hash64(data)
}

func City32(data []byte) uint32 {
	return city.Hash32(data)
}
