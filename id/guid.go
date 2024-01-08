package id

import (
	"github.com/beevik/guid"
)

func GUID() string {
	return guid.NewString()
}
