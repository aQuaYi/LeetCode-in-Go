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
	dp := make([][]string, len(s)+1)
	dp[0] = []string{""}
	n := len(s)

	for i := 0; i <= n; i++ {
		if dp[i] == nil {
			continue
		}

		for _, size := range sizes {
			if i+size <= n && dict[s[i:i+size]] {
				for _, str := range dp[i] {
					if str == "" {
						str = s[i : i+size]
					} else {
						str += " " + s[i:i+size]
					}

					dp[i+size] = append(dp[i+size], str)
				}
			}
		}
	}

	return dp[n]
}
