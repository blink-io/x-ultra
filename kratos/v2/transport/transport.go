package transport

import (
	"context"

	"github.com/go-kratos/kratos/v2/transport"
)

const (
	KindHTTP  Kind = transport.KindHTTP
	KindHTTP3 Kind = "http3"
)

type Kind = transport.Kind

type Validator interface {
	Validate(context.Context) error
}
