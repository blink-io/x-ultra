package id

import (
	"github.com/segmentio/ksuid"
)

func KSUID() string {
	return ksuid.New().String()
}
