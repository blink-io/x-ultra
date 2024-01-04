package g

import (
	"google.golang.org/grpc"
)

type RegistrarFunc[S any] func(grpc.ServiceRegistrar, S)

type Handler[S any] interface {
	HandleGRPC(grpc.ServiceRegistrar)
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

func (h handler[S]) HandleGRPC(r grpc.ServiceRegistrar) {
	h.f(r, h.s)
}
