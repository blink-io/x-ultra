package id

import (
	gofrsuuid "github.com/gofrs/uuid/v5"
)

func UUIDV4() string {
	u, _ := gofrsuuid.NewV4()
	return u.String()
}
