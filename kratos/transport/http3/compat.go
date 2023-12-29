package http3

import (
	khttp "github.com/go-kratos/kratos/v2/transport/http"
)

type (
	// Context is an HTTP Context.
	Context = khttp.Context

	ClientOption = khttp.ClientOption

	FilterFunc = khttp.FilterFunc

	// HandlerFunc defines a function to serve HTTP requests.
	HandlerFunc = khttp.HandlerFunc

	// RouteInfo is an HTTP route info.
	RouteInfo = khttp.RouteInfo

	// WalkRouteFunc is the type of the function called for each route visited by Walk.
	WalkRouteFunc = khttp.WalkRouteFunc

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

	WithTimeout = khttp.WithTimeout

	WithMiddleware = khttp.WithMiddleware

	WithTLSConfig = khttp.WithTLSConfig

	WithRequestEncoder = khttp.WithRequestEncoder

	WithResponseDecoder = khttp.WithResponseDecoder

	WithErrorDecoder = khttp.WithErrorDecoder

	WithUserAgent = khttp.WithUserAgent

	FilterChain = khttp.FilterChain
)
