package quic

import (
	"context"
	"crypto/tls"
	"net"

	"github.com/quic-go/quic-go"
)

type listener struct {
	conn *net.UDPConn
	ln   *quic.Listener
}

var _ net.Listener = (*listener)(nil)

// Listen creates a QUIC listener on the given network interface
func Listen(network, addr string, tlsConfig *tls.Config, qconf *quic.Config) (net.Listener, error) {
	udpAddr, err := net.ResolveUDPAddr(network, addr)
	if err != nil {
		return nil, &net.OpError{Op: "listen", Net: network, Source: nil, Addr: nil, Err: err}
	}
	conn, err := net.ListenUDP(network, udpAddr)
	if err != nil {
		return nil, err
	}

	ln, err := quic.Listen(conn, tlsConfig, qconf)
	if err != nil {
		return nil, err
	}
	return &listener{
		conn: conn,
		ln:   ln,
	}, nil
}

// Accept waits for and returns the next connection to the listener.
func (s *listener) Accept() (net.Conn, error) {
	conn, err := s.ln.Accept(context.Background())
	if err != nil {
		return nil, err
	}
	stream, err := conn.AcceptStream(context.Background())
	if err != nil {
		return nil, err
	}

	qconn := &Conn{
		conn:   s.conn,
		qconn:  conn,
		stream: stream,
	}
	if err != nil {
		return nil, err
	}
	return qconn, nil
}

// Close closes the listener.
// Any blocked Accept operations will be unblocked and return errors.
func (s *listener) Close() error {
	return s.ln.Close()
}

// Addr returns the listener's network address.
func (s *listener) Addr() net.Addr {
	return s.ln.Addr()
}
