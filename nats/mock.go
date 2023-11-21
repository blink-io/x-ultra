package nats

import (
	"context"
	"crypto/tls"
	"log/slog"
	"net"
	"time"

	"github.com/nats-io/nats.go"
)

type mockConn struct {
	serverID      string
	serverName    string
	serverVersion string

	addr        string
	url         string
	clusterName string
	newInbox    string
	respInbox   string

	servers []string

	maxPayload       int64
	headersSupported bool
	authRequired     bool
	tlsRequired      bool
	isClosed         bool
	isReconnecting   bool
	isDraining       bool
	isConnected      bool

	clientID uint64
	clientIP net.IP

	rtt          time.Duration
	flushTimeout time.Duration

	bufferd          int
	numSubscriptions int

	status nats.Status

	stats nats.Statistics

	reqMsg *nats.Msg
	rwcMsg *nats.Msg

	tlsConnState tls.ConnectionState

	qsSub   *nats.Subscription
	cqsSub  *nats.Subscription
	csSub   *nats.Subscription
	sSub    *nats.Subscription
	ssSub   *nats.Subscription
	qssSub  *nats.Subscription
	qsscSub *nats.Subscription

	ceb  nats.ConnErrHandler
	rcb  nats.ConnHandler
	dscb nats.ConnHandler
	eb   nats.ErrHandler
	cb   nats.ConnHandler
	dcb  nats.ConnHandler

	jsCtx nats.JetStreamContext

	chStatus chan nats.Status
}

func newMockConn() IConn {
	mcc := &mockConn{}
	return mcc
}

func (m *mockConn) RequestMsgWithContext(ctx context.Context, msg *nats.Msg) (*nats.Msg, error) {
	return m.rwcMsg, nil
}

func (m *mockConn) RequestWithContext(ctx context.Context, subject string, data []byte) (*nats.Msg, error) {
	return m.rwcMsg, nil
}

func (m *mockConn) FlushWithContext(ctx context.Context) error {
	return nil
}

func (m *mockConn) RemoveMsgFilter(subject string) {
	slog.Info("invoke RemoveMsgFilter", slog.String("subject", subject))
}

func (m *mockConn) CloseTCPConn() {
	slog.Info("invoke CloseTCPConn")
}

func (m *mockConn) JetStream(opts ...nats.JSOpt) (nats.JetStreamContext, error) {
	return m.jsCtx, nil
}

func (m *mockConn) SetDisconnectHandler(dcb nats.ConnHandler) {
	m.dcb = dcb
}

func (m *mockConn) SetDisconnectErrHandler(ceb nats.ConnErrHandler) {
	m.ceb = ceb
}

func (m *mockConn) DisconnectErrHandler() nats.ConnErrHandler {
	return m.ceb
}

func (m *mockConn) SetReconnectHandler(rcb nats.ConnHandler) {
	m.rcb = rcb
}

func (m *mockConn) ReconnectHandler() nats.ConnHandler {
	return nil
}

func (m *mockConn) SetDiscoveredServersHandler(dscb nats.ConnHandler) {
	m.dscb = dscb
}

func (m *mockConn) DiscoveredServersHandler() nats.ConnHandler {
	return m.dscb
}

func (m *mockConn) SetClosedHandler(cb nats.ConnHandler) {
	m.cb = cb
}

func (m *mockConn) ClosedHandler() nats.ConnHandler {
	return m.cb
}

func (m *mockConn) SetErrorHandler(eb nats.ErrHandler) {
	m.eb = eb
}

func (m *mockConn) ErrorHandler() nats.ErrHandler {
	return m.eb
}

func (m *mockConn) TLSConnectionState() (tls.ConnectionState, error) {
	return m.tlsConnState, nil
}

func (m *mockConn) ConnectedUrl() string {
	return m.url
}

func (m *mockConn) ConnectedUrlRedacted() string {
	return m.url
}

func (m *mockConn) ConnectedAddr() string {
	return m.addr
}

func (m *mockConn) ConnectedServerId() string {
	return m.serverID
}

