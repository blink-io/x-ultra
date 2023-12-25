package thrift

import (
	"sync"

	"github.com/apache/thrift/lib/go/thrift"
)

type TQUIC struct {
}

// TServerQUIC defines Thrift server over QUIC
type TServerQUIC struct {

	// Protects the interrupted value to make it thread safe.
	mu sync.RWMutex
}

func (T *TServerQUIC) Listen() error {
	return nil
}

func (T *TServerQUIC) Accept() (thrift.TTransport, error) {
	//TODO implement me
	panic("implement me")
}

func (T *TServerQUIC) Close() error {
	return nil
}

func (T *TServerQUIC) Interrupt() error {
	return nil
}

var _ thrift.TServerTransport = (*TServerQUIC)(nil)
