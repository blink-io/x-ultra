package id

import (
	"github.com/oklog/ulid/v2"
)

func ULID() string {
	return ulid.Make().String()
}
