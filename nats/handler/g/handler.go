package g

import (
	"context"
	"time"

	"github.com/nats-io/nats.go"
)

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

type CtxRegistrarFunc[S any] func(context.Context, ServiceRegistrar, S)

type RegistrarErrFunc[S any] func(ServiceRegistrar, S) error

type CtxRegistrarErrFunc[S any] func(context.Context, ServiceRegistrar, S) error

type Handler interface {
	HandleNATS(context.Context, ServiceRegistrar) error
}

type handler[S any] struct {
	s S
	f CtxRegistrarErrFunc[S]
}

var _ Handler = (*handler[any])(nil)

func (h handler[S]) HandleNATS(ctx context.Context, r ServiceRegistrar) error {
	return h.f(ctx, r, h.s)
}

func NewHandler[S any](s S, f RegistrarFunc[S]) Handler {
	cf := func(ctx context.Context, r ServiceRegistrar, s S) error {
		f(r, s)
		return nil
	}
	return NewCtxErrHandler(s, cf)
}

func NewCtxHandler[S any](s S, f CtxRegistrarFunc[S]) Handler {
	cf := func(ctx context.Context, r ServiceRegistrar, s S) error {
		f(ctx, r, s)
		return nil
	}
	return NewCtxErrHandler(s, cf)
}

func NewErrHandler[S any](s S, f RegistrarErrFunc[S]) Handler {
	cf := func(ctx context.Context, r ServiceRegistrar, s S) error {
		return f(r, s)
	}
	return NewCtxErrHandler(s, cf)
}

func NewCtxErrHandler[S any](s S, f CtxRegistrarErrFunc[S]) Handler {
	h := &handler[S]{
		s: s,
		f: f,
	}
	return h
}
