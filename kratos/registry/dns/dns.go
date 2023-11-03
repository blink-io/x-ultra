package dns

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/registry"
	"github.com/miekg/dns"
)

var (
	_ registry.Registrar = (*Registry)(nil)
	_ registry.Discovery = (*Registry)(nil)
)

// Option is etcd registry option.
type Option func(o *options)

type options struct {
	ctx       context.Context
	namespace string
	ttl       time.Duration
	maxRetry  int
}

// Namespace with registry namespace.
func Namespace(ns string) Option {
	return func(o *options) { o.namespace = ns }
}

// RegisterTTL with register ttl.
func RegisterTTL(ttl time.Duration) Option {
	return func(o *options) { o.ttl = ttl }
}

func MaxRetry(num int) Option {
	return func(o *options) { o.maxRetry = num }
}

type Registry struct {
	client *dns.Client
	opts   *options
}

func New(client *dns.Client, opts ...Option) (r *Registry) {
	op := &options{
		ctx:       context.Background(),
		namespace: "/microservices",
		ttl:       time.Second * 15,
		maxRetry:  5,
	}
	for _, o := range opts {
		o(op)
	}
	return &Registry{
		opts:   op,
		client: client,
	}
}

func (r *Registry) GetService(ctx context.Context, serviceName string) ([]*registry.ServiceInstance, error) {
	//TODO implement me
	panic("implement me")
}

func (r *Registry) Watch(ctx context.Context, serviceName string) (registry.Watcher, error) {
	//TODO implement me
	panic("implement me")
}

func (r *Registry) Register(ctx context.Context, service *registry.ServiceInstance) error {
	return nil
}

func (r *Registry) Deregister(ctx context.Context, service *registry.ServiceInstance) error {
	//TODO implement me
	panic("implement me")
}
