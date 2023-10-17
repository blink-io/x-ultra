package generics

import (
	"net/http"
)

type options struct {
	method string
}

type Option func(*options)

func GET() Option {
	return func(o *options) {
		o.method = http.MethodGet
	}
}

func POST() Option {
	return func(o *options) {
		o.method = http.MethodPost
	}
}

func PUT() Option {
	return func(o *options) {
		o.method = http.MethodPut
	}
}

func PATCH() Option {
	return func(o *options) {
		o.method = http.MethodPatch
	}
}

func DELETE() Option {
	return func(o *options) {
		o.method = http.MethodDelete
	}
}

func OPTIONS() Option {
	return func(o *options) {
		o.method = http.MethodOptions
	}
}
