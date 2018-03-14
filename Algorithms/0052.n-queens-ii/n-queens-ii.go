package problem0052

func totalNQueens(n int) int {
	if n == 0 {
		return 0
	}

	cols := make([]bool, n)
	// 记录 '\' 方向的对角线的占用情况
	d1 := make([]bool, 2*n)
	// 记录 '/' 方向的对角线的占用情况
	d2 := make([]bool, 2*n)

	res := 0

	dfs(0, cols, d1, d2, &res)

	return res
}

func dfs(r int, cols, d1, d2 []bool, res *int) {
	n := len(cols)

	if r == n {
		*res++
		return
	}

	for c := 0; c < n; c++ {
		id1 := r - c + n
		id2 := 2*n - r - c - 1
		if !cols[c] && !d1[id1] && !d2[id2] {

			// 标记占用
			cols[c], d1[id1], d2[id2] = true, true, true

			dfs(r+1, cols, d1, d2, res)

			// 解除标记
			cols[c], d1[id1], d2[id2] = false, false, false
		}
	}
}
