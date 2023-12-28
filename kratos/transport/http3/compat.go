package http3

import (
	khttp "github.com/go-kratos/kratos/v2/transport/http"
)

// Context is an HTTP Context.
type Context = khttp.Context

type (
	DecodeRequestFunc  = khttp.DecodeRequestFunc
	EncodeResponseFunc = khttp.EncodeResponseFunc
	EncodeErrorFunc    = khttp.EncodeErrorFunc
)

var (
	DefaultRequestVars     = khttp.DefaultRequestVars
	DefaultRequestQuery    = khttp.DefaultRequestQuery
	DefaultRequestDecoder  = khttp.DefaultRequestDecoder
	DefaultResponseEncoder = khttp.DefaultResponseEncoder
	DefaultErrorEncoder    = khttp.DefaultErrorEncoder
)

var (
	NewClient = khttp.NewClient

	WithEndpoint = khttp.WithEndpoint

	WithTransport = khttp.WithTransport
)
