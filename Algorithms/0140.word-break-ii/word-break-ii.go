package Problem0140

import "sort"

func wordBreak(s string, wordDict []string) []string {
	if len(wordDict) == 0 {
		return nil
	}

	dict := make(map[string]bool, len(wordDict))
	length := make(map[int]bool, len(wordDict))

	for _, w := range wordDict {
		dict[w] = true
		length[len(w)] = true
	}

	sizes := make([]int, 0, len(length))
	for k := range length {
		sizes = append(sizes, k)
	}
	sort.Ints(sizes)

	res := make([]string, 0, len(s))

	var dfs func(int, string)
	dfs = func(i int, str string) {
		if i == len(s) {
			res = append(res, str[1:])
			return
		}

		for _, size := range sizes {
			if i+size <= len(s) && dict[s[i:i+size]] {
				dfs(i+size, str+" "+s[i:i+size])
			}
		}
	}

	dfs(0, "")

	return res
}
