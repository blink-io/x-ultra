package id

import (
	"github.com/jaevor/go-nanoid"
)

func NanoID(len int) string {
	gen, _ := nanoid.Standard(len)
	return gen()
}
