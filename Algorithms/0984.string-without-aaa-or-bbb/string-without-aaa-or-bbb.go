package problem0984

import "strings"

func strWithout3a3b(A int, B int) string {
	var sb strings.Builder
	sb.Grow(A + B)

	a, b := 'a', 'b'
	if A < B {
		A, B = B, A
		a, b = b, a
	}

	for A > 0 || B > 0 {
		for i := 2; i > 0 && A > 0; i-- {
			sb.WriteRune(a)
			A--
		}
		for i := 2; i > 0 && B > 0; i-- {
			sb.WriteRune(b)
			B--
		}
	}

	return sb.String()
}
