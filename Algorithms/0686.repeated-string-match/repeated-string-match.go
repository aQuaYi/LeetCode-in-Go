package Problem0686

import (
	"strings"
)

func repeatedStringMatch(a, b string) int {
	lenA := len(a)
	lenB := len(b)
	i := strings.Index(b, a)
	j := i
	if i != -1 {
		for j+lenA <= lenB && b[j:j+lenA] == a {
			j += lenA
		}
	}

	res := (j - i) / lenA
	t := b
	if i < j {
		t = b[:i] + b[j:]
	}

	lenT := len(t)
	if t == "" {
		return res
	}

	if (lenT <= lenA && a[:lenT] == t) ||
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

}
