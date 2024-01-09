package http

import (
	"context"
	"net/http"

	khttp "github.com/go-kratos/kratos/v2/transport/http"
)

type WithHandler interface {
	HTTPHandler() Handler
}

type Handler interface {
	HandleHTTP(context.Context, ServerRouter)
}

type StdHandlerFunc http.HandlerFunc

func (h StdHandlerFunc) Handle(ctx khttp.Context) error {
	w := ctx.Response()
	r := ctx.Request()
	h(w, r)
	return nil
}

func StdHandler(h http.Handler) khttp.HandlerFunc {
	return func(ctx khttp.Context) error {
		w := ctx.Response()
		r := ctx.Request()
		h.ServeHTTP(w, r)
		return nil
	}
}
