package Problem0686

import (
	"strings"
)

func repeatedStringMatch(a, b string) int {
	if len(a) >= len(b) {
		if a[:len(b)] == b {
			return 1
		}
		return -1
	}

	i := strings.Index(b, a)

	switch i {
	case -1:
		return -1
	case 0:
		return 1 + repeatedStringMatch(a, b[len(a):])
	default:
		if i < len(a) && a[len(a)-i:] == b[:i] {
			return 2 + repeatedStringMatch(a, b[i+len(a):])
		}
		return -1
	}
}
