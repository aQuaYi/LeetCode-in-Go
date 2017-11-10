package Problem0472

import (
	"sort"
)

func findAllConcatenatedWordsInADict(words []string) []string {
	res := []string{}

	size := len(words)
	if size <= 1 {
		return res
	}

	sort.Strings(words)
	r := make(map[byte]int, 26)
	head := words[0][0]
	r[head] = 0
	for i := 1; i < size; i++ {
		if words[i][0] == head {
			continue
		}
		head = words[i][0]
		r[head] = i
	}

	var isCheck func(string, bool) bool
	isCheck = func(s string, isFirst bool) bool {
		if s == "" {
			return true
		}

		i, ok := r[s[0]]
		for ok && i < size && words[i] < s || (!isFirst && words[i] == s) {
			wLen := len(words[i])
			if s[:wLen] == words[i] && isCheck(s[wLen:], false) {
				return true
			}
			i++
		}

		return false
	}

	for i := 1; i < size; i++ {
		if isCheck(words[i], true) {
			res = append(res, words[i])
		}
	}

	return res
}
