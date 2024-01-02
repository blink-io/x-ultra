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
)

func NewServer(opts ...ServerOption) Server {
	a := NewAdapter(DefaultOptions)
	opts = append(opts, httpbase.Adapter(a))
	s := httpbase.NewServer(opts...)
	return s
}
