package problem0468

import (
	"strconv"
	"strings"
)

func validIPAddress(IP string) string {
	switch {
	case isIPv4(IP):
		return "IPv4"
	case isIPv6(IP):
		return "IPv6"
	default:
		return "Neither"
	}
}

func isIPv4(IP string) bool {
	if !strings.Contains(IP, ".") {
		return false
	}

	ss := strings.Split(IP, ".")

	if len(ss) != 4 {
		return false
	}

	for _, s := range ss {
		if !isV4Num(s) {
			return false
		}
	}

	return true
}

func isV4Num(s string) bool {
	if len(s) == 0 {
		return false
	}

	if len(s) > 1 &&
		(s[0] < '1' || '9' < s[0]) {
		return false
	}

	n, err := strconv.Atoi(s)

	return err == nil && 0 <= n && n < 256
}

func isIPv6(IP string) bool {
	if !strings.Contains(IP, ":") {
		return false
	}

	ss := strings.Split(IP, ":")

	if len(ss) != 8 {
		return false
	}

	for _, s := range ss {
		if !isV6Num(s) {
			return false
		}
	}

	return true
}

func isV6Num(s string) bool {
	if len(s) == 0 || len(s) > 4 {
		return false
	}

	if !('0' <= s[0] && s[0] <= '9') &&
		!('a' <= s[0] && s[0] <= 'z') &&
		!('A' <= s[0] && s[0] <= 'Z') {
		return false
	}

	n, err := strconv.ParseInt(s, 16, 64)

	return err == nil && 0 <= n && n < 1<<16
}
