package http3

import (
	"net/http"

	khttp3 "github.com/blink-io/x/kratos/transport/http3"
)

type Router interface {
	GET(path string, h khttp3.HandlerFunc, m ...khttp3.FilterFunc)

	HEAD(path string, h khttp3.HandlerFunc, m ...khttp3.FilterFunc)

	POST(path string, h khttp3.HandlerFunc, m ...khttp3.FilterFunc)

	PUT(path string, h khttp3.HandlerFunc, m ...khttp3.FilterFunc)

	PATCH(path string, h khttp3.HandlerFunc, m ...khttp3.FilterFunc)

	DELETE(path string, h khttp3.HandlerFunc, m ...khttp3.FilterFunc)

	CONNECT(path string, h khttp3.HandlerFunc, m ...khttp3.FilterFunc)

	OPTIONS(path string, h khttp3.HandlerFunc, m ...khttp3.FilterFunc)

	TRACE(path string, h khttp3.HandlerFunc, m ...khttp3.FilterFunc)

	Handle(method, relativePath string, h khttp3.HandlerFunc, filters ...khttp3.FilterFunc)

	Group(prefix string, filters ...khttp3.FilterFunc) *khttp3.Router
}

var _ Router = (*khttp3.Router)(nil)

type RouteRegistrar interface {
	Route(prefix string, filters ...khttp3.FilterFunc) *khttp3.Router

	Handle(path string, h http.Handler)

	HandlePrefix(prefix string, h http.Handler)

	HandleFunc(path string, h http.HandlerFunc)

	HandleHeader(key, val string, h http.HandlerFunc)
}

type Handler interface {
	HandleHTTP3(RouteRegistrar)
}

type StdHandlerFunc http.HandlerFunc

func (h StdHandlerFunc) Handle(ctx khttp3.Context) error {
	w := ctx.Response()
	r := ctx.Request()
	h(w, r)
	return nil
}
