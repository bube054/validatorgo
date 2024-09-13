package validatorgo

import (
	"net"
	"strings"
)

func IsIPRange(str, version string) bool {
	switch version {
	case "4":
		return isValidIPv4Range(str)
	case "6":
		return isValidIPv6Range(str)
	default:
		return isValidIPv4Range(str) || isValidIPv6Range(str)
	}
}

func isValidIPv4Range(str string) bool {
	if strings.Contains(str, "/") {
		return isValidCIDR(str, 32)
	} else if strings.Contains(str, "-") {
		return isValidDashSeparatedRange(str, "4")
	}
	return false
}

func isValidIPv6Range(str string) bool {
	if strings.Contains(str, "/") {
		return isValidCIDR(str, 128)
	} else if strings.Contains(str, "-") {
		return isValidDashSeparatedRange(str, "6")
	}
	return false
}

func isValidCIDR(str string, maxPrefixLength int) bool {
	_, ipNet, err := net.ParseCIDR(str)
	if err != nil {
		return false
	}
	ones, bits := ipNet.Mask.Size()
	return bits == maxPrefixLength && ones >= 0 && ones <= maxPrefixLength
}

func isValidDashSeparatedRange(str string, version string) bool {
	parts := strings.Split(str, "-")
	if len(parts) != 2 {
		return false
	}
	startIP := strings.TrimSpace(parts[0])
	endIP := strings.TrimSpace(parts[1])

	// Parse the start and end IPs
	start := net.ParseIP(startIP)
	end := net.ParseIP(endIP)

	// Ensure both IPs are valid and of the same version
	if start == nil || end == nil || !isSameIPVersion(start, end, version) {
		return false
	}

	// Check that start IP is less than or equal to end IP
	return isIPLessThanOrEqual(start, end)
}

func isSameIPVersion(ip1, ip2 net.IP, version string) bool {
	if version == "4" {
		return ip1.To4() != nil && ip2.To4() != nil
	} else if version == "6" {
		return ip1.To16() != nil && ip2.To16() != nil && ip1.To4() == nil && ip2.To4() == nil
	}
	return false
}

func isIPLessThanOrEqual(start, end net.IP) bool {
	for i := range start {
		if start[i] < end[i] {
			return true
		} else if start[i] > end[i] {
			return false
		}
	}
	return true
}
