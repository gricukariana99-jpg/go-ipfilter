package ipfilter

import (
	"bytes"
	"net"
	"strings"
)

func IsValid(value string) bool {
	if strings.Contains(value, "/") {
		ip, _, err := net.ParseCIDR(value)

		if err != nil || ip.To4() == nil {
			return false
		}

		return true
	}

	if strings.Contains(value, "-") {
		parts := strings.Split(value, "-")

		if len(parts) != 2 {
			return false
		}

		startIP := net.ParseIP(strings.TrimSpace(parts[0]))
		endIP := net.ParseIP(strings.TrimSpace(parts[1]))

		if startIP == nil || endIP == nil {
			return false
		}

		if startIP.To4() == nil || endIP.To4() == nil {
			return false
		}

		return bytes.Compare(startIP.To4(), endIP.To4()) <= 0
	}

	ip := net.ParseIP(value)

	return ip != nil && ip.To4() != nil
}

func Match(ruleValue string, ip string) bool {
	parsedIP := net.ParseIP(ip)

	if parsedIP == nil || parsedIP.To4() == nil {
		return false
	}

	if strings.Contains(ruleValue, "/") {
		_, network, err := net.ParseCIDR(ruleValue)

		if err != nil {
			return false
		}

		return network.Contains(parsedIP)
	}

	if strings.Contains(ruleValue, "-") {
		parts := strings.Split(ruleValue, "-")

		if len(parts) != 2 {
			return false
		}

		startIP := net.ParseIP(strings.TrimSpace(parts[0]))
		endIP := net.ParseIP(strings.TrimSpace(parts[1]))

		if startIP == nil || endIP == nil {
			return false
		}

		if startIP.To4() == nil || endIP.To4() == nil {
			return false
		}

		return bytes.Compare(parsedIP.To4(), startIP.To4()) >= 0 &&
			bytes.Compare(parsedIP.To4(), endIP.To4()) <= 0
	}

	return ruleValue == ip
}
