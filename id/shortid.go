package id

import (
	"github.com/teris-io/shortid"
)

func ShortID() string {
	s, _ := shortid.Generate()
	return s
}
