package problem1027

func longestArithSeqLength(A []int) int {
	n := len(A)

	indexes := [10001][]int{}
	for i := 0; i < n; i++ {
		indexes[A[i]] = append(indexes[A[i]], i)
	}

	var dfs func(int, int) int
	dfs = func(j, diff int) int {
		next := A[j] + diff
		if next < 0 || 10000 < next {
			return 0
		}
		for _, k := range indexes[next] {
			if j < k {
				return 1 + dfs(k, diff)
			}
		}
		return 0
	}

	res := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			diff := A[j] - A[i]
			res = max(res, 2+dfs(j, diff))
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
