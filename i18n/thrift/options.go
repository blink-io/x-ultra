package thrift

import "crypto/tls"

type options struct {
	protocol   Protocol
	useHTTP    bool
	framed     bool
	buffered   bool
	bufferSize int
	tlsConfig  *tls.Config
}

func applyOptions(ops ...Option) *options {
	opt := new(options)
	for _, o := range ops {
		o(opt)
	}
	return opt
}

type Option func(*options)

func WithTProtocol(protocol Protocol) Option {
	return func(o *options) {
		o.protocol = protocol
	}
}

func WithTFramed(framed bool) Option {
	return func(o *options) {
		o.framed = framed
	}
}

func WithUseHTTP() Option {
	return func(o *options) {
		o.useHTTP = true
	}
}

func WithBuffered(bufferSize int) Option {
	return func(o *options) {
		o.buffered = true
		o.bufferSize = bufferSize
	}
}

func WithTLSConfig(tlsConfig *tls.Config) Option {
	return func(o *options) {
		o.tlsConfig = tlsConfig
	}
}
