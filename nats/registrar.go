package nats

import (
	"context"
	"time"

	"github.com/nats-io/nats.go"
)

type RegisterToNATSFunc func(context.Context, ServiceRegistrar) error

type WithRegistrar interface {
	NATSRegistrar(context.Context) RegisterToNATSFunc
}

type ServiceRegistrar interface {
	RequestMsgWithContext(ctx context.Context, msg *nats.Msg) (*nats.Msg, error)

	RequestWithContext(ctx context.Context, subj string, data []byte) (*nats.Msg, error)

	JetStream(opts ...nats.JSOpt) (nats.JetStreamContext, error)

	Publish(subj string, data []byte) error

	PublishMsg(m *nats.Msg) error

	PublishRequest(subj, reply string, data []byte) error

	RequestMsg(msg *nats.Msg, timeout time.Duration) (*nats.Msg, error)

	Request(subj string, data []byte, timeout time.Duration) (*nats.Msg, error)

	Subscribe(subj string, cb nats.MsgHandler) (*nats.Subscription, error)

	ChanSubscribe(subj string, ch chan *nats.Msg) (*nats.Subscription, error)

	ChanQueueSubscribe(subj, group string, ch chan *nats.Msg) (*nats.Subscription, error)

	SubscribeSync(subj string) (*nats.Subscription, error)

	QueueSubscribe(subj, queue string, cb nats.MsgHandler) (*nats.Subscription, error)

	QueueSubscribeSync(subj, queue string) (*nats.Subscription, error)

	QueueSubscribeSyncWithChan(subj, queue string, ch chan *nats.Msg) (*nats.Subscription, error)
}

type RegistrarFunc[S any] func(ServiceRegistrar, S)

type RegistrarFuncWithErr[S any] func(ServiceRegistrar, S) error

type CtxRegistrarFunc[S any] func(context.Context, ServiceRegistrar, S)

type CtxRegistrarFuncWithErr[S any] func(context.Context, ServiceRegistrar, S) error

type Registrar interface {
	RegisterToNATS(context.Context, ServiceRegistrar) error
}

type registrar[S any] struct {
	s S
	f CtxRegistrarFuncWithErr[S]
}

var _ Registrar = (*registrar[any])(nil)

func (h *registrar[S]) RegisterToNATS(ctx context.Context, r ServiceRegistrar) error {
	return h.f(ctx, r, h.s)
}

func NewRegistrar[S any](s S, f RegistrarFunc[S]) Registrar {
	cf := func(ctx context.Context, r ServiceRegistrar, s S) error {
		f(r, s)
		return nil
	}
	return NewCtxRegistrarWithErr(s, cf)
}

func NewRegistrarWithErr[S any](s S, f RegistrarFuncWithErr[S]) Registrar {
	cf := func(ctx context.Context, r ServiceRegistrar, s S) error {
		return f(r, s)
	}
	return NewCtxRegistrarWithErr(s, cf)
}

func NewCtxRegistrar[S any](s S, f CtxRegistrarFunc[S]) Registrar {
	cf := func(ctx context.Context, r ServiceRegistrar, s S) error {
		f(ctx, r, s)
		return nil
	}
	return NewCtxRegistrarWithErr(s, cf)
}

func NewCtxRegistrarWithErr[S any](s S, f CtxRegistrarFuncWithErr[S]) Registrar {
	h := &registrar[S]{
		s: s,
		f: f,
	}
	return h
}
