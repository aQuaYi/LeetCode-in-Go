package Problem0720

import "sort"

func longestWord(words []string) string {
	res := ""

	sort.Strings(words)

	i := 0
	for i < len(words) {
		for i < len(words) && len(words[i]) != 1 {
			i++
		}

		j := i

		for j+1 < len(words) &&
			len(words[j])+1 == len(words[j+1]) &&
			words[j] == words[j+1][:len(words[j])] {
			j++
		}

		if j < len(words) && len(res) < len(words[j]) {
			res = words[j]
		}

		i = j + 1
	}

	return res
}
