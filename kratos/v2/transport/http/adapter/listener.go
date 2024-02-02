package adapter

import (
	"io"
	"net"
)

type Listener interface {
	io.Closer

	Addr() net.Addr
}
