package id

import "github.com/rs/xid"

func XID() string {
	return xid.New().String()
}
