package transport

import (
	"context"

	"github.com/go-kratos/kratos/v2/transport"
)

const (
	KindHTTP3 transport.Kind = "http3"
)

type Validator interface {
	Validate(context.Context) error
}
