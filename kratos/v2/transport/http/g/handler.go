package g

import (
	khttp "github.com/blink-io/x/kratos/v2/transport/http"
)

type RegistrarFunc[S any] func(khttp.ServerRouter, S)

type Handler[S any] interface {
	HandleHTTP(khttp.ServerRouter)
}

type handler[S any] struct {
	s S
	f RegistrarFunc[S]
}

var _ Handler[any] = (*handler[any])(nil)

func NewHandler[S any](s S, f RegistrarFunc[S]) Handler[S] {
	h := &handler[S]{
		s: s,
		f: f,
	}
	return h
}

func (h handler[S]) HandleHTTP(r khttp.ServerRouter) {
	h.f(r, h.s)
}
