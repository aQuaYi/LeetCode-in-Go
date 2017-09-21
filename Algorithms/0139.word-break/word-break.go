package Problem0139

import (
	"sort"
)

func wordBreak(s string, wordDict []string) bool {
	if len(wordDict) == 0 {
		return false
	}

	dict := make(map[string]bool, len(wordDict))
	length := make(map[int]bool, len(wordDict))

	for _, w := range wordDict {
		length[len(w)] = true
		dict[w] = true
	}

	sizes := make([]int, 0, len(length))
	for k := range length {
		sizes = append(sizes, k)
	}

	sort.Ints(sizes)
	min, max := sizes[0], sizes[len(sizes)-1]

	var dfs func(int) bool
	dfs = func(i int) bool {
		if len(s) < i+min {
			return false
		}

		if i+min <= len(s) && len(s) <= i+max {
			return dict[s[i:]]
		}

		for _, size := range sizes {
			if dict[s[i:i+size]] && dfs(i+size) {
				return true
			}
		}

		return false
	}

	return dfs(0)
}