func (m *mockConn) ConnectedServerName() string {
	return m.serverName
}

func (m *mockConn) ConnectedServerVersion() string {
	return m.serverVersion
}

func (m *mockConn) ConnectedClusterName() string {
	return m.clusterName
}

func (m *mockConn) LastError() error {
	return nil
}

func (m *mockConn) Publish(subject string, data []byte) error {
	return nil
}

func (m *mockConn) PublishMsg(msg *nats.Msg) error {
	return nil
}

func (m *mockConn) PublishRequest(subj, reply string, data []byte) error {
	return nil
}

func (m *mockConn) RequestMsg(msg *nats.Msg, timeout time.Duration) (*nats.Msg, error) {
	return m.reqMsg, nil
}

func (m *mockConn) Request(subject string, data []byte, timeout time.Duration) (*nats.Msg, error) {
	return m.reqMsg, nil
}

func (m *mockConn) NewInbox() string {
	return m.newInbox
}

func (m *mockConn) NewRespInbox() string {
	return m.respInbox
}

func (m *mockConn) Subscribe(subject string, cb nats.MsgHandler) (*nats.Subscription, error) {
	return m.sSub, nil
}

func (m *mockConn) ChanSubscribe(subject string, ch chan *nats.Msg) (*nats.Subscription, error) {
	return m.csSub, nil
}

func (m *mockConn) ChanQueueSubscribe(subject, group string, chn chan *nats.Msg) (*nats.Subscription, error) {
	return m.cqsSub, nil
}

func (m *mockConn) SubscribeSync(subject string) (*nats.Subscription, error) {
	return m.ssSub, nil
}

func (m *mockConn) QueueSubscribe(subject, queue string, cb nats.MsgHandler) (*nats.Subscription, error) {
	return m.qsSub, nil
}

func (m *mockConn) QueueSubscribeSync(subject, queue string) (*nats.Subscription, error) {
	return m.qssSub, nil
}

func (m *mockConn) QueueSubscribeSyncWithChan(subject, queue string, chn chan *nats.Msg) (*nats.Subscription, error) {
	return m.qsscSub, nil
}

func (m *mockConn) NumSubscriptions() int {
	return m.numSubscriptions
}

func (m *mockConn) FlushTimeout(timeout time.Duration) error {
	return nil
}

func (m *mockConn) RTT() (time.Duration, error) {
	return m.rtt, nil
}

func (m *mockConn) Flush() error {
	return nil
}

func (m *mockConn) Buffered() (int, error) {
	return m.bufferd, nil
}

func (m *mockConn) Close() {
	slog.Info("invoke Close")
}

func (m *mockConn) IsClosed() bool {
	return m.isClosed
}

func (m *mockConn) IsReconnecting() bool {
	return m.isReconnecting
}

func (m *mockConn) IsConnected() bool {
	return m.isConnected
}

func (m *mockConn) Drain() error {
	return nil
}

func (m *mockConn) IsDraining() bool {
	return m.isDraining
}

func (m *mockConn) Servers() []string {
	return m.servers
}

func (m *mockConn) DiscoveredServers() []string {
	return m.servers
}

func (m *mockConn) Status() nats.Status {
	return m.status
}

func (m *mockConn) Stats() nats.Statistics {
	return m.stats
}

func (m *mockConn) MaxPayload() int64 {
	return m.maxPayload
}

func (m *mockConn) HeadersSupported() bool {
	return m.headersSupported
}

func (m *mockConn) AuthRequired() bool {
	return m.authRequired
}

func (m *mockConn) TLSRequired() bool {
	return m.tlsRequired
}

func (m *mockConn) Barrier(f func()) error {
	return nil
}

func (m *mockConn) GetClientIP() (net.IP, error) {
	return m.clientIP, nil
}

func (m *mockConn) GetClientID() (uint64, error) {
	return m.clientID, nil
}

func (m *mockConn) StatusChanged(statuses ...nats.Status) chan nats.Status {
	return m.chStatus
}
