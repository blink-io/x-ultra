//go:build other

package id

import (
	gofrsuuid "github.com/gofrs/uuid/v5"
)

func UUID() string {
	u, _ := gofrsuuid.NewV4()
	return u.String()
}
