package kvstore

import (
	"context"
)

const testStoreName = "mock"

func newStore(ctx context.Context, endpoints []string, options Config) (Store, error) {
	cfg, ok := options.(*Config)
	if !ok && cfg != nil {
		return nil, &InvalidConfigurationError{Store: testStoreName, Config: options}
	}

	return New(ctx, endpoints, cfg)
}

type Mock struct {
	cfg *Config
}

// New creates a new Mock client.
//
//nolint:gocritic
func New(_ context.Context, _ []string, cfg *Config) (*Mock, error) {
	return &Mock{cfg: cfg}, nil
}

func (m Mock) Put(_ context.Context, _ string, _ []byte, _ *WriteOptions) error {
	panic("implement me")
}

func (m Mock) Get(_ context.Context, _ string, _ *ReadOptions) (*KVPair, error) {
	panic("implement me")
}

func (m Mock) Delete(_ context.Context, _ string) error {
	panic("implement me")
}

func (m Mock) Exists(_ context.Context, _ string, _ *ReadOptions) (bool, error) {
	panic("implement me")
}

func (m Mock) Watch(_ context.Context, _ string, _ *ReadOptions) (<-chan *KVPair, error) {
	panic("implement me")
}

func (m Mock) WatchTree(_ context.Context, _ string, _ *ReadOptions) (<-chan []*KVPair, error) {
	panic("implement me")
}

func (m Mock) NewLock(_ context.Context, _ string, _ *LockOptions) (Locker, error) {
	panic("implement me")
}

func (m Mock) List(_ context.Context, _ string, _ *ReadOptions) ([]*KVPair, error) {
	panic("implement me")
}

func (m Mock) DeleteTree(_ context.Context, _ string) error {
	panic("implement me")
}

func (m Mock) AtomicPut(_ context.Context, _ string, _ []byte, _ *KVPair, _ *WriteOptions) (bool, *KVPair, error) {
	panic("implement me")
}

func (m Mock) AtomicDelete(_ context.Context, _ string, _ *KVPair) (bool, error) {
	panic("implement me")
}

func (m Mock) Close() error {
	panic("implement me")
}
