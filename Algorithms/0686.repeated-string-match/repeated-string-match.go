package Problem0686

import (
	"strings"
)

func repeatedStringMatch(a, b string) int {
	if len(a) >= len(b) {
		if a[:len(b)] == b || a[len(a)-len(b):] == b {
			return 1
		}
		for i := 1; i < len(b)-1; i++ {
			if b[:i] == a[len(a)-i:] && b[i:] == a[:len(b)-i] {
				return 2
			}
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
