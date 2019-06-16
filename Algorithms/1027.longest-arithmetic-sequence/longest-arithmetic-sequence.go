package problem1027

func longestArithSeqLength(A []int) int {
	n := len(A)

	indexes := make(map[int][]int, n)
	for i := 0; i < n; i++ {
		indexes[A[i]] = append(indexes[A[i]], i)
	}

	// diff :=0

	var dfs func(int, int) int
	dfs = func(diff, j int) int {
		for _, k := range indexes[A[j]+diff] {
			if j < k {
				return 1 + dfs(diff, k)
			}
		}

		return 0
	}

	res := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			diff := A[j] - A[i]
			res = max(res, dfs(diff, j)+2)
		}
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
