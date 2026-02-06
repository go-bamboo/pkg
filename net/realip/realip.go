package realip

import (
	"errors"
	"net"
	"net/http"
	"strings"
)

var cidrs []*net.IPNet

func init() {
	maxCidrBlocks := []string{
		"127.0.0.1/8",    // localhost
		"10.0.0.0/8",     // 24-bit block
		"172.16.0.0/12",  // 20-bit block
		"192.168.0.0/16", // 16-bit block
		"169.254.0.0/16", // link local address
		"::1/128",        // localhost IPv6
		"fc00::/7",       // unique local address IPv6
		"fe80::/10",      // link local address IPv6
	}

	cidrs = make([]*net.IPNet, len(maxCidrBlocks))
	for i, maxCidrBlock := range maxCidrBlocks {
		_, cidr, _ := net.ParseCIDR(maxCidrBlock)
		cidrs[i] = cidr
	}
}

// isLocalAddress works by checking if the address is under private CIDR blocks.
// List of private CIDR blocks can be seen on :
//
// https://en.wikipedia.org/wiki/Private_network
//
// https://en.wikipedia.org/wiki/Link-local_address
func isPrivateAddress(address string) (bool, error) {
	ipAddress := net.ParseIP(address)
	if ipAddress == nil {
		return false, errors.New("address is not valid")
	}

	for i := range cidrs {
		if cidrs[i].Contains(ipAddress) {
			return true, nil
		}
	}

	return false, nil
}

// FromRequest returns the client's real IP from http request headers.
// Priority: X-Forwarded-For (leftmost) > X-Real-IP > CF-Connecting-IP > True-Client-IP > X-Client-IP > RemoteAddr
// X-Forwarded-For format is "client, proxy1, proxy2", so the leftmost IP is the real client.
func FromRequest(r *http.Request) string {
	// 1. X-Forwarded-For: 标准格式为 "客户端IP, 代理1, 代理2"，最左侧为真实客户端
	if xff := r.Header.Get("X-Forwarded-For"); xff != "" {
		parts := strings.Split(xff, ",")
		if len(parts) > 0 {
			for i := len(parts) - 1; i >= 0; i-- {
				addr := strings.TrimSpace(parts[i])
				if isPrivate, err := isPrivateAddress(addr); !isPrivate && err == nil {
					return addr
				}
			}
		}
	}

	// 2. X-Real-IP（常见于 Nginx 等）
	if xri := r.Header.Get("X-Real-IP"); xri != "" {
		xri = strings.TrimSpace(xri)
		if ip := net.ParseIP(xri); ip != nil && !ip.IsPrivate() {
			if isPrivate, err := isPrivateAddress(xri); err == nil && !isPrivate {
				return xri
			}
		}
	}

	// 4. True-Client-IP（Akamai/Cloudflare 等）
	if tc := r.Header.Get("True-Client-IP"); tc != "" {
		tc = strings.TrimSpace(tc)
		if ip := net.ParseIP(tc); ip != nil && !ip.IsPrivate() {
			if isPrivate, err := isPrivateAddress(tc); err == nil && !isPrivate {
				return tc
			}
		}
	}

	// 5. X-Client-IP
	if xc := r.Header.Get("X-Client-IP"); xc != "" {
		xc = strings.TrimSpace(xc)
		if ip := net.ParseIP(xc); ip != nil && !ip.IsPrivate() {
			if isPrivate, err := isPrivateAddress(xc); err == nil && !isPrivate {
				return xc
			}
		}
	}

	// 6. 最后回退到 RemoteAddr
	remoteIP := r.RemoteAddr
	if strings.ContainsRune(r.RemoteAddr, ':') {
		remoteIP, _, _ = net.SplitHostPort(r.RemoteAddr)
	}
	return remoteIP
}

// RealIP is depreciated, use FromRequest instead
func RealIP(r *http.Request) string {
	return FromRequest(r)
}
