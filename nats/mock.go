package nats

import (
	"context"
	"crypto/tls"
	"log/slog"
	"net"
	"time"

	"github.com/jellydator/ttlcache/v3"
	"github.com/nats-io/nats.go"
)

type MockConn struct {
	Cache            *ttlcache.Cache[string, any]
	ServerIDVar      string
	ServerNameVar    string
	ServerVersionVar string

	AddrVar         string
	URLVar          string
	URLRedactedVar  string
	ClusterNameVar  string
	NewInboxVar     string
	NewRespInboxVar string

	ServersVar []string

	MaxPayloadVar       int64
	HeadersSupportedVar bool
	AuthRequiredVar     bool
	TLSRequiredVar      bool
	IsClosedVar         bool
	IsReconnectingVar   bool
	IsDrainingVar       bool
	IsConnectedVar      bool

	ClientIDVar uint64
	ClientIPVar net.IP

	RTTVar       time.Duration
	flushTimeout time.Duration

	BufferedVar         int
	NumSubscriptionsVar int

	StatusVar nats.Status

	StatsVar nats.Statistics

	RequestVar               *nats.Msg
	RequestMsgVar            *nats.Msg
	RequestWithContextVar    *nats.Msg
	RequestMsgWithContextVar *nats.Msg

	TLSConnState tls.ConnectionState

	QueueSubscribeVar             *nats.Subscription
	ChanQueueSubscribeVar         *nats.Subscription
	ChanSubscribeVar              *nats.Subscription
	SubscribeVar                  *nats.Subscription
	SubscribeSyncVar              *nats.Subscription
	QueueSubscribeSyncVar         *nats.Subscription
	QueueSubscribeSyncWithChanVar *nats.Subscription

	DisconnectErrHandlerVar     nats.ConnErrHandler
	ReconnectHandlerVar         nats.ConnHandler
	DiscoveredServersHandlerVar nats.ConnHandler
	ErrorHandlerVar             nats.ErrHandler
	ClosedHandlerVar            nats.ConnHandler
	DisconnectHandlerVar        nats.ConnHandler

	JSCtx nats.JetStreamContext

	StatusChangedVar chan nats.Status

	BufferdErr        error
	BarrierErr        error
	DrainErr          error
	FlushErr          error
	FlushTimeoutErr   error
	PublishRequestErr error

	LastErrorErr  error
	PublishErr    error
	PublishMsgErr error

	FlushWithContextErr           error
	RequestWithContextErr         error
	RequestMsgWithContextErr      error
	JSCtxErr                      error
	RequestMsgErr                 error
	tlsConnStateErr               error
	RequestErr                    error
	SubscribeErr                  error
	SubscribeSyncErr              error
	ChanSubscribeErr              error
	ChanQueueSubscribeErr         error
	QueueSubscribeErr             error
	QueueSubscribeSyncErr         error
	QueueSubscribeSyncWithChanErr error
}

func NewMockConn() *MockConn {
	mcc := &MockConn{
		Cache: ttlcache.New[string, any](
			ttlcache.WithVersion[string, any](true),
		),
	}
	return mcc
}

func (m *MockConn) RequestMsgWithContext(ctx context.Context, msg *nats.Msg) (*nats.Msg, error) {
	return m.RequestMsgWithContextVar, m.RequestMsgWithContextErr
}

func (m *MockConn) RequestWithContext(ctx context.Context, subject string, data []byte) (*nats.Msg, error) {
	return m.RequestWithContextVar, m.RequestWithContextErr
}

func (m *MockConn) FlushWithContext(ctx context.Context) error {
	return m.FlushWithContextErr
}

func (m *MockConn) RemoveMsgFilter(subject string) {
	slog.Info("invoke RemoveMsgFilter", slog.String("subject", subject))
}

func (m *MockConn) CloseTCPConn() {
	slog.Info("invoke CloseTCPConn")
}

func (m *MockConn) JetStream(opts ...nats.JSOpt) (nats.JetStreamContext, error) {
	return m.JSCtx, m.JSCtxErr
}

func (m *MockConn) SetDisconnectHandler(dcb nats.ConnHandler) {
	m.DisconnectHandlerVar = dcb
}

func (m *MockConn) SetDisconnectErrHandler(ceb nats.ConnErrHandler) {
	m.DisconnectErrHandlerVar = ceb
}

func (m *MockConn) DisconnectErrHandler() nats.ConnErrHandler {
	return m.DisconnectErrHandlerVar
}

func (m *MockConn) SetReconnectHandler(rcb nats.ConnHandler) {
	m.ReconnectHandlerVar = rcb
}

func (m *MockConn) ReconnectHandler() nats.ConnHandler {
	return m.ReconnectHandlerVar
}

func (m *MockConn) SetDiscoveredServersHandler(dscb nats.ConnHandler) {
	m.DiscoveredServersHandlerVar = dscb
}

func (m *MockConn) DiscoveredServersHandler() nats.ConnHandler {
	return m.DiscoveredServersHandlerVar
}

func (m *MockConn) SetClosedHandler(cb nats.ConnHandler) {
	m.ClosedHandlerVar = cb
}

func (m *MockConn) ClosedHandler() nats.ConnHandler {
	return m.ClosedHandlerVar
}

func (m *MockConn) SetErrorHandler(eb nats.ErrHandler) {
	m.ErrorHandlerVar = eb
}

