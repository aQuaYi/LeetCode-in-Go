package Problem0077

func combine(n int, k int) [][]int {
	r := make([]bool, n+1)
	combination := make([]int, k)
	res := [][]int{}

	var dfs func(int, int)
	dfs = func(idx, begin int) {
		if idx == k {
			temp := make([]int, k)
			copy(temp, combination)
			res = append(res, temp)
			return
		}

		for i := begin; i <= n+1-k+idx; i++ {
			r[i] = true
			combination[idx] = i
			dfs(idx+1, i+1)
			r[i] = false
		}
	}

	dfs(0, 1)

	return res
}
