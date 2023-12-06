package realip

import (
	"net"
	"strings"
)

type Options struct {
	Headers        []string `json:"headers" yaml:"headers" toml:"headers" msgpack:"headers"`
	PrivateSubnets []Range  `json:"private_subnets" yaml:"private_subnets" toml:"private_subnets" msgpack:"private_subnets"`
}

// DefaultOptions is an `Options` value with some default headers and private subnets.
// See `Get` method.
var DefaultOptions = &Options{
	Headers: []string{
		HeaderXRealIP,
		HeaderXClientIP,
		HeaderXForwardedFor,
		HeaderCFConnectingIP,
	},
	PrivateSubnets: PrivateSubnets,
}

// AddRange adds a private subnet to "opts".
// Should be called before any use of `Get`.
func (o *Options) AddRange(start, end string) *Options {
	o.PrivateSubnets = append(o.PrivateSubnets, Range{
		Start: net.ParseIP(start),
		End:   net.ParseIP(end),
	})
	return o
}

// AddHeader adds a proxy remote address header to "opts".
// Should be called before any use of `Get`.
func (o *Options) AddHeader(header string) *Options {
	o.Headers = append(o.Headers, header)
	return o
}

// Get extracts the real client's remote IP Address.
func (o *Options) Get(addrs ...string) string {
	for _, addr := range addrs {
		addrs := strings.Split(addr, ",")
		if ip, ok := GetIPAddress(addrs, o.PrivateSubnets); ok {
			return ip
		}
	}
	return ""
}
