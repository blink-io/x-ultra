package bytes

import (
	"github.com/go-faster/xor"
)

func Xor(dst, a, b []byte) int {
	return xor.Bytes(dst, a, b)
}
