package problem0796

import (
	"strings"
)

func rotateString(A string, B string) bool {
	return len(A) == len(B) &&
		strings.Contains(A+A, B)
}
