package iputil

import (
	"net"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestIsIPv4(t *testing.T) {
	tests := []struct {
		ip   net.IPAddr
		want bool
	}{
		{ip: net.IPAddr{IP: net.ParseIP("192.168.0.1")}, want: true},
		{ip: net.IPAddr{IP: net.ParseIP("2001:db8::1")}, want: false},
		{ip: net.IPAddr{}, want: false},
	}
	for _, tc := range tests {
		got := IsIPv4(tc.ip)
		if got != tc.want {
			t.Errorf("IsIPv4(%q)=%v, want: %v", tc.ip.String(), got, tc.want)
		}
	}
}

func TestIsIPv6(t *testing.T) {
	tests := []struct {
		ip   net.IPAddr
		want bool
	}{
		{ip: net.IPAddr{IP: net.ParseIP("192.168.0.1")}, want: false},
		{ip: net.IPAddr{IP: net.ParseIP("2001:db8::1")}, want: true},
		{ip: net.IPAddr{}, want: false},
	}
	for _, tc := range tests {
		got := IsIPv6(tc.ip)
		if got != tc.want {
			t.Errorf("IsIPv6(%q)=%v, want: %v", tc.ip.String(), got, tc.want)
		}
	}
}

func TestFindIPv4(t *testing.T) {
	tests := []struct {
		ip   []net.IPAddr
		want net.IP
	}{
		{
			ip: []net.IPAddr{
				net.IPAddr{IP: net.ParseIP("2001:db8::1")},
				net.IPAddr{IP: net.ParseIP("2001:db8::2")},
			},
			want: nil,
		},
		{
			ip: []net.IPAddr{
				net.IPAddr{IP: net.ParseIP("192.168.0.1")},
				net.IPAddr{IP: net.ParseIP("2001:db8::1")},
			},
			want: net.ParseIP("192.168.0.1"),
		},
		{
			ip: []net.IPAddr{
				net.IPAddr{IP: net.ParseIP("2001:db8::1")},
				net.IPAddr{IP: net.ParseIP("192.168.0.1")},
			},
			want: net.ParseIP("192.168.0.1"),
		},
		{
			ip: []net.IPAddr{
				net.IPAddr{IP: net.ParseIP("192.168.0.1")},
				net.IPAddr{IP: net.ParseIP("2001:db8::1")},
				net.IPAddr{IP: net.ParseIP("192.168.0.2")},
			},
			want: net.ParseIP("192.168.0.1"),
		},
		{
			ip:   []net.IPAddr{},
			want: nil,
		},
	}
	for _, tc := range tests {
		got, _ := FindFirstIPv4(tc.ip)
		if !cmp.Equal(got, tc.want, cmpopts.EquateEmpty()) {
			t.Errorf("FindFirstIPv4(%v)=%v, want: %v", tc.ip, got, tc.want)
		}
	}
}

func TestFindIPv6(t *testing.T) {
	tests := []struct {
		ip   []net.IPAddr
		want net.IP
	}{
		{
			ip: []net.IPAddr{
				net.IPAddr{IP: net.ParseIP("192.168.0.1")},
				net.IPAddr{IP: net.ParseIP("192.168.0.2")},
			},
			want: nil,
		},
		{
			ip: []net.IPAddr{
				net.IPAddr{IP: net.ParseIP("2001:db8::1")},
				net.IPAddr{IP: net.ParseIP("192.168.0.1")},
			},
			want: net.ParseIP("2001:db8::1"),
		},
		{
			ip: []net.IPAddr{
				net.IPAddr{IP: net.ParseIP("192.168.0.1")},
				net.IPAddr{IP: net.ParseIP("2001:db8::1")},
			},
			want: net.ParseIP("2001:db8::1"),
		},
		{
			ip: []net.IPAddr{
				net.IPAddr{IP: net.ParseIP("2001:db8::1")},
				net.IPAddr{IP: net.ParseIP("192.168.0.1")},
				net.IPAddr{IP: net.ParseIP("2001:db8::2")},
			},
			want: net.ParseIP("2001:db8::1"),
		},
		{
			ip:   []net.IPAddr{},
			want: nil,
		},
	}
	for _, tc := range tests {
		got, _ := FindFirstIPv6(tc.ip)
		if !cmp.Equal(got, tc.want, cmpopts.EquateEmpty()) {
			t.Errorf("FindFirstIPv6(%v)=%v, want: %v", tc.ip, got, tc.want)
		}
	}
}
