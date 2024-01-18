package otel

import (
	"context"
	"crypto/tls"
	"net"
	"time"

	xnats "github.com/blink-io/x/nats"
	"github.com/nats-io/nats.go"
)

type otelConn struct {
	cc xnats.IConn

	opts *options
}

func New(cc xnats.IConn) xnats.IConn {
	occ := &otelConn{}
	return occ
}

func (o *otelConn) RequestMsgWithContext(ctx context.Context, msg *nats.Msg) (*nats.Msg, error) {
	//TODO implement me
	panic("implement me")
}

func (o *otelConn) RequestWithContext(ctx context.Context, subj string, data []byte) (*nats.Msg, error) {
	//TODO implement me
	panic("implement me")
}

func (o *otelConn) FlushWithContext(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (o *otelConn) JetStream(opts ...nats.JSOpt) (nats.JetStreamContext, error) {
	//TODO implement me
	panic("implement me")
}

func (o *otelConn) SetDisconnectHandler(dcb nats.ConnHandler) {
	//TODO implement me
	panic("implement me")
}

func (o *otelConn) SetDisconnectErrHandler(dcb nats.ConnErrHandler) {
	//TODO implement me
	panic("implement me")
}

func (o *otelConn) DisconnectErrHandler() nats.ConnErrHandler {
	//TODO implement me
	panic("implement me")
}

func (o *otelConn) SetReconnectHandler(rcb nats.ConnHandler) {
	//TODO implement me
	panic("implement me")
}

func (o *otelConn) ReconnectHandler() nats.ConnHandler {
	//TODO implement me
	panic("implement me")
}

func (o *otelConn) SetDiscoveredServersHandler(dscb nats.ConnHandler) {
	//TODO implement me
	panic("implement me")
}

func (o *otelConn) DiscoveredServersHandler() nats.ConnHandler {
	//TODO implement me
	panic("implement me")
}

func (o *otelConn) SetClosedHandler(cb nats.ConnHandler) {
	//TODO implement me
	panic("implement me")
}

func (o *otelConn) ClosedHandler() nats.ConnHandler {
	//TODO implement me
	panic("implement me")
}

func (o *otelConn) SetErrorHandler(cb nats.ErrHandler) {
	//TODO implement me
	panic("implement me")
}

func (o *otelConn) ErrorHandler() nats.ErrHandler {
	//TODO implement me
	panic("implement me")
}

func (o *otelConn) TLSConnectionState() (tls.ConnectionState, error) {
	//TODO implement me
	panic("implement me")
}

func (o *otelConn) ConnectedUrl() string {
	//TODO implement me
	panic("implement me")
}

func (o *otelConn) ConnectedUrlRedacted() string {
	//TODO implement me
	panic("implement me")
}

func (o *otelConn) ConnectedAddr() string {
	//TODO implement me
	panic("implement me")
}

func (o *otelConn) ConnectedServerId() string {
	//TODO implement me
	panic("implement me")
}

func (o *otelConn) ConnectedServerName() string {
	//TODO implement me
	panic("implement me")
}

func (o *otelConn) ConnectedServerVersion() string {
	//TODO implement me
	panic("implement me")
}

func (o *otelConn) ConnectedClusterName() string {
	//TODO implement me
	panic("implement me")
}

func (o *otelConn) LastError() error {
	//TODO implement me
	panic("implement me")
}

func (o *otelConn) Publish(subj string, data []byte) error {
	//TODO implement me
	panic("implement me")
}

func (o *otelConn) PublishMsg(m *nats.Msg) error {
	//TODO implement me
	panic("implement me")
}

func (o *otelConn) PublishRequest(subj, reply string, data []byte) error {
	//TODO implement me
	panic("implement me")
}

func (o *otelConn) RequestMsg(msg *nats.Msg, timeout time.Duration) (*nats.Msg, error) {
	//TODO implement me
	panic("implement me")
}

func (o *otelConn) Request(subj string, data []byte, timeout time.Duration) (*nats.Msg, error) {
	//TODO implement me
	panic("implement me")
}

func (o *otelConn) NewInbox() string {
	//TODO implement me
	panic("implement me")
}

func (o *otelConn) NewRespInbox() string {
	//TODO implement me
	panic("implement me")
}

func (o *otelConn) Subscribe(subj string, cb nats.MsgHandler) (*nats.Subscription, error) {
	//TODO implement me
	panic("implement me")
}

func (o *otelConn) ChanSubscribe(subj string, ch chan *nats.Msg) (*nats.Subscription, error) {
	//TODO implement me
	panic("implement me")
}

func (o *otelConn) ChanQueueSubscribe(subj, group string, ch chan *nats.Msg) (*nats.Subscription, error) {
	//TODO implement me
	panic("implement me")
}

func (o *otelConn) SubscribeSync(subj string) (*nats.Subscription, error) {
	//TODO implement me
	panic("implement me")
}

func (o *otelConn) QueueSubscribe(subj, queue string, cb nats.MsgHandler) (*nats.Subscription, error) {
	//TODO implement me
	panic("implement me")
}

func (o *otelConn) QueueSubscribeSync(subj, queue string) (*nats.Subscription, error) {
	//TODO implement me
	panic("implement me")
}

func (o *otelConn) QueueSubscribeSyncWithChan(subj, queue string, ch chan *nats.Msg) (*nats.Subscription, error) {
	//TODO implement me
	panic("implement me")
}

func (o *otelConn) NumSubscriptions() int {
	//TODO implement me
	panic("implement me")
}

func (o *otelConn) FlushTimeout(timeout time.Duration) (err error) {
	//TODO implement me
	panic("implement me")
}

func (o *otelConn) RTT() (time.Duration, error) {
	//TODO implement me
	panic("implement me")
}

func (o *otelConn) Flush() error {
	//TODO implement me
	panic("implement me")
}

func (o *otelConn) Buffered() (int, error) {
	//TODO implement me
	panic("implement me")
}

func (o *otelConn) Close() {
	//TODO implement me
	panic("implement me")
}

func (o *otelConn) IsClosed() bool {
	//TODO implement me
	panic("implement me")
}

func (o *otelConn) IsReconnecting() bool {
	//TODO implement me
	panic("implement me")
}

func (o *otelConn) IsConnected() bool {
	//TODO implement me
	panic("implement me")
}

func (o *otelConn) Drain() error {
	//TODO implement me
	panic("implement me")
}

func (o *otelConn) IsDraining() bool {
	//TODO implement me
	panic("implement me")
}

func (o *otelConn) Servers() []string {
	//TODO implement me
	panic("implement me")
}

func (o *otelConn) DiscoveredServers() []string {
	//TODO implement me
	panic("implement me")
}

func (o *otelConn) Status() nats.Status {
	//TODO implement me
	panic("implement me")
}

func (o *otelConn) Stats() nats.Statistics {
	//TODO implement me
	panic("implement me")
}

func (o *otelConn) MaxPayload() int64 {
	//TODO implement me
	panic("implement me")
}

func (o *otelConn) HeadersSupported() bool {
	//TODO implement me
	panic("implement me")
}

func (o *otelConn) AuthRequired() bool {
	//TODO implement me
	panic("implement me")
}

func (o *otelConn) TLSRequired() bool {
	//TODO implement me
	panic("implement me")
}

func (o *otelConn) Barrier(f func()) error {
	//TODO implement me
	panic("implement me")
}

func (o *otelConn) GetClientIP() (net.IP, error) {
	//TODO implement me
	panic("implement me")
}

func (o *otelConn) GetClientID() (uint64, error) {
	//TODO implement me
	panic("implement me")
}

func (o *otelConn) StatusChanged(statuses ...nats.Status) chan nats.Status {
	//TODO implement me
	panic("implement me")
}
