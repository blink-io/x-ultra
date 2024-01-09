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

func StdHandlerFunc(h http.HandlerFunc) khttp.HandlerFunc {
	return func(ctx khttp.Context) error {
		w := ctx.Response()
		r := ctx.Request()
		h(w, r)
		return nil
	}
}

func StdHandler(h http.Handler) khttp.HandlerFunc {
	return func(ctx khttp.Context) error {
		w := ctx.Response()
		r := ctx.Request()
		h.ServeHTTP(w, r)
		return nil
	}
}

func RouteHandlerFunc(r Router, f khttp.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		ctx := &wrapper{router: r}
		ctx.Reset(res, req)
		err := f(ctx)
		if err != nil {
			r.server().EncodeError()(res, req, err)
		}
	})
}
