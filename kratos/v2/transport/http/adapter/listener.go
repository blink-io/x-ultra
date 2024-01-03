package adapter

import (
	"io"
	"net"
)

type Listener interface {
	Addr() net.Addr
	io.Closer
}
