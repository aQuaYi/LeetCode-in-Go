package problem0884

import (
	"strings"
)

func uncommonFromSentences(A string, B string) []string {
	m := make(map[string]int, 400)

	tmp := strings.Split(A, " ")
	for _, t := range tmp {
		m[t]++
	}

	tmp = strings.Split(B, " ")
	for _, t := range tmp {
		m[t]++
	}

	res := make([]string, 0)

	for w, c := range m {
		if c == 1 {
			res = append(res, w)
		}
	}

	return res
}
