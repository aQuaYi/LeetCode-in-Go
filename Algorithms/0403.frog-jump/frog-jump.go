package problem0403

func canCross(stones []int) bool {
	n := len(stones)
	if n == 0 || stones[1] != 1 {
		return false
	}

	if n == 1 || n == 2 {
		return true
	}

	last := stones[n-1]

	hs := make(map[int]bool, n)

	for i := 0; i < n; i++ {
		if i > 3 && stones[i] > stones[i-1]*2 {
			return false
		}
		hs[stones[i]] = true
	}

	var dfs func(map[int]bool, int, int) bool
	dfs = func(hs map[int]bool, pos, jump int) bool {
		if pos+jump-1 == last || pos+jump == last || pos+jump+1 == last {
			return true
		}
		// i--
		// 先跨大步
		for i := 1; -1 <= i; i-- {
			if jump+i > 0 && hs[pos+jump+i] {
				if dfs(hs, pos+jump+i, jump+i) {
					return true
				}
			}
		}

		return false
	}

	return dfs(hs, 1, 1)
}
