package Problem0139

import "sort"

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

	sort.Sort(sort.Reverse(sort.IntSlice(sizes)))

	var dfs func(int) bool
	dfs = func(i int) bool {
		if i == len(s) {
			return true
		}

		for _, size := range sizes {
			if i+size <= len(s) && dict[s[i:i+size]] && dfs(i+size) {
				return true
			}
		}

		return false
	}

	return dfs(0)
}
