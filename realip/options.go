package realip

import (
	"net"
	"strings"
)

// Range is a structure that holds the start and end of a range of IP Addresses.
type Range struct {
	Start net.IP `json:"start" yaml:"Start" toml:"Start"`
	End   net.IP `json:"end" yaml:"End" toml:"End"` // End, e.g. 255 is not a private one, but 254 is.
}

type Options struct {
	Headers        []string `json:"headers" yaml:"Headers" toml:"Headers"`
	PrivateSubnets []Range  `json:"privateSubnets" yaml:"PrivateSubnets" toml:"PrivateSubnets"`
}

// DefaultOptions is an `Options` value with some default headers and private subnets.
// See `Get` method.
var DefaultOptions = &Options{
	Headers: []string{
		"X-Real-Ip",
		"X-Forwarded-For",
		"CF-Connecting-IP",
	},
	PrivateSubnets: []Range{
		{
			Start: net.ParseIP("10.0.0.0"),
			End:   net.ParseIP("10.255.255.255"),
		},
		{
			Start: net.ParseIP("100.64.0.0"),
			End:   net.ParseIP("100.127.255.255"),
		},
		{
			Start: net.ParseIP("172.16.0.0"),
			End:   net.ParseIP("172.31.255.255"),
		},
		{
			Start: net.ParseIP("192.0.0.0"),
			End:   net.ParseIP("192.0.0.255"),
		},
		{
			Start: net.ParseIP("192.168.0.0"),
			End:   net.ParseIP("192.168.255.255"),
		},
		{
			Start: net.ParseIP("198.18.0.0"),
			End:   net.ParseIP("198.19.255.255"),
		},
	},
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
func (o *Options) AddHeader(headerKey string) *Options {
	o.Headers = append(o.Headers, headerKey)
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
