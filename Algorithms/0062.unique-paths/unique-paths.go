package Problem0062

func uniquePaths(m int, n int) int {
	res := 0

	search(0, m, 0, n, &res)

	return res
}

func search(i, m, j, n int, res *int) {
	if i == m-1 && j == n-1 {
		*res++
		return
	}

	if i == m || j == n {
		return
	}

	search(i+1, m, j, n, res)

	search(i, m, j+1, n, res)
}
