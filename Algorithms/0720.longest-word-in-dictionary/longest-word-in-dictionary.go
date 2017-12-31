package Problem0720

import "sort"

func longestWord(words []string) string {
	res := ""

	sort.Strings(words)

	for i := 1; i < len(words); i++ {
		n := len(words[i-1])
		if n+1 == len(words[i]) &&
			words[i-1] == words[i][:n] &&
			len(res) < n+1 {
			res = words[i]
		}
	}

	return res
}
