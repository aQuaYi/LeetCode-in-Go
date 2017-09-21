package Problem0140

import "sort"

func wordBreak(s string, wordDict []string) []string {
	if len(wordDict) == 0 {
		return []string{}
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

	// dp[i] 等于 wordBreak(s[:i+1], wordDict)
	dp := make([]int, len(s)+1)
	dp[0] = 1
	n := len(s)

	for i := 0; i <= n; i++ {
		if dp[i] == 0 {
			continue
		}

		for _, size := range sizes {
			if i+size <= n && dict[s[i:i+size]] {
				dp[i+size] += dp[i]
			}
		}
	}

	if dp[n] == 0 {
		return []string{}
	}

	res := make([]string, 0, dp[n])

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
