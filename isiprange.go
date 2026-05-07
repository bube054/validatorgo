package validatorgo

import (
	"net"
	"strings"
)

// A validator that checks if the string is an IP Range (version 4 or 6). If version is not provide, both versions "4" and "6" will be checked.
//
//	ok, _ := validatorgo.IsIPRange("192.168.0.0/24", "4")
//	fmt.Println(ok) // true
//	ok, _ = validatorgo.IsIPRange("192.168.0.0/33", "4")
//	fmt.Println(ok) // false
func IsIPRange(str, version string) (bool, error) {
	switch version {
	case "4":
		ok, _ := isValidIPv4Range(str)
		if ok {
			return true, nil
		}
		return false, newValidationError("IsIPRange", ErrInvalidFormat, "invalid iprange")
	case "6":
		ok, _ := isValidIPv6Range(str)
		if ok {
			return true, nil
		}
		return false, newValidationError("IsIPRange", ErrInvalidFormat, "invalid iprange")
	case "":
		ok4, _ := isValidIPv4Range(str)
		ok6, _ := isValidIPv6Range(str)
		if ok4 || ok6 {
			return true, nil
		}
		return false, newValidationError("IsIPRange", ErrInvalidFormat, "invalid iprange")
	default:
		return false, newValidationError("IsIPRange", ErrInvalidFormat, "invalid iprange")
	}
}

func isValidIPv4Range(str string) (bool, error) {
	if strings.Contains(str, "/") {
		ok, _ := isValidCIDR(str, 32)
		if ok {
			return true, nil
		}
		return false, newValidationError("isValidIPv4Range", ErrInvalidFormat, "invalid isvalidipv4range")
	} else if strings.Contains(str, "-") {
		ok, _ := isValidDashSeparatedRange(str, "4")
		if ok {
			return true, nil
		}
		return false, newValidationError("isValidIPv4Range", ErrInvalidFormat, "invalid isvalidipv4range")
	}
	return false, newValidationError("isValidIPv4Range", ErrInvalidFormat, "invalid isvalidipv4range")
}

func isValidIPv6Range(str string) (bool, error) {
	if strings.Contains(str, "/") {
		ok, _ := isValidCIDR(str, 128)
		if ok {
			return true, nil
		}
		return false, newValidationError("isValidIPv6Range", ErrInvalidFormat, "invalid isvalidipv6range")
	} else if strings.Contains(str, "-") {
		ok, _ := isValidDashSeparatedRange(str, "6")
		if ok {
			return true, nil
		}
		return false, newValidationError("isValidIPv6Range", ErrInvalidFormat, "invalid isvalidipv6range")
	}
	return false, newValidationError("isValidIPv6Range", ErrInvalidFormat, "invalid isvalidipv6range")
}

func isValidCIDR(str string, maxPrefixLength int) (bool, error) {
	_, ipNet, err := net.ParseCIDR(str)
	if err != nil {
		return false, newValidationError("isValidCIDR", ErrInvalidFormat, "invalid isvalidcidr")
	}
	ones, bits := ipNet.Mask.Size()
	if bits == maxPrefixLength && ones >= 0 && ones <= maxPrefixLength {
		return true, nil
	}
	return false, newValidationError("isValidCIDR", ErrInvalidFormat, "invalid isvalidcidr")
}

func isValidDashSeparatedRange(str string, version string) (bool, error) {
	parts := strings.Split(str, "-")
	if len(parts) != 2 {
		return false, newValidationError("isValidDashSeparatedRange", ErrInvalidFormat, "invalid isvaliddashseparatedrange")
	}
	startIP := strings.TrimSpace(parts[0])
	endIP := strings.TrimSpace(parts[1])

	// Parse the start and end IPs
	start := net.ParseIP(startIP)
	end := net.ParseIP(endIP)

	// Ensure both IPs are valid and of the same version
	sameVer, _ := isSameIPVersion(start, end, version)
	if start == nil || end == nil || !sameVer {
		return false, newValidationError("isValidDashSeparatedRange", ErrInvalidFormat, "invalid isvaliddashseparatedrange")
	}

	// Check that start IP is less than or equal to end IP
	lessEq, _ := isIPLessThanOrEqual(start, end)
	if lessEq {
		return true, nil
	}
	return false, newValidationError("isValidDashSeparatedRange", ErrInvalidFormat, "invalid isvaliddashseparatedrange")
}

func isSameIPVersion(ip1, ip2 net.IP, version string) (bool, error) {
	if version == "4" {
		if ip1.To4() != nil && ip2.To4() != nil {
			return true, nil
		}
		return false, newValidationError("isSameIPVersion", ErrInvalidFormat, "invalid issameipversion")
	} else if version == "6" {
		if ip1.To16() != nil && ip2.To16() != nil && ip1.To4() == nil && ip2.To4() == nil {
			return true, nil
		}
		return false, newValidationError("isSameIPVersion", ErrInvalidFormat, "invalid issameipversion")
	}
	return false, newValidationError("isSameIPVersion", ErrInvalidFormat, "invalid issameipversion")
}

func isIPLessThanOrEqual(start, end net.IP) (bool, error) {
	for i := range start {
		if start[i] < end[i] {
			return true, nil
		} else if start[i] > end[i] {
			return false, newValidationError("isIPLessThanOrEqual", ErrInvalidFormat, "invalid isiplessthanorequal")
		}
	}
	return true, nil
}
