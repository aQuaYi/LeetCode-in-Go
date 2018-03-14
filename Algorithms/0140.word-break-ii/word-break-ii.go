package problem0140

import "sort"

func wordBreak(s string, wordDict []string) []string {
	if len(wordDict) == 0 {
		return []string{}
	}

	// dict 方便查找 wordDict 中的单词
	dict := make(map[string]bool, len(wordDict))
	length := make(map[int]bool, len(wordDict))

	for _, w := range wordDict {
		dict[w] = true
		length[len(w)] = true
	}

	// sizes 是 wordDict 中单词长度的集合
	// sizes 可以加快 for 循环
	sizes := make([]int, 0, len(length))
	for k := range length {
		sizes = append(sizes, k)
	}
	sort.Ints(sizes)

	n := len(s)

	// dp[i] 等于 len(wordBreak(s[:i+1], wordDict))
	dp := make([]float64, len(s)+1)
	dp[0] = 1

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

	// 利用 dp[n] 统计解的数量，可以避免无用功
	// 取消下一行的注释符号，看看各个 test case 的 dp
	// fmt.Println(dp)
	if dp[n] == 0 {
		return []string{}
	}

	res := make([]string, 0, int(dp[n]))

	// 利用 dfs 获取所有的解
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
