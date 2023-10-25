package realip

import (
	"context"
	"net"
	"strings"

	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
)

func (o *Options) GetFromGRPC(ctx context.Context) string {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ""
	}
	for _, header := range o.Headers {
		mdvals := md.Get(header)
		for _, mdval := range mdvals {
			addrs := strings.Split(mdval, ",")
			if ip, ok := GetIPAddress(addrs, o.PrivateSubnets); ok {
				return ip
			}
		}
	}

	addr := GetGRPCPeerAddr(ctx)
	if addr != "" {
		if ip, _, err := net.SplitHostPort(addr); err == nil {
			return ip
		}
	}

	return addr
}

// GetGRPCPeerAddr get peer addr
func GetGRPCPeerAddr(ctx context.Context) string {
	var addr string
	if pr, ok := peer.FromContext(ctx); ok {
		if tcpAddr, ok := pr.Addr.(*net.TCPAddr); ok {
			addr = tcpAddr.IP.String()
		} else {
			addr = pr.Addr.String()
		}
	}
	return addr
}