func (m *MockConn) ErrorHandler() nats.ErrHandler {
	return m.ErrorHandlerVar
}

func (m *MockConn) TLSConnectionState() (tls.ConnectionState, error) {
	return m.TLSConnState, m.tlsConnStateErr
}

func (m *MockConn) ConnectedUrl() string {
	return m.URLVar
}

func (m *MockConn) ConnectedUrlRedacted() string {
	return m.URLRedactedVar
}

func (m *MockConn) ConnectedAddr() string {
	return m.AddrVar
}

func (m *MockConn) ConnectedServerId() string {
	return m.ServerIDVar
}

func (m *MockConn) ConnectedServerName() string {
	return m.ServerNameVar
}

func (m *MockConn) ConnectedServerVersion() string {
	return m.ServerVersionVar
}

func (m *MockConn) ConnectedClusterName() string {
	return m.ClusterNameVar
}

func (m *MockConn) LastError() error {
	return m.LastErrorErr
}

func (m *MockConn) Publish(subject string, data []byte) error {
	return m.PublishErr
}

func (m *MockConn) PublishMsg(msg *nats.Msg) error {
	return m.PublishMsgErr
}

func (m *MockConn) PublishRequest(subj, reply string, data []byte) error {
	return m.PublishRequestErr
}

func (m *MockConn) RequestMsg(msg *nats.Msg, timeout time.Duration) (*nats.Msg, error) {
	return m.RequestMsgVar, m.RequestMsgErr
}

func (m *MockConn) Request(subject string, data []byte, timeout time.Duration) (*nats.Msg, error) {
	return m.RequestVar, m.RequestErr
}

func (m *MockConn) NewInbox() string {
	return m.NewInboxVar
}

func (m *MockConn) NewRespInbox() string {
	return m.NewRespInboxVar
}

func (m *MockConn) Subscribe(subject string, cb nats.MsgHandler) (*nats.Subscription, error) {
	return m.SubscribeVar, m.SubscribeErr
}

func (m *MockConn) ChanSubscribe(subject string, ch chan *nats.Msg) (*nats.Subscription, error) {
	return m.ChanSubscribeVar, m.ChanSubscribeErr
}

func (m *MockConn) ChanQueueSubscribe(subject, group string, chn chan *nats.Msg) (*nats.Subscription, error) {
	return m.ChanQueueSubscribeVar, m.ChanQueueSubscribeErr
}

func (m *MockConn) SubscribeSync(subject string) (*nats.Subscription, error) {
	return m.SubscribeSyncVar, m.SubscribeSyncErr
}

func (m *MockConn) QueueSubscribe(subject, queue string, cb nats.MsgHandler) (*nats.Subscription, error) {
	return m.QueueSubscribeVar, m.QueueSubscribeErr
}

func (m *MockConn) QueueSubscribeSync(subject, queue string) (*nats.Subscription, error) {
	return m.QueueSubscribeSyncVar, m.QueueSubscribeSyncErr
}

func (m *MockConn) QueueSubscribeSyncWithChan(subject, queue string, chn chan *nats.Msg) (*nats.Subscription, error) {
	return m.QueueSubscribeSyncWithChanVar, m.QueueSubscribeSyncWithChanErr
}

func (m *MockConn) NumSubscriptions() int {
	return m.NumSubscriptionsVar
}

func (m *MockConn) FlushTimeout(timeout time.Duration) error {
	return m.FlushTimeoutErr
}

func (m *MockConn) RTT() (time.Duration, error) {
	return m.RTTVar, nil
}

func (m *MockConn) Flush() error {
	return m.FlushErr
}

func (m *MockConn) Buffered() (int, error) {
	return m.BufferedVar, m.BufferdErr
}

func (m *MockConn) Close() {
	slog.Info("invoke Close")
}

func (m *MockConn) IsClosed() bool {
	return m.IsClosedVar
}

func (m *MockConn) IsReconnecting() bool {
	return m.IsReconnectingVar
}

func (m *MockConn) IsConnected() bool {
	return m.IsConnectedVar
}

func (m *MockConn) Drain() error {
	return m.DrainErr
}

func (m *MockConn) IsDraining() bool {
	return m.IsDrainingVar
}

func (m *MockConn) Servers() []string {
	return m.ServersVar
}

func (m *MockConn) DiscoveredServers() []string {
	return m.ServersVar
}

func (m *MockConn) Status() nats.Status {
	return m.StatusVar
}

func (m *MockConn) Stats() nats.Statistics {
	return m.StatsVar
}

func (m *MockConn) MaxPayload() int64 {
	return m.MaxPayloadVar
}

func (m *MockConn) HeadersSupported() bool {
	return m.HeadersSupportedVar
}

func (m *MockConn) AuthRequired() bool {
	return m.AuthRequiredVar
}

func (m *MockConn) TLSRequired() bool {
	return m.TLSRequiredVar
}

func (m *MockConn) Barrier(f func()) error {
	return m.BarrierErr
}

func (m *MockConn) GetClientIP() (net.IP, error) {
	return m.ClientIPVar, nil
}

func (m *MockConn) GetClientID() (uint64, error) {
	return m.ClientIDVar, nil
}

func (m *MockConn) StatusChanged(statuses ...nats.Status) chan nats.Status {
	return m.StatusChangedVar
}
