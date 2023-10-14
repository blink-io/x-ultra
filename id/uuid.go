//go:build !other

package id

import (
	"github.com/google/uuid"
)

func UUID() string {
	return uuid.NewString()
}
