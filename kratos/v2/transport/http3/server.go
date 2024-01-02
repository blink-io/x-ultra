package http3

import (
	"github.com/blink-io/x/kratos/v2/transport/httpbase"
)

type (
	Server = httpbase.Server

	ServerOption = httpbase.ServerOption
)

var (
	Endpoint = httpbase.Endpoint

	Network = httpbase.Network

	Address = httpbase.Address

	Timeout = httpbase.Timeout

	Middleware = httpbase.Middleware

	TLSConfig = httpbase.TLSConfig

	StrictSlash = httpbase.StrictSlash

	Adapter = httpbase.Adapter
)

func NewServer(opts ...ServerOption) Server {
	s := httpbase.NewServer(opts...)
	return s
}

// DefaultServer has an HTTP3 adapter with default options
func DefaultServer(opts ...ServerOption) Server {
	a := NewAdapter(DefaultOptions)
	opts = append(opts, httpbase.Adapter(a))
	s := httpbase.NewServer(opts...)
	return s
}
