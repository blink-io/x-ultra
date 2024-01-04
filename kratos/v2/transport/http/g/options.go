package g

import (
	"net/http"
)

type options struct {
	method string
}

type DoOption func(*options)

func applyOptions(ops ...DoOption) *options {
	opt := new(options)
	for _, o := range ops {
		o(opt)
	}
	if len(opt.method) == 0 {
		opt.method = http.MethodGet
	}
	return opt
}
