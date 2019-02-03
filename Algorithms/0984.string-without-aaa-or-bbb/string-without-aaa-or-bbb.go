package problem0984

import "strings"

func strWithout3a3b(A int, B int) string {
	var sb strings.Builder
	sb.Grow(A + B)

	a, b := byte('a'), byte('b')
	if A < B {
		A, B = B, A
		a, b = b, a
	}

	for A > 0 {
		sb.WriteByte(a)
		A--
		if A > B {
			sb.WriteByte(a)
			A--
		}
		if B > 0 {
			sb.WriteByte(b)
			B--
		}
	}

	return sb.String()
}
