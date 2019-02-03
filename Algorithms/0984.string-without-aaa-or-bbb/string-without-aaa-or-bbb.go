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

	aab := []byte{a, a, b}

	for A > B && B > 0 {
		sb.Write(aab)
		A -= 2
		B--
	}

	for A > 0 || B > 0 {
		sb.WriteByte(a)
		A--
		if B > 0 {
			sb.WriteByte(b)
			B--
		}
	}

	return sb.String()
}
