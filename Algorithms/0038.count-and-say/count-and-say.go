package Problem0038

import (
	"strconv"
)

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	return say(countAndSay(n - 1))
}

func say(s string) string {
	res := ""

	words := split(s)
	for _, w := range words {
		res += strconv.Itoa(len(w)) + w[0:1]
	}

	return res
}

func split(s string) []string {
	res := []string{}
	i, j := 0, 1

	for j < len(s) {
		if s[i] != s[j] {
			res = append(res, s[i:j])
			i = j
		}
		j++
	}

	res = append(res, s[i:])

	return res
}
