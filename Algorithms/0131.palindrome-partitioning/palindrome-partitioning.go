package problem0131

func partition(s string) [][]string {
	result := [][]string{}
	current := make([]string, 0, len(s))
	dfs(s, 0, current, &result)
	return result
}

func dfs(s string, i int, cur []string, result *[][]string) {
	if i == len(s) {
		tmp := make([]string, len(cur))
		copy(tmp, cur)
		*result = append(*result, tmp)
		return
	}

	for j := i; j < len(s); j++ {
		// i == 0 时，
		// 按照 len(cur[0]) 的不同，来划分 res
		// 并以此类推
		if par(s[i : j+1]) {
			dfs(s, j+1, append(cur, s[i:j+1]), result)
		}
	}
}

// s 为回文，则返回 true
func par(s string) bool {
	if len(s) <= 1 {
		return true
	}
	a, b := 0, len(s)-1
	for a < b {
		if s[a] != s[b] {
			return false
		}
		a++
		b--
	}
	return true
}
