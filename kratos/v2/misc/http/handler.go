package http

import (
	"net/http"

	khttp "github.com/go-kratos/kratos/v2/transport/http"
)

type Router interface {
	GET(path string, h khttp.HandlerFunc, m ...khttp.FilterFunc)

	HEAD(path string, h khttp.HandlerFunc, m ...khttp.FilterFunc)

	POST(path string, h khttp.HandlerFunc, m ...khttp.FilterFunc)

	PUT(path string, h khttp.HandlerFunc, m ...khttp.FilterFunc)

	PATCH(path string, h khttp.HandlerFunc, m ...khttp.FilterFunc)

	DELETE(path string, h khttp.HandlerFunc, m ...khttp.FilterFunc)

	CONNECT(path string, h khttp.HandlerFunc, m ...khttp.FilterFunc)

	OPTIONS(path string, h khttp.HandlerFunc, m ...khttp.FilterFunc)

	TRACE(path string, h khttp.HandlerFunc, m ...khttp.FilterFunc)

	Handle(method, relativePath string, h khttp.HandlerFunc, filters ...khttp.FilterFunc)

	Group(prefix string, filters ...khttp.FilterFunc) *khttp.Router
}

var _ Router = (*khttp.Router)(nil)

type RouteRegistrar interface {
	Route(prefix string, filters ...khttp.FilterFunc) *khttp.Router

	Handle(path string, h http.Handler)

	HandlePrefix(prefix string, h http.Handler)

	HandleFunc(path string, h http.HandlerFunc)

	HandleHeader(key, val string, h http.HandlerFunc)
}

type Handler interface {
	HandleHTTP(RouteRegistrar)
}

type StdHandlerFunc http.HandlerFunc

func (h StdHandlerFunc) Handle(ctx khttp.Context) error {
	w := ctx.Response()
	r := ctx.Request()
	h(w, r)
	return nil
}
