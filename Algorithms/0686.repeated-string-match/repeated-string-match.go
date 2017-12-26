package Problem0686

import (
	"strings"
)

func repeatedStringMatch(a, b string) int {
	lenA := len(a)
	lenB := len(b)
	i := strings.Index(b, a)
	j := i + lenA
	if i != -1 {
		for j+lenA <= lenB && b[j:j+lenA] == a {
			j += lenA
		}
	}

	t := b[:i] + b[j:]

	lenT := len(t)
	res := (lenB - lenT) / lenA
	if t == "" {
		return res
	}

	if (lenT < lenA && a[:lenT] == t) ||
		(0 <= lenA-lenT && a[lenA-lenT:] == t) {
		return res + 1
	}

	for i := 1; i < lenT-1; i++ {
		if 0 <= lenA-i && t[:i] == a[lenA-i:] &&
			lenT-i <= lenA && t[i:] == a[:lenT-i] {
			return res + 2
		}
	}

	return -1

	// if len(a) >= len(b) {
	// 	if a[:len(b)] == b || a[len(a)-len(b):] == b {
	// 		return 1
	// 	}
	// 	for i := 1; i < len(b)-1; i++ {
	// 		if b[:i] == a[len(a)-i:] && b[i:] == a[:len(b)-i] {
	// 			return 2
	// 		}
	// 	}
	// 	return -1
	// }

	// i := strings.Index(b, a)

	// switch i {
	// case -1:
	// 	return -1
	// case 0:
	// 	return 1 + repeatedStringMatch(a, b[len(a):])
	// default:
	// 	if i < len(a) && a[len(a)-i:] == b[:i] {
	// 		return 2 + repeatedStringMatch(a, b[i+len(a):])
	// 	}
	// 	return -1
	// }

}
