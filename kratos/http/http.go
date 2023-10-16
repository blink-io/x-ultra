package http

import (
	"net/http"

	khttp "github.com/go-kratos/kratos/v2/transport/http"
)

type HandlerFunc func(http.ResponseWriter, *http.Request)

func (h HandlerFunc) Handle(ctx khttp.Context) error {
	w := ctx.Response()
	r := ctx.Request()
	h(w, r)
	return nil
}
