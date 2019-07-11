package problem1071

import "strings"

func gcdOfStrings(s1, s2 string) string {
	l1, l2 := len(s1), len(s2)
	d := gcd(max(l1, l2), min(l1, l2))
	p := s2[:d]
	if s1 == strings.Repeat(p, l1/d) &&
		s2 == strings.Repeat(p, l2/d) {
		return p
	}
	return ""
}

// a >= b
func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
