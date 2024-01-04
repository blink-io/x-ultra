package transport

import (
	"context"

	"github.com/go-kratos/kratos/v2/transport"
)

const (
	KindHTTP  Kind = transport.KindHTTP
	KindHTTP3 Kind = "http3"
)

var (
	NewClientContext = transport.NewClientContext

	FromClientContext = transport.FromClientContext

	NewServerContext = transport.NewServerContext

	FromServerContext = transport.FromServerContext
)

type (
	Kind = transport.Kind

	Server = transport.Server

	Transporter = transport.Transporter

	Endpointer = transport.Endpointer

	Header = transport.Header
)

type Validator interface {
	Validate(context.Context) error
}
