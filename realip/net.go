package realip

import (
	"net"
)

// Range is a structure that holds the start and end of a range of IP Addresses.
type Range struct {
	Start net.IP `json:"start" yaml:"start" toml:"start" msgpack:"start"`
	End   net.IP `json:"end" yaml:"end" toml:"end" msgpack:"end"` // End, e.g. 255 is not a private one, but 254 is.
}

// PrivateSubnets defines private subnets
var PrivateSubnets = []Range{
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
}

var CIDRs []*net.IPNet

var MaxCidrBlocks = []string{
	"127.0.0.1/8",    // localhost
	"10.0.0.0/8",     // 24-bit block
	"172.16.0.0/12",  // 20-bit block
	"192.168.0.0/16", // 16-bit block
	"169.254.0.0/16", // link local address
	"::1/128",        // localhost IPv6
	"fc00::/7",       // unique local address IPv6
	"fe80::/10",      // link local address IPv6
}

var HeaderXForwardedFor = "X-Forwarded-For"

// HeaderXClientIP defines standard header used by Amazon EC2, Heroku, and others
var HeaderXClientIP = "X-Client-IP"

// HeaderXRealIP defines Nginx proxy/FastCGI
var HeaderXRealIP = "X-Real-IP"

// HeaderCFConnectingIP defines Cloudflare
// @see https://support.cloudflare.com/hc/en-us/articles/200170986-How-does-Cloudflare-handle-HTTP-Request-headers-
// CF-Connecting-IP - applied to every request to the origin.
var HeaderCFConnectingIP = "CF-Connecting-IP"

// HeaderFastlyClientIP defines Fastly CDN and Firebase hosting header when forwared to a cloud function
var HeaderFastlyClientIP = "Fastly-Client-Ip"

// HeaderTrueClientIP defines Akamai and Cloudflare
var HeaderTrueClientIP = "True-Client-Ip"

func init() {
	CIDRs = make([]*net.IPNet, len(MaxCidrBlocks))
	for i, maxCidrBlock := range MaxCidrBlocks {
		_, cidr, _ := net.ParseCIDR(maxCidrBlock)
		CIDRs[i] = cidr
	}
}
