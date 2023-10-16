package i18n

import (
	"context"

	"github.com/blink-io/x/i18n"

	"github.com/go-kratos/kratos/v2/middleware"
)

type Options = i18n.Options

var DefaultOptions = i18n.DefaultOptions

func Default() middleware.Middleware {
	return New(DefaultOptions)
}

func New(c *Options) middleware.Middleware {
	return func(h middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			return h(ctx, req)
		}
	}
}
