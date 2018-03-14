package problem0691

const intMax = 1<<63 - 1

func minStickers(stickers []string, target string) int {
	size := len(stickers)

	// ss[i][j] == k 表示 stickers[i] 有 k 个字母（'a'+j)
	ss := make([][]int, size)
	for i := range ss {
		ss[i] = make([]int, 26)
		for _, c := range stickers[i] {
			ss[i][c-'a']++
		}
	}

	// dp["abc"] == m 表示，最少需要 m 个 sticker，可以组成 "abc"
	dp := make(map[string]int, size)
	dp[""] = 0

	helper(dp, ss, target)

	return dp[target]
}

func helper(dp map[string]int, ss [][]int, target string) int {
	// 如果曾经检查过 target ，直接返回答案
	if minimum, ok := dp[target]; ok {
		return minimum
	}

	// 获取 target 内各个字母的个数
	tar := make([]int, 26)
	for _, c := range target {
		tar[c-'a']++
	}

	res := intMax

	for _, s := range ss {
		// 只有当 s 包含了 target 的首字母时
		// 才从 target 中减去 s
		// 这样就确保了，每次递归时候，target 都在变小
		if s[target[0]-'a'] == 0 {
			continue
		}

		// t = target - s
		t := make([]byte, 0, len(target))
		for i := 0; i < 26; i++ {
			for j := tar[i] - s[i]; 0 < j; j-- {
				t = append(t, byte('a'+i))
			}
		}

		// 递归求解 t
		temp := helper(dp, ss, string(t))
		if temp != -1 {
			res = min(res, temp+1)
		}
	}

	// res 依然等于 intMax 说明无法组成 target
	if res == intMax {
		res = -1
	}

	dp[target] = res

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
