package http3

import (
	"github.com/go-kratos/kratos/v2/transport"
	_ "github.com/quic-go/quic-go"
)

const (
	KindHTTP3 transport.Kind = "http3"
)

var _ transport.Transporter = (*http3)(nil)

type http3 struct {
}

func (h *http3) Kind() transport.Kind {
	return KindHTTP3
}

func (h *http3) Endpoint() string {
	//TODO implement me
	panic("implement me")
}

func (h *http3) Operation() string {
	//TODO implement me
	panic("implement me")
}

func (h *http3) RequestHeader() transport.Header {
	//TODO implement me
	panic("implement me")
}

func (h *http3) ReplyHeader() transport.Header {
	//TODO implement me
	panic("implement me")
}
