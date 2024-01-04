package nats

import (
	"context"
	"crypto/tls"
	"net"
	"time"

	"github.com/nats-io/nats.go"
)

var _ IConn = (*nats.Conn)(nil)

// IConn is extracted from nats.Conn struct.
type IConn interface {
	RequestMsgWithContext(ctx context.Context, msg *nats.Msg) (*nats.Msg, error)
	RequestWithContext(ctx context.Context, subj string, data []byte) (*nats.Msg, error)
	FlushWithContext(ctx context.Context) error
	JetStream(opts ...nats.JSOpt) (nats.JetStreamContext, error)
	SetDisconnectHandler(dcb nats.ConnHandler)
	SetDisconnectErrHandler(dcb nats.ConnErrHandler)
	DisconnectErrHandler() nats.ConnErrHandler
	SetReconnectHandler(rcb nats.ConnHandler)
	ReconnectHandler() nats.ConnHandler
	SetDiscoveredServersHandler(dscb nats.ConnHandler)
	DiscoveredServersHandler() nats.ConnHandler
	SetClosedHandler(cb nats.ConnHandler)
	ClosedHandler() nats.ConnHandler
	SetErrorHandler(cb nats.ErrHandler)
	ErrorHandler() nats.ErrHandler
	TLSConnectionState() (tls.ConnectionState, error)
	ConnectedUrl() string
	ConnectedUrlRedacted() string
	ConnectedAddr() string
	ConnectedServerId() string
	ConnectedServerName() string
	ConnectedServerVersion() string
	ConnectedClusterName() string
	LastError() error
	Publish(subj string, data []byte) error
	PublishMsg(m *nats.Msg) error
	PublishRequest(subj, reply string, data []byte) error
	RequestMsg(msg *nats.Msg, timeout time.Duration) (*nats.Msg, error)
	Request(subj string, data []byte, timeout time.Duration) (*nats.Msg, error)
	NewInbox() string
	NewRespInbox() string
	Subscribe(subj string, cb nats.MsgHandler) (*nats.Subscription, error)
	ChanSubscribe(subj string, ch chan *nats.Msg) (*nats.Subscription, error)
	ChanQueueSubscribe(subj, group string, ch chan *nats.Msg) (*nats.Subscription, error)
	SubscribeSync(subj string) (*nats.Subscription, error)
	QueueSubscribe(subj, queue string, cb nats.MsgHandler) (*nats.Subscription, error)
	QueueSubscribeSync(subj, queue string) (*nats.Subscription, error)
	QueueSubscribeSyncWithChan(subj, queue string, ch chan *nats.Msg) (*nats.Subscription, error)
	NumSubscriptions() int
	FlushTimeout(timeout time.Duration) (err error)
	RTT() (time.Duration, error)
	Flush() error
	Buffered() (int, error)
	Close()
	IsClosed() bool
	IsReconnecting() bool
	IsConnected() bool
	Drain() error
	IsDraining() bool
	Servers() []string
	DiscoveredServers() []string
	Status() nats.Status
	Stats() nats.Statistics
	MaxPayload() int64
	HeadersSupported() bool
	AuthRequired() bool
	TLSRequired() bool
	Barrier(f func()) error
	GetClientIP() (net.IP, error)
	GetClientID() (uint64, error)
	StatusChanged(statuses ...nats.Status) chan nats.Status
}
