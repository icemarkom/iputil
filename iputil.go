package iputil

import (
	"errors"
	"net"
)

// IsIPv4 reports whether supplied IP is IPv4.
func IsIPv4(a net.IPAddr) bool {
	return a.IP.To4() != nil
}

// IsIPv6 reports whether supplied IP is IPv6.
func IsIPv6(a net.IPAddr) bool {
	return a.IP.To4() == nil && a.IP.To16() != nil
}

// FindFirstIPv4 returns first available IPv4 address from the supplied list.
func FindFirstIPv4(al []net.IPAddr) (net.IP, error) {
	for _, a := range al {
		if IsIPv4(a) {
			return a.IP, nil
		}
	}
	return nil, errors.New("no IPv4 addresses found")
}

// FindFirstIPv6 returns first available IPv6 address from the supplied list.
func FindFirstIPv6(al []net.IPAddr) (net.IP, error) {
	for _, a := range al {
		if IsIPv6(a) {
			return a.IP, nil
		}
	}
	return nil, errors.New("no IPv6 addresses found")
}
