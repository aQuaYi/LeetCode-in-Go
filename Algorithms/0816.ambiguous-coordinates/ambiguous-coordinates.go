package problem0816

import (
	"fmt"
	"strconv"
)

func ambiguousCoordinates(S string) []string {
	res := make([]string, 0, len(S))
	s := S[1 : len(S)-1]
	for i := 1; i < len(s); i++ {
		lefts, rights := addDot(s[:i]), addDot(s[i:])
		for _, l := range lefts {
			for _, r := range rights {
				res = append(res, connect(l, r))
			}
		}
	}
	return res
}

func addDot(s string) []string {
	res := make([]string, 0, len(s))
	if isValid(s) {
		res = append(res, s)
	}

	for i := 1; i < len(s); i++ {
		t := s[:i] + "." + s[i:]
		if isValid(t) {
			res = append(res, t)
		}
	}

	return res
}

func isValid(s string) bool {
	f, _ := strconv.ParseFloat(s, 64)
	fs := strconv.FormatFloat(f, 'f', -1, 64)
	return s == fs
}

func connect(left, right string) string {
	return fmt.Sprintf("(%s, %s)", left, right)
}
