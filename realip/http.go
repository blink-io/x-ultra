package realip

import (
	"net"
	"net/http"
	"strings"
)

// GetFromHTTP extracts the real client's remote IP Address.
//
// Based on proxy headers of `Headers` and `PrivateSubnets`.
//
// Fallbacks to the request's `RemoteAddr` field which is filled by the transport.
func (o *Options) GetFromHTTP(r *http.Request) string {
	for _, header := range o.Headers {
		addrs := strings.Split(r.Header.Get(header), ",")
		if ip, ok := GetIPAddress(addrs, o.PrivateSubnets); ok {
			return ip
		}
	}

	addr := strings.TrimSpace(r.RemoteAddr)
	if addr != "" {
		if ip, _, err := net.SplitHostPort(addr); err == nil {
			return ip
		}
	}
	return addr
}
