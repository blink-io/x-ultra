package http

import (
	"net/http"

	khttp "github.com/go-kratos/kratos/v2/transport/http"
)

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

type HandlerFunc http.HandlerFunc

func (h HandlerFunc) Handle(ctx khttp.Context) error {
	w := ctx.Response()
	r := ctx.Request()
	h(w, r)
	return nil
}
