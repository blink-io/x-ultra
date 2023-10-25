package realip

import (
	"bytes"
	"net"
	"strings"
)

// InRange reports whether a given IP Address is within a range given.
func InRange(r Range, ipAddress net.IP) bool {
	return bytes.Compare(ipAddress, r.Start) >= 0 && bytes.Compare(ipAddress, r.End) < 0
}

// IsPrivateSubnet reports whether this "ipAddress" is in a private subnet.
func IsPrivateSubnet(ipAddress net.IP, privateRanges []Range) bool {
	if ipCheck := ipAddress.To4(); ipCheck != nil {
		for _, r := range privateRanges {
			if InRange(r, ipAddress) {
				return true
			}
		}
	}
	return false
}

// GetIPAddress returns a valid public IP Address from a collection of IP Addresses
// and a range of private subnets.
//
// Reports whether a valid IP was found.
func GetIPAddress(addrs []string, privateRanges []Range) (string, bool) {
	for i := len(addrs) - 1; i >= 0; i-- {
		ip := strings.TrimSpace(addrs[i])
		realIP := net.ParseIP(ip)
		if !realIP.IsGlobalUnicast() || IsPrivateSubnet(realIP, privateRanges) {
			continue
		}
		return ip, true

	}

	return "", false
}
