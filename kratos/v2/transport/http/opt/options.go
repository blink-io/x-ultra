package opt

import (
	"net/http"
)

type keyValue struct {
	key   string
	value string
}

type options struct {
	method      string
	statusCode  int
	ahs         []*keyValue
	shs         map[string]string
	skipQuery   bool
	skipVars    bool
	skipReqBody bool
	skipResBody bool
}

type DoOption func(*options)

func applyOptions(ops ...DoOption) *options {
	opt := &options{
		ahs: make([]*keyValue, 0),
		shs: make(map[string]string),
	}
	for _, o := range ops {
		o(opt)
	}
	if len(opt.method) == 0 {
		opt.method = http.MethodGet
	}
	if opt.statusCode == 0 {
		opt.statusCode = http.StatusOK
	}
	return opt
}

func StatusCode(statusCode int) DoOption {
	return func(o *options) {
		o.statusCode = statusCode
	}
}

func AddHeader(key string, value string) DoOption {
	return func(o *options) {
		o.ahs = append(o.ahs, &keyValue{key, value})
	}
}

func SetHeader(key string, value string) DoOption {
	return func(o *options) {
		o.shs[key] = value
	}
}

func SkipVars() DoOption {
	return func(o *options) {
		o.skipVars = true
	}
}

func SkipQuery() DoOption {
	return func(o *options) {
		o.skipQuery = true
	}
}

func SkipReqBody() DoOption {
	return func(o *options) {
		o.skipReqBody = true
	}
}

func SkipResBody() DoOption {
	return func(o *options) {
		o.skipResBody = true
	}
}
