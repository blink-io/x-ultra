package httpbase

import (
	"io"
	"net"
)

type Listener interface {
	Addr() net.Addr
	io.Closer
}
