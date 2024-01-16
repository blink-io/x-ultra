package thrift

import (
	"crypto/tls"
	"net"
	"time"

	"github.com/apache/thrift/lib/go/thrift"
)

type quicConn struct {
}

var _ thrift.TConfigurationSetter = (*TQUIC)(nil)

type TQUIC struct {
	conn *quicConn
	// hostPort contains host:port (e.g. "asdf.com:12345"). The field is
	// only valid if addr is nil.
	hostPort string
	// addr is nil when hostPort is not "", and is only used when the
	// TSSLSocket is constructed from a net.Addr.
	addr net.Addr

	cfg *thrift.TConfiguration
}

// NewTQUICConf creates a net.Conn-backed TTransport, given a host and port.
//
// Example:
//
//	trans := thrift.NewTSSLSocketConf("localhost:9090", &TConfiguration{
//	    ConnectTimeout: time.Second, // Use 0 for no timeout
//	    SocketTimeout:  time.Second, // Use 0 for no timeout
//
//	    TLSConfig: &tls.Config{
//	        // Fill in tls config here.
//	    }
//	})
func NewTQUICConf(hostPort string, conf *thrift.TConfiguration) *TQUIC {
	if cfg := conf.GetTLSConfig(); cfg != nil && cfg.MinVersion == 0 {
		cfg.MinVersion = tls.VersionTLS13
	}
	return &TQUIC{
		hostPort: hostPort,
		cfg:      conf,
	}
}

// NewTQUICFromConnTimeout creates a TQUIC from an existing net.Conn.
func NewTQUICFromConnTimeout(addr net.Addr, conf *thrift.TConfiguration) (*TQUIC, error) {
	return &TQUIC{
		addr: addr,
		cfg:  conf,
	}, nil
}

// SetTConfiguration implements TConfigurationSetter.
//
// It can be used to change connect and socket timeouts.
func (p *TQUIC) SetTConfiguration(conf *thrift.TConfiguration) {
	p.cfg = conf
}

// SetConnTimeout sets the connect timeout
func (p *TQUIC) SetConnTimeout(timeout time.Duration) error {
	if p.cfg == nil {
		p.cfg = &thrift.TConfiguration{}
	}
	p.cfg.ConnectTimeout = timeout
	return nil
}

// SetSocketTimeout sets the socket timeout
func (p *TQUIC) SetSocketTimeout(timeout time.Duration) error {
	if p.cfg == nil {
		p.cfg = &thrift.TConfiguration{}
	}
	p.cfg.SocketTimeout = timeout
	return nil
}
