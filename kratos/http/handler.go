package http

import (
	"net/http"

	khttp "github.com/go-kratos/kratos/v2/transport/http"
)

type Handler interface {
	HandleHTTP(*khttp.Server)
}

type StdHandlerFunc http.HandlerFunc

func (h StdHandlerFunc) Handle(ctx khttp.Context) error {
	w := ctx.Response()
	r := ctx.Request()
	h(w, r)
	return nil
}
