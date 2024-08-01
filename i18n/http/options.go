package http

import (
	"net/http"

	"github.com/blink-io/x/i18n"
)

type options struct {
	client *http.Client

	requestFunc func(*http.Client, string) (*http.Request, error)

	extractFunc func(string) string
}

type Option func(*options)

func applyOptions(ops ...Option) *options {
	opts := &options{
		client: &http.Client{Timeout: i18n.DefaultTimeout},
		extractFunc: func(s string) string {
			return s
		},
	}
	for _, op := range ops {
		op(opts)
	}
	return opts
}

func WithHTTPClient(client *http.Client) Option {
	return func(o *options) {
		o.client = client
	}
}

func WithRequestFunc(requestFunc func(*http.Client, string) (*http.Request, error)) Option {
	return func(o *options) {
		o.requestFunc = requestFunc
	}
}

func WithExtractFunc(extractFunc func(string) string) Option {
	return func(o *options) {
		o.extractFunc = extractFunc
	}
}
