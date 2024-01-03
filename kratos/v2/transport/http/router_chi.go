package http

import (
	"github.com/go-chi/chi/v5"
)

type chiRouter struct {
	chi.Router
}

func newChiRouter() *chiRouter {
	r := &chiRouter{
		Router: chi.NewRouter(),
	}
	return r
}
