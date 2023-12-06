package tls

import (
	"time"
)

type options struct {
	timeout time.Duration
}

type Option func(*options)

func applyOptions(ops ...Option) *options {
	opt := &options{
		timeout: 5 * time.Second,
	}
	for _, f := range ops {
		f(opt)
	}
	return opt
}

func Timeout(timeout time.Duration) Option {
	return func(o *options) {
		o.timeout = timeout
	}
}
