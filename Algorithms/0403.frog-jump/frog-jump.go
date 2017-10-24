package Problem0403

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

	var i int
	for i = 0; i < n; i++ {
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
		if hs[pos+jump+1] {
			if dfs(hs, pos+jump+1, jump+1) {
				return true
			}
		}
		if hs[pos+jump] {
			if dfs(hs, pos+jump, jump) {
				return true
			}
		}
		if jump > 1 && hs[pos+jump-1] {
			if dfs(hs, pos+jump-1, jump-1) {
				return true
			}
		}
		return false
	}

	return dfs(hs, 1, 1)
}
