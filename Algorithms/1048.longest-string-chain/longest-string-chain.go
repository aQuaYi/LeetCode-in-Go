package problem1048

func longestStrChain(words []string) int {
	lengths := make([][]string, 18)
	minLen := 18
	for _, w := range words {
		l := len(w)
		minLen = min(minLen, l)
		lengths[l] = append(lengths[l], w)
	}

	res := 0
	var dfs func(int, string)
	dfs = func(count int, w1 string) {
		res = max(res, count)
		l := len(w1) + 1
		for i, w2 := range lengths[l] {
			if isPredecessor(w1, w2) {
				dfs(count+1, w2)
				lengths[l][i] = ""
			}
		}
	}
	for i := 0; i < 17; i++ {
		for _, w1 := range lengths[i] {
			if w1 != "" {
				dfs(1, w1)
			}
		}
	}
	return res
}

func isPredecessor(w1, w2 string) bool {
	for i := 0; i < len(w2); i++ {
		if w1[:i] == w2[:i] && w1[i:] == w2[i+1:] {
			return true
		}
	}
	return false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
