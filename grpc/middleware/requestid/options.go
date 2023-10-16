package requestid

import (
	"github.com/blink-io/x/requestid"
)

type options = requestid.Options

type Option func(*options)

func initOption(ops ...Option) *options {
	opt := requestid.DefaultOptions
	for _, o := range ops {
		o(opt)
	}
	return opt
}

func Generator(g func() string) Option {
	return func(o *options) {
		if g != nil {
			o.Generator = g
		}
	}
}

func Header(h string) Option {
	return func(o *options) {
		if len(h) > 0 {
			o.Header = h
		}
	}
}
