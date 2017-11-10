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
	head := byte(0)
	for i := 0; i < size; i++ {
		if len(words[i]) == 0 || words[i][0] == head {
			continue
		}
		head = words[i][0]
		r[head] = i
	}

	var isCheck func(string, bool) bool
	isCheck = func(s string, isFirst bool) bool {
		i, ok := r[s[0]]
		for ok && i < size && words[i] < s {
			wLen := len(words[i])
			if wLen < len(s) && s[:wLen] == words[i] && isCheck(s[wLen:], false) {
				return true
			}
			i++
		}

		if !isFirst && i < size && words[i] == s {
			return true
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
