package Problem0115

func numDistinct(s string, t string) int {
	if len(s) < len(t) {
		return 0
	}

	res := 0
	sLen, tLen := len(s), len(t)

	var dfs func(int, int)
	dfs = func(idx, begin int) {
		if idx == tLen {
			res++
			return
		}
		for i := begin; i <= sLen-tLen+idx; i++ {
			if s[i] == t[idx] {
				dfs(idx+1, i+1)
			}
		}
	}

	dfs(0, 0)

	return res
}
